package login

import (
	"time"

	"github.com/smartems/smartems/pkg/bus"
	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/setting"
)

var (
	maxInvalidLoginAttempts int64 = 5
	loginAttemptsWindow           = time.Minute * 5
)

var validateLoginAttempts = func(username string) error {
	if setting.DisableBruteForceLoginProtection {
		return nil
	}

	loginAttemptCountQuery := m.GetUserLoginAttemptCountQuery{
		Username: username,
		Since:    time.Now().Add(-loginAttemptsWindow),
	}

	if err := bus.Dispatch(&loginAttemptCountQuery); err != nil {
		return err
	}

	if loginAttemptCountQuery.Result >= maxInvalidLoginAttempts {
		return ErrTooManyLoginAttempts
	}

	return nil
}

var saveInvalidLoginAttempt = func(query *m.LoginUserQuery) error {
	if setting.DisableBruteForceLoginProtection {
		return nil
	}

	loginAttemptCommand := m.CreateLoginAttemptCommand{
		Username:  query.Username,
		IpAddress: query.IpAddress,
	}

	return bus.Dispatch(&loginAttemptCommand)
}
