package middleware

import (
	"context"
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gopkg.in/macaron.v1"

	"github.com/smartems/smartems/pkg/api/dtos"
	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/infra/remotecache"
	authproxy "github.com/smartems/smartems/pkg/middleware/auth_proxy"
	"github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/services/auth"
	"github.com/smartems/smartems/pkg/services/login"
	"github.com/smartems/smartems/pkg/setting"
	"github.com/smartems/smartems/pkg/util"
)

const errorTemplate = "error-template"

func mockGetTime() {
	var timeSeed int64
	getTime = func() time.Time {
		fakeNow := time.Unix(timeSeed, 0)
		timeSeed++
		return fakeNow
	}
}

func resetGetTime() {
	getTime = time.Now
}

func TestMiddleWareSecurityHeaders(t *testing.T) {
	setting.ERR_TEMPLATE_NAME = errorTemplate

	Convey("Given the smartems middleware", t, func() {

		middlewareScenario(t, "middleware should get correct x-xss-protection header", func(sc *scenarioContext) {
			setting.XSSProtectionHeader = true
			sc.fakeReq("GET", "/api/").exec()
			So(sc.resp.Header().Get("X-XSS-Protection"), ShouldEqual, "1; mode=block")
		})

		middlewareScenario(t, "middleware should not get x-xss-protection when disabled", func(sc *scenarioContext) {
			setting.XSSProtectionHeader = false
			sc.fakeReq("GET", "/api/").exec()
			So(sc.resp.Header().Get("X-XSS-Protection"), ShouldBeEmpty)
		})

		middlewareScenario(t, "middleware should add correct Strict-Transport-Security header", func(sc *scenarioContext) {
			setting.StrictTransportSecurity = true
			setting.Protocol = setting.HTTPS
			setting.StrictTransportSecurityMaxAge = 64000
			sc.fakeReq("GET", "/api/").exec()
			So(sc.resp.Header().Get("Strict-Transport-Security"), ShouldEqual, "max-age=64000")
			setting.StrictTransportSecurityPreload = true
			sc.fakeReq("GET", "/api/").exec()
			So(sc.resp.Header().Get("Strict-Transport-Security"), ShouldEqual, "max-age=64000; preload")
			setting.StrictTransportSecuritySubDomains = true
			sc.fakeReq("GET", "/api/").exec()
			So(sc.resp.Header().Get("Strict-Transport-Security"), ShouldEqual, "max-age=64000; preload; includeSubDomains")
		})
	})
}

