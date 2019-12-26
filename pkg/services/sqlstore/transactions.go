package sqlstore

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/infra/log"
	"github.com/smartems/smartems/pkg/util/errutil"
	sqlite3 "github.com/mattn/go-sqlite3"
)

// WithTransactionalDbSession calls the callback with an session within a transaction
func (ss *SqlStore) WithTransactionalDbSession(ctx context.Context, callback dbTransactionFunc) error {
	return inTransactionWithRetryCtx(ctx, ss.engine, callback, 0)
}

func (ss *SqlStore) InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return ss.inTransactionWithRetry(ctx, fn, 0)
}

func (ss *SqlStore) inTransactionWithRetry(ctx context.Context, fn func(ctx context.Context) error, retry int) error {
	return inTransactionWithRetryCtx(ctx, ss.engine, func(sess *DBSession) error {
		withValue := context.WithValue(ctx, ContextSessionName, sess)
		return fn(withValue)
	}, retry)
}

func inTransactionWithRetry(callback dbTransactionFunc, retry int) error {
	return inTransactionWithRetryCtx(context.Background(), x, callback, retry)
}

func inTransactionWithRetryCtx(ctx context.Context, engine *xorm.Engine, callback dbTransactionFunc, retry int) error {
	sess, err := startSession(ctx, engine, true)
	if err != nil {
		return err
	}

	defer sess.Close()

	err = callback(sess)

	// special handling of database locked errors for sqlite, then we can retry 5 times
	if sqlError, ok := err.(sqlite3.Error); ok && retry < 5 && sqlError.Code ==
		sqlite3.ErrLocked || sqlError.Code == sqlite3.ErrBusy {
		if rollErr := sess.Rollback(); rollErr != nil {
			return errutil.Wrapf(err, "Rolling back transaction due to error failed: %s", rollErr)
		}

		time.Sleep(time.Millisecond * time.Duration(10))
		sqlog.Info("Database locked, sleeping then retrying", "error", err, "retry", retry)
		return inTransactionWithRetry(callback, retry+1)
	}

	if err != nil {
		if rollErr := sess.Rollback(); rollErr != nil {
			return errutil.Wrapf(err, "Rolling back transaction due to error failed: %s", rollErr)
		}
		return err
	}
	if err := sess.Commit(); err != nil {
		return err
	}

	if len(sess.events) > 0 {
		for _, e := range sess.events {
			if err = bus.Publish(e); err != nil {
				log.Error(3, "Failed to publish event after commit. error: %v", err)
			}
		}
	}

	return nil
}

func inTransaction(callback dbTransactionFunc) error {
	return inTransactionWithRetry(callback, 0)
}

func inTransactionCtx(ctx context.Context, callback dbTransactionFunc) error {
	return inTransactionWithRetryCtx(ctx, x, callback, 0)
}
