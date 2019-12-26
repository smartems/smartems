package usagestats

import (
	"context"
	"time"

	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/login/social"
	"github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/services/sqlstore"

	"github.com/smartems/smartems/pkg/infra/log"
	"github.com/smartems/smartems/pkg/registry"
	"github.com/smartems/smartems/pkg/setting"
)

var metricsLogger log.Logger = log.New("metrics")

func init() {
	registry.RegisterService(&UsageStatsService{})
}

type UsageStatsService struct {
	Cfg      *setting.Cfg       `inject:""`
	Bus      bus.Bus            `inject:""`
	SQLStore *sqlstore.SqlStore `inject:""`
	License  models.Licensing   `inject:""`

	oauthProviders map[string]bool
}

func (uss *UsageStatsService) Init() error {

	uss.oauthProviders = social.GetOAuthProviders(uss.Cfg)
	return nil
}

func (uss *UsageStatsService) Run(ctx context.Context) error {
	uss.updateTotalStats()

	onceEveryDayTick := time.NewTicker(time.Hour * 24)
	everyMinuteTicker := time.NewTicker(time.Minute)
	defer onceEveryDayTick.Stop()
	defer everyMinuteTicker.Stop()

	for {
		select {
		case <-onceEveryDayTick.C:
			uss.sendUsageStats(uss.oauthProviders)
		case <-everyMinuteTicker.C:
			uss.updateTotalStats()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