func TestMiddlewareContext(t *testing.T) {
	setting.ERR_TEMPLATE_NAME = errorTemplate

	Convey("Given the smartems middleware", t, func() {
		middlewareScenario(t, "middleware should add context to injector", func(sc *scenarioContext) {
			sc.fakeReq("GET", "/").exec()
			So(sc.context, ShouldNotBeNil)
		})

		middlewareScenario(t, "Default middleware should allow get request", func(sc *scenarioContext) {
			sc.fakeReq("GET", "/").exec()
			So(sc.resp.Code, ShouldEqual, 200)
		})

		middlewareScenario(t, "middleware should add Cache-Control header for requests to API", func(sc *scenarioContext) {
			sc.fakeReq("GET", "/api/search").exec()
			So(sc.resp.Header().Get("Cache-Control"), ShouldEqual, "no-cache")
			So(sc.resp.Header().Get("Pragma"), ShouldEqual, "no-cache")
			So(sc.resp.Header().Get("Expires"), ShouldEqual, "-1")
		})

		middlewareScenario(t, "middleware should not add Cache-Control header for requests to datasource proxy API", func(sc *scenarioContext) {
			sc.fakeReq("GET", "/api/datasources/proxy/1/test").exec()
			So(sc.resp.Header().Get("Cache-Control"), ShouldBeEmpty)
			So(sc.resp.Header().Get("Pragma"), ShouldBeEmpty)
			So(sc.resp.Header().Get("Expires"), ShouldBeEmpty)
		})

		middlewareScenario(t, "middleware should add Cache-Control header for requests with html response", func(sc *scenarioContext) {
			sc.handler(func(c *models.ReqContext) {
				data := &dtos.IndexViewData{
					User:     &dtos.CurrentUser{},
					Settings: map[string]interface{}{},
					NavTree:  []*dtos.NavLink{},
				}
				c.HTML(200, "index-template", data)
			})
			sc.fakeReq("GET", "/").exec()
			So(sc.resp.Code, ShouldEqual, 200)
			So(sc.resp.Header().Get("Cache-Control"), ShouldEqual, "no-cache")
			So(sc.resp.Header().Get("Pragma"), ShouldEqual, "no-cache")
			So(sc.resp.Header().Get("Expires"), ShouldEqual, "-1")
		})

		middlewareScenario(t, "middleware should add X-Frame-Options header with deny for request when not allowing embedding", func(sc *scenarioContext) {
			sc.fakeReq("GET", "/api/search").exec()
			So(sc.resp.Header().Get("X-Frame-Options"), ShouldEqual, "deny")
		})

		middlewareScenario(t, "middleware should not add X-Frame-Options header for request when allowing embedding", func(sc *scenarioContext) {
			setting.AllowEmbedding = true
			sc.fakeReq("GET", "/api/search").exec()
			So(sc.resp.Header().Get("X-Frame-Options"), ShouldBeEmpty)
		})

		middlewareScenario(t, "Invalid api key", func(sc *scenarioContext) {
			sc.apiKey = "invalid_key_test"
			sc.fakeReq("GET", "/").exec()

			Convey("Should not init session", func() {
				So(sc.resp.Header().Get("Set-Cookie"), ShouldBeEmpty)
			})

			Convey("Should return 401", func() {
				So(sc.resp.Code, ShouldEqual, 401)
				So(sc.respJson["message"], ShouldEqual, errStringInvalidAPIKey)
			})
		})

		middlewareScenario(t, "Valid api key", func(sc *scenarioContext) {
			keyhash, err := util.EncodePassword("v5nAwpMafFP6znaS4urhdWDLS5511M42", "asd")
			So(err, ShouldBeNil)

			bus.AddHandler("test", func(query *models.GetApiKeyByNameQuery) error {
				query.Result = &models.ApiKey{OrgId: 12, Role: models.ROLE_EDITOR, Key: keyhash}
				return nil
			})

			sc.fakeReq("GET", "/").withValidApiKey().exec()

			Convey("Should return 200", func() {
				So(sc.resp.Code, ShouldEqual, 200)
			})

			Convey("Should init middleware context", func() {
				So(sc.context.IsSignedIn, ShouldEqual, true)
				So(sc.context.OrgId, ShouldEqual, 12)
				So(sc.context.OrgRole, ShouldEqual, models.ROLE_EDITOR)
			})
		})

		middlewareScenario(t, "Valid api key, but does not match db hash", func(sc *scenarioContext) {
			keyhash := "Something_not_matching"

			bus.AddHandler("test", func(query *models.GetApiKeyByNameQuery) error {
				query.Result = &models.ApiKey{OrgId: 12, Role: models.ROLE_EDITOR, Key: keyhash}
				return nil
			})

			sc.fakeReq("GET", "/").withValidApiKey().exec()

			Convey("Should return api key invalid", func() {
				So(sc.resp.Code, ShouldEqual, 401)
				So(sc.respJson["message"], ShouldEqual, errStringInvalidAPIKey)
			})
		})

		middlewareScenario(t, "Valid api key, but expired", func(sc *scenarioContext) {
			mockGetTime()
			defer resetGetTime()

			keyhash, err := util.EncodePassword("v5nAwpMafFP6znaS4urhdWDLS5511M42", "asd")
			So(err, ShouldBeNil)

			bus.AddHandler("test", func(query *models.GetApiKeyByNameQuery) error {
				// api key expired one second before
				expires := getTime().Add(-1 * time.Second).Unix()
				query.Result = &models.ApiKey{OrgId: 12, Role: models.ROLE_EDITOR, Key: keyhash,
					Expires: &expires}
				return nil
			})

			sc.fakeReq("GET", "/").withValidApiKey().exec()

			Convey("Should return 401", func() {
				So(sc.resp.Code, ShouldEqual, 401)
				So(sc.respJson["message"], ShouldEqual, "Expired API key")
			})
		})

		middlewareScenario(t, "Non-expired auth token in cookie which not are being rotated", func(sc *scenarioContext) {
			sc.withTokenSessionCookie("token")

			bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
				query.Result = &models.SignedInUser{OrgId: 2, UserId: 12}
				return nil
			})

			sc.userAuthTokenService.LookupTokenProvider = func(ctx context.Context, unhashedToken string) (*models.UserToken, error) {
				return &models.UserToken{
					UserId:        12,
					UnhashedToken: unhashedToken,
				}, nil
			}

			sc.fakeReq("GET", "/").exec()

			Convey("Should init context with user info", func() {
				So(sc.context.IsSignedIn, ShouldBeTrue)
				So(sc.context.UserId, ShouldEqual, 12)
				So(sc.context.UserToken.UserId, ShouldEqual, 12)
				So(sc.context.UserToken.UnhashedToken, ShouldEqual, "token")
			})

			Convey("Should not set cookie", func() {
				So(sc.resp.Header().Get("Set-Cookie"), ShouldEqual, "")
			})
		})

		middlewareScenario(t, "Non-expired auth token in cookie which are being rotated", func(sc *scenarioContext) {
			sc.withTokenSessionCookie("token")

			bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
				query.Result = &models.SignedInUser{OrgId: 2, UserId: 12}
				return nil
			})

			sc.userAuthTokenService.LookupTokenProvider = func(ctx context.Context, unhashedToken string) (*models.UserToken, error) {
				return &models.UserToken{
					UserId:        12,
					UnhashedToken: "",
				}, nil
			}

			sc.userAuthTokenService.TryRotateTokenProvider = func(ctx context.Context, userToken *models.UserToken, clientIP, userAgent string) (bool, error) {
				userToken.UnhashedToken = "rotated"
				return true, nil
			}

			maxAgeHours := (time.Duration(setting.LoginMaxLifetimeDays) * 24 * time.Hour)
			maxAge := (maxAgeHours + time.Hour).Seconds()

			sameSitePolicies := []http.SameSite{
				http.SameSiteDefaultMode,
				http.SameSiteLaxMode,
				http.SameSiteStrictMode,
			}
			for _, sameSitePolicy := range sameSitePolicies {
				setting.CookieSameSite = sameSitePolicy
				expectedCookie := &http.Cookie{
					Name:     setting.LoginCookieName,
					Value:    "rotated",
					Path:     setting.AppSubUrl + "/",
					HttpOnly: true,
					MaxAge:   int(maxAge),
					Secure:   setting.CookieSecure,
				}
				if sameSitePolicy != http.SameSiteDefaultMode {
					expectedCookie.SameSite = sameSitePolicy
				}

				sc.fakeReq("GET", "/").exec()

				Convey(fmt.Sprintf("Should init context with user info and setting.SameSite=%v", sameSitePolicy), func() {
					So(sc.context.IsSignedIn, ShouldBeTrue)
					So(sc.context.UserId, ShouldEqual, 12)
					So(sc.context.UserToken.UserId, ShouldEqual, 12)
					So(sc.context.UserToken.UnhashedToken, ShouldEqual, "rotated")
				})

				Convey(fmt.Sprintf("Should set cookie with setting.SameSite=%v", sameSitePolicy), func() {
					So(sc.resp.Header().Get("Set-Cookie"), ShouldEqual, expectedCookie.String())
				})
			}
		})

		middlewareScenario(t, "Invalid/expired auth token in cookie", func(sc *scenarioContext) {
			sc.withTokenSessionCookie("token")

			sc.userAuthTokenService.LookupTokenProvider = func(ctx context.Context, unhashedToken string) (*models.UserToken, error) {
				return nil, models.ErrUserTokenNotFound
			}

			sc.fakeReq("GET", "/").exec()

			Convey("Should not init context with user info", func() {
				So(sc.context.IsSignedIn, ShouldBeFalse)
				So(sc.context.UserId, ShouldEqual, 0)
				So(sc.context.UserToken, ShouldBeNil)
			})
		})

		middlewareScenario(t, "When anonymous access is enabled", func(sc *scenarioContext) {
			setting.AnonymousEnabled = true
			setting.AnonymousOrgName = "test"
			setting.AnonymousOrgRole = string(models.ROLE_EDITOR)

			bus.AddHandler("test", func(query *models.GetOrgByNameQuery) error {
				So(query.Name, ShouldEqual, "test")

				query.Result = &models.Org{Id: 2, Name: "test"}
				return nil
			})

			sc.fakeReq("GET", "/").exec()

			Convey("Should init context with org info", func() {
				So(sc.context.UserId, ShouldEqual, 0)
				So(sc.context.OrgId, ShouldEqual, 2)
				So(sc.context.OrgRole, ShouldEqual, models.ROLE_EDITOR)
			})

			Convey("context signed in should be false", func() {
				So(sc.context.IsSignedIn, ShouldBeFalse)
			})
		})

		Convey("auth_proxy", func() {
			setting.AuthProxyEnabled = true
			setting.AuthProxyWhitelist = ""
			setting.AuthProxyAutoSignUp = true
			setting.LDAPEnabled = true
			setting.AuthProxyHeaderName = "X-WEBAUTH-USER"
			setting.AuthProxyHeaderProperty = "username"
			setting.AuthProxyHeaders = map[string]string{"Groups": "X-WEBAUTH-GROUPS"}
			name := "markelog"
			group := "smartems-core-team"

			middlewareScenario(t, "Should not sync the user if it's in the cache", func(sc *scenarioContext) {
				bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
					query.Result = &models.SignedInUser{OrgId: 4, UserId: query.UserId}
					return nil
				})

				key := fmt.Sprintf(authproxy.CachePrefix, base32.StdEncoding.EncodeToString([]byte(name+"-"+group)))
				err := sc.remoteCacheService.Set(key, int64(33), 0)
				So(err, ShouldBeNil)
				sc.fakeReq("GET", "/")

				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.req.Header.Add("X-WEBAUTH-GROUPS", group)
				sc.exec()

				Convey("Should init user via cache", func() {
					So(sc.context.IsSignedIn, ShouldBeTrue)
					So(sc.context.UserId, ShouldEqual, 33)
					So(sc.context.OrgId, ShouldEqual, 4)
				})
			})

			middlewareScenario(t, "Should respect auto signup option", func(sc *scenarioContext) {
				setting.LDAPEnabled = false
				setting.AuthProxyAutoSignUp = false
				var actualAuthProxyAutoSignUp *bool = nil

				bus.AddHandler("test", func(cmd *models.UpsertUserCommand) error {
					actualAuthProxyAutoSignUp = &cmd.SignupAllowed
					return login.ErrInvalidCredentials
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.exec()

				assert.False(t, *actualAuthProxyAutoSignUp)
				assert.Equal(t, sc.resp.Code, 407)
				assert.Nil(t, sc.context)
			})

			middlewareScenario(t, "Should create an user from a header", func(sc *scenarioContext) {
				setting.LDAPEnabled = false
				setting.AuthProxyAutoSignUp = true

				bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
					if query.UserId > 0 {
						query.Result = &models.SignedInUser{OrgId: 4, UserId: 33}
						return nil
					}
					return models.ErrUserNotFound
				})

				bus.AddHandler("test", func(cmd *models.UpsertUserCommand) error {
					cmd.Result = &models.User{Id: 33}
					return nil
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.exec()

				Convey("Should create user from header info", func() {
					So(sc.context.IsSignedIn, ShouldBeTrue)
					So(sc.context.UserId, ShouldEqual, 33)
					So(sc.context.OrgId, ShouldEqual, 4)
				})
			})

			middlewareScenario(t, "Should get an existing user from header", func(sc *scenarioContext) {
				setting.LDAPEnabled = false

				bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
					query.Result = &models.SignedInUser{OrgId: 2, UserId: 12}
					return nil
				})

				bus.AddHandler("test", func(cmd *models.UpsertUserCommand) error {
					cmd.Result = &models.User{Id: 12}
					return nil
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.exec()

				Convey("Should init context with user info", func() {
					So(sc.context.IsSignedIn, ShouldBeTrue)
					So(sc.context.UserId, ShouldEqual, 12)
					So(sc.context.OrgId, ShouldEqual, 2)
				})
			})

			middlewareScenario(t, "Should allow the request from whitelist IP", func(sc *scenarioContext) {
				setting.AuthProxyWhitelist = "192.168.1.0/24, 2001::0/120"
				setting.LDAPEnabled = false

				bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
					query.Result = &models.SignedInUser{OrgId: 4, UserId: 33}
					return nil
				})

				bus.AddHandler("test", func(cmd *models.UpsertUserCommand) error {
					cmd.Result = &models.User{Id: 33}
					return nil
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.req.RemoteAddr = "[2001::23]:12345"
				sc.exec()

				Convey("Should init context with user info", func() {
					So(sc.context.IsSignedIn, ShouldBeTrue)
					So(sc.context.UserId, ShouldEqual, 33)
					So(sc.context.OrgId, ShouldEqual, 4)
				})
			})

			middlewareScenario(t, "Should not allow the request from whitelist IP", func(sc *scenarioContext) {
				setting.AuthProxyWhitelist = "8.8.8.8"
				setting.LDAPEnabled = false

				bus.AddHandler("test", func(query *models.GetSignedInUserQuery) error {
					query.Result = &models.SignedInUser{OrgId: 4, UserId: 33}
					return nil
				})

				bus.AddHandler("test", func(cmd *models.UpsertUserCommand) error {
					cmd.Result = &models.User{Id: 33}
					return nil
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.req.RemoteAddr = "[2001::23]:12345"
				sc.exec()

				Convey("Should return 407 status code", func() {
					So(sc.resp.Code, ShouldEqual, 407)
					So(sc.context, ShouldBeNil)
				})
			})

			middlewareScenario(t, "Should return 407 status code if LDAP says no", func(sc *scenarioContext) {
				bus.AddHandler("LDAP", func(cmd *models.UpsertUserCommand) error {
					return errors.New("Do not add user")
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.exec()

				Convey("Should return 407 status code", func() {
					So(sc.resp.Code, ShouldEqual, 407)
					So(sc.context, ShouldBeNil)
				})
			})

			middlewareScenario(t, "Should return 407 status code if there is cache mishap", func(sc *scenarioContext) {
				bus.AddHandler("Do not have the user", func(query *models.GetSignedInUserQuery) error {
					return errors.New("Do not add user")
				})

				sc.fakeReq("GET", "/")
				sc.req.Header.Add(setting.AuthProxyHeaderName, name)
				sc.exec()

				Convey("Should return 407 status code", func() {
					So(sc.resp.Code, ShouldEqual, 407)
					So(sc.context, ShouldBeNil)
				})
			})
		})
	})
}

func middlewareScenario(t *testing.T, desc string, fn scenarioFunc) {
	Convey(desc, func() {
		defer bus.ClearBusHandlers()

		setting.LoginCookieName = "smartems_session"
		setting.LoginMaxLifetimeDays = 30

		sc := &scenarioContext{}

		viewsPath, _ := filepath.Abs("../../public/views")

		sc.m = macaron.New()
		sc.m.Use(AddDefaultResponseHeaders())
		sc.m.Use(macaron.Renderer(macaron.RenderOptions{
			Directory: viewsPath,
			Delims:    macaron.Delims{Left: "[[", Right: "]]"},
		}))

		sc.userAuthTokenService = auth.NewFakeUserAuthTokenService()
		sc.remoteCacheService = remotecache.NewFakeStore(t)

		sc.m.Use(GetContextHandler(sc.userAuthTokenService, sc.remoteCacheService))

		sc.m.Use(OrgRedirect())

		sc.defaultHandler = func(c *models.ReqContext) {
			sc.context = c
			if sc.handlerFunc != nil {
				sc.handlerFunc(sc.context)
			}
		}

		sc.m.Get("/", sc.defaultHandler)

		fn(sc)
	})
}
