package alerting

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tlog "github.com/opentracing/opentracing-go/log"

	"github.com/benbjohnson/clock"
	"github.com/smartems/smartems/pkg/infra/log"
	"github.com/smartems/smartems/pkg/registry"
	"github.com/smartems/smartems/pkg/services/rendering"
	"github.com/smartems/smartems/pkg/setting"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
)

// AlertEngine is the background process that
// schedules alert evaluations and makes sure notifications
// are sent.
type AlertEngine struct {
	RenderService rendering.Service `inject:""`

	execQueue     chan *Job
	ticker        *Ticker
	scheduler     scheduler
	evalHandler   evalHandler
	ruleReader    ruleReader
	log           log.Logger
	resultHandler resultHandler
}

func init() {
	registry.RegisterService(&AlertEngine{})
}

// IsDisabled returns true if the alerting service is disable for this instance.
func (e *AlertEngine) IsDisabled() bool {
	return !setting.AlertingEnabled || !setting.ExecuteAlerts
}

// Init initalizes the AlertingService.
func (e *AlertEngine) Init() error {
	e.ticker = NewTicker(time.Now(), time.Second*0, clock.New())
	e.execQueue = make(chan *Job, 1000)
	e.scheduler = newScheduler()
	e.evalHandler = NewEvalHandler()
	e.ruleReader = newRuleReader()
	e.log = log.New("alerting.engine")
	e.resultHandler = newResultHandler(e.RenderService)
	return nil
}

// Run starts the alerting service background process.
func (e *AlertEngine) Run(ctx context.Context) error {
	alertGroup, ctx := errgroup.WithContext(ctx)
	alertGroup.Go(func() error { return e.alertingTicker(ctx) })
	alertGroup.Go(func() error { return e.runJobDispatcher(ctx) })

	err := alertGroup.Wait()
	return err
}

func (e *AlertEngine) alertingTicker(smartemsCtx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e.log.Error("Scheduler Panic: stopping alertingTicker", "error", err, "stack", log.Stack(1))
		}
	}()

	tickIndex := 0

	for {
		select {
		case <-smartemsCtx.Done():
			return smartemsCtx.Err()
		case tick := <-e.ticker.C:
			// TEMP SOLUTION update rules ever tenth tick
			if tickIndex%10 == 0 {
				e.scheduler.Update(e.ruleReader.fetch())
			}

			e.scheduler.Tick(tick, e.execQueue)
			tickIndex++
		}
	}
}

func (e *AlertEngine) runJobDispatcher(smartemsCtx context.Context) error {
	dispatcherGroup, alertCtx := errgroup.WithContext(smartemsCtx)

	for {
		select {
		case <-smartemsCtx.Done():
			return dispatcherGroup.Wait()
		case job := <-e.execQueue:
			dispatcherGroup.Go(func() error { return e.processJobWithRetry(alertCtx, job) })
		}
	}
}

var (
	unfinishedWorkTimeout = time.Second * 5
)

func (e *AlertEngine) processJobWithRetry(smartemsCtx context.Context, job *Job) error {
	defer func() {
		if err := recover(); err != nil {
			e.log.Error("Alert Panic", "error", err, "stack", log.Stack(1))
		}
	}()

	cancelChan := make(chan context.CancelFunc, setting.AlertingMaxAttempts*2)
	attemptChan := make(chan int, 1)

	// Initialize with first attemptID=1
	attemptChan <- 1
	job.SetRunning(true)

	for {
		select {
		case <-smartemsCtx.Done():
			// In case smartems server context is cancel, let a chance to job processing
			// to finish gracefully - by waiting a timeout duration - before forcing its end.
			unfinishedWorkTimer := time.NewTimer(unfinishedWorkTimeout)
			select {
			case <-unfinishedWorkTimer.C:
				return e.endJob(smartemsCtx.Err(), cancelChan, job)
			case <-attemptChan:
				return e.endJob(nil, cancelChan, job)
			}
		case attemptID, more := <-attemptChan:
			if !more {
				return e.endJob(nil, cancelChan, job)
			}
			go e.processJob(attemptID, attemptChan, cancelChan, job)
		}
	}
}

func (e *AlertEngine) endJob(err error, cancelChan chan context.CancelFunc, job *Job) error {
	job.SetRunning(false)
	close(cancelChan)
	for cancelFn := range cancelChan {
		cancelFn()
	}
	return err
}

func (e *AlertEngine) processJob(attemptID int, attemptChan chan int, cancelChan chan context.CancelFunc, job *Job) {
	defer func() {
		if err := recover(); err != nil {
			e.log.Error("Alert Panic", "error", err, "stack", log.Stack(1))
		}
	}()

	alertCtx, cancelFn := context.WithTimeout(context.Background(), setting.AlertingEvaluationTimeout)
	cancelChan <- cancelFn
	span := opentracing.StartSpan("alert execution")
	alertCtx = opentracing.ContextWithSpan(alertCtx, span)

	evalContext := NewEvalContext(alertCtx, job.Rule)
	evalContext.Ctx = alertCtx

	go func() {
		defer func() {
			if err := recover(); err != nil {
				e.log.Error("Alert Panic", "error", err, "stack", log.Stack(1))
				ext.Error.Set(span, true)
				span.LogFields(
					tlog.Error(fmt.Errorf("%v", err)),
					tlog.String("message", "failed to execute alert rule. panic was recovered."),
				)
				span.Finish()
				close(attemptChan)
			}
		}()

		e.evalHandler.Eval(evalContext)

		span.SetTag("alertId", evalContext.Rule.ID)
		span.SetTag("dashboardId", evalContext.Rule.DashboardID)
		span.SetTag("firing", evalContext.Firing)
		span.SetTag("nodatapoints", evalContext.NoDataFound)
		span.SetTag("attemptID", attemptID)

		if evalContext.Error != nil {
			ext.Error.Set(span, true)
			span.LogFields(
				tlog.Error(evalContext.Error),
				tlog.String("message", "alerting execution attempt failed"),
			)
			if attemptID < setting.AlertingMaxAttempts {
				span.Finish()
				e.log.Debug("Job Execution attempt triggered retry", "timeMs", evalContext.GetDurationMs(), "alertId", evalContext.Rule.ID, "name", evalContext.Rule.Name, "firing", evalContext.Firing, "attemptID", attemptID)
				attemptChan <- (attemptID + 1)
				return
			}
		}

		// create new context with timeout for notifications
		resultHandleCtx, resultHandleCancelFn := context.WithTimeout(context.Background(), setting.AlertingNotificationTimeout)
		cancelChan <- resultHandleCancelFn

		// override the context used for evaluation with a new context for notifications.
		// This makes it possible for notifiers to execute when datasources
		// dont respond within the timeout limit. We should rewrite this so notifications
		// dont reuse the evalContext and get its own context.
		evalContext.Ctx = resultHandleCtx
		evalContext.Rule.State = evalContext.GetNewState()
		if err := e.resultHandler.handle(evalContext); err != nil {
			if xerrors.Is(err, context.Canceled) {
				e.log.Debug("Result handler returned context.Canceled")
			} else if xerrors.Is(err, context.DeadlineExceeded) {
				e.log.Debug("Result handler returned context.DeadlineExceeded")
			} else {
				e.log.Error("Failed to handle result", "err", err)
			}
		}

		span.Finish()
		e.log.Debug("Job Execution completed", "timeMs", evalContext.GetDurationMs(), "alertId", evalContext.Rule.ID, "name", evalContext.Rule.Name, "firing", evalContext.Firing, "attemptID", attemptID)
		close(attemptChan)
	}()
}
