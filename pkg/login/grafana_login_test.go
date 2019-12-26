package login

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/smartems/smartems/pkg/bus"
	m "github.com/smartems/smartems/pkg/models"
)

func TestGrafanaLogin(t *testing.T) {
	Convey("Login using Grafana DB", t, func() {
		smartemsLoginScenario("When login with non-existing user", func(sc *smartemsLoginScenarioContext) {
			sc.withNonExistingUser()
			err := loginUsingGrafanaDB(sc.loginUserQuery)

			Convey("it should result in user not found error", func() {
				So(err, ShouldEqual, m.ErrUserNotFound)
			})

			Convey("it should not call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeFalse)
			})

			Convey("it should not pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldBeNil)
			})
		})

		smartemsLoginScenario("When login with invalid credentials", func(sc *smartemsLoginScenarioContext) {
			sc.withInvalidPassword()
			err := loginUsingGrafanaDB(sc.loginUserQuery)

			Convey("it should result in invalid credentials error", func() {
				So(err, ShouldEqual, ErrInvalidCredentials)
			})

			Convey("it should call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeTrue)
			})

			Convey("it should not pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldBeNil)
			})
		})

		smartemsLoginScenario("When login with valid credentials", func(sc *smartemsLoginScenarioContext) {
			sc.withValidCredentials()
			err := loginUsingGrafanaDB(sc.loginUserQuery)

			Convey("it should not result in error", func() {
				So(err, ShouldBeNil)
			})

			Convey("it should call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeTrue)
			})

			Convey("it should pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldNotBeNil)
				So(sc.loginUserQuery.User.Login, ShouldEqual, sc.loginUserQuery.Username)
				So(sc.loginUserQuery.User.Password, ShouldEqual, sc.loginUserQuery.Password)
			})
		})

		smartemsLoginScenario("When login with disabled user", func(sc *smartemsLoginScenarioContext) {
			sc.withDisabledUser()
			err := loginUsingGrafanaDB(sc.loginUserQuery)

			Convey("it should return user is disabled error", func() {
				So(err, ShouldEqual, ErrUserDisabled)
			})

			Convey("it should not call password validation", func() {
				So(sc.validatePasswordCalled, ShouldBeFalse)
			})

			Convey("it should not pupulate user object", func() {
				So(sc.loginUserQuery.User, ShouldBeNil)
			})
		})
	})
}

type smartemsLoginScenarioContext struct {
	loginUserQuery         *m.LoginUserQuery
	validatePasswordCalled bool
}

type smartemsLoginScenarioFunc func(c *smartemsLoginScenarioContext)

func smartemsLoginScenario(desc string, fn smartemsLoginScenarioFunc) {
	Convey(desc, func() {
		origValidatePassword := validatePassword

		sc := &smartemsLoginScenarioContext{
			loginUserQuery: &m.LoginUserQuery{
				Username:  "user",
				Password:  "pwd",
				IpAddress: "192.168.1.1:56433",
			},
			validatePasswordCalled: false,
		}

		defer func() {
			validatePassword = origValidatePassword
		}()

		fn(sc)
	})
}

func mockPasswordValidation(valid bool, sc *smartemsLoginScenarioContext) {
	validatePassword = func(providedPassword string, userPassword string, userSalt string) error {
		sc.validatePasswordCalled = true

		if !valid {
			return ErrInvalidCredentials
		}

		return nil
	}
}

func (sc *smartemsLoginScenarioContext) getUserByLoginQueryReturns(user *m.User) {
	bus.AddHandler("test", func(query *m.GetUserByLoginQuery) error {
		if user == nil {
			return m.ErrUserNotFound
		}

		query.Result = user
		return nil
	})
}

func (sc *smartemsLoginScenarioContext) withValidCredentials() {
	sc.getUserByLoginQueryReturns(&m.User{
		Id:       1,
		Login:    sc.loginUserQuery.Username,
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(true, sc)
}

func (sc *smartemsLoginScenarioContext) withNonExistingUser() {
	sc.getUserByLoginQueryReturns(nil)
}

func (sc *smartemsLoginScenarioContext) withInvalidPassword() {
	sc.getUserByLoginQueryReturns(&m.User{
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(false, sc)
}

func (sc *smartemsLoginScenarioContext) withDisabledUser() {
	sc.getUserByLoginQueryReturns(&m.User{
		IsDisabled: true,
	})
}
