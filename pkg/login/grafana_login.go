package login

import (
	"crypto/subtle"

	"github.com/smartems/smartems/pkg/bus"
	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/util"
)

var validatePassword = func(providedPassword string, userPassword string, userSalt string) error {
	passwordHashed, err := util.EncodePassword(providedPassword, userSalt)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare([]byte(passwordHashed), []byte(userPassword)) != 1 {
		return ErrInvalidCredentials
	}

	return nil
}

var loginUsingGrafanaDB = func(query *m.LoginUserQuery) error {
	userQuery := m.GetUserByLoginQuery{LoginOrEmail: query.Username}

	if err := bus.Dispatch(&userQuery); err != nil {
		return err
	}

	user := userQuery.Result

	if user.IsDisabled {
		return ErrUserDisabled
	}

	if err := validatePassword(query.Password, user.Password, user.Salt); err != nil {
		return err
	}

	query.User = user
	return nil
}
