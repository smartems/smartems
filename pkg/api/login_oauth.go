package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/infra/log"
	"github.com/smartems/smartems/pkg/infra/metrics"
	"github.com/smartems/smartems/pkg/login"
	"github.com/smartems/smartems/pkg/login/social"
	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/setting"
)

var (
	oauthLogger          = log.New("oauth")
	OauthStateCookieName = "oauth_state"
)

func GenStateString() (string, error) {
	rnd := make([]byte, 32)
	if _, err := rand.Read(rnd); err != nil {
		oauthLogger.Error("failed to generate state string", "err", err)
		return "", err
	}
	return base64.URLEncoding.EncodeToString(rnd), nil
}

func (hs *HTTPServer) OAuthLogin(ctx *m.ReqContext) {
	if setting.OAuthService == nil {
		ctx.Handle(404, "OAuth not enabled", nil)
		return
	}

	name := ctx.Params(":name")
	connect, ok := social.SocialMap[name]
	if !ok {
		ctx.Handle(404, fmt.Sprintf("No OAuth with name %s configured", name), nil)
		return
	}

	errorParam := ctx.Query("error")
	if errorParam != "" {
		errorDesc := ctx.Query("error_description")
		oauthLogger.Error("failed to login ", "error", errorParam, "errorDesc", errorDesc)
		hs.redirectWithError(ctx, login.ErrProviderDeniedRequest, "error", errorParam, "errorDesc", errorDesc)
		return
	}

	code := ctx.Query("code")
	if code == "" {
		state, err := GenStateString()
		if err != nil {
			ctx.Logger.Error("Generating state string failed", "err", err)
			ctx.Handle(500, "An internal error occurred", nil)
			return
		}

		hashedState := hashStatecode(state, setting.OAuthService.OAuthInfos[name].ClientSecret)
		hs.writeCookie(ctx.Resp, OauthStateCookieName, hashedState, 60, hs.Cfg.CookieSameSite)
		if setting.OAuthService.OAuthInfos[name].HostedDomain == "" {
			ctx.Redirect(connect.AuthCodeURL(state, oauth2.AccessTypeOnline))
		} else {
			ctx.Redirect(connect.AuthCodeURL(state, oauth2.SetAuthURLParam("hd", setting.OAuthService.OAuthInfos[name].HostedDomain), oauth2.AccessTypeOnline))
		}
		return
	}

	cookieState := ctx.GetCookie(OauthStateCookieName)

	// delete cookie
	ctx.Resp.Header().Del("Set-Cookie")
	hs.deleteCookie(ctx.Resp, OauthStateCookieName, hs.Cfg.CookieSameSite)

	if cookieState == "" {
		ctx.Handle(500, "login.OAuthLogin(missing saved state)", nil)
		return
	}

	queryState := hashStatecode(ctx.Query("state"), setting.OAuthService.OAuthInfos[name].ClientSecret)
	oauthLogger.Info("state check", "queryState", queryState, "cookieState", cookieState)
	if cookieState != queryState {
		ctx.Handle(500, "login.OAuthLogin(state mismatch)", nil)
		return
	}

	// handle call back
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: setting.OAuthService.OAuthInfos[name].TlsSkipVerify,
		},
	}
	oauthClient := &http.Client{
		Transport: tr,
	}

	if setting.OAuthService.OAuthInfos[name].TlsClientCert != "" || setting.OAuthService.OAuthInfos[name].TlsClientKey != "" {
		cert, err := tls.LoadX509KeyPair(setting.OAuthService.OAuthInfos[name].TlsClientCert, setting.OAuthService.OAuthInfos[name].TlsClientKey)
		if err != nil {
			ctx.Logger.Error("Failed to setup TlsClientCert", "oauth", name, "error", err)
			ctx.Handle(500, "login.OAuthLogin(Failed to setup TlsClientCert)", nil)
			return
		}

		tr.TLSClientConfig.Certificates = append(tr.TLSClientConfig.Certificates, cert)
	}

	if setting.OAuthService.OAuthInfos[name].TlsClientCa != "" {
		caCert, err := ioutil.ReadFile(setting.OAuthService.OAuthInfos[name].TlsClientCa)
		if err != nil {
			ctx.Logger.Error("Failed to setup TlsClientCa", "oauth", name, "error", err)
			ctx.Handle(500, "login.OAuthLogin(Failed to setup TlsClientCa)", nil)
			return
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tr.TLSClientConfig.RootCAs = caCertPool
	}

	oauthCtx := context.WithValue(context.Background(), oauth2.HTTPClient, oauthClient)

	// get token from provider
	token, err := connect.Exchange(oauthCtx, code)
	if err != nil {
		ctx.Handle(500, "login.OAuthLogin(NewTransportWithCode)", err)
		return
	}
	// token.TokenType was defaulting to "bearer", which is out of spec, so we explicitly set to "Bearer"
	token.TokenType = "Bearer"

	oauthLogger.Debug("OAuthLogin Got token", "token", token)

	// set up oauth2 client
	client := connect.Client(oauthCtx, token)

	// get user info
	userInfo, err := connect.UserInfo(client, token)
	if err != nil {
		if sErr, ok := err.(*social.Error); ok {
			hs.redirectWithError(ctx, sErr)
		} else {
			ctx.Handle(500, fmt.Sprintf("login.OAuthLogin(get info from %s)", name), err)
		}
		return
	}

	oauthLogger.Debug("OAuthLogin got user info", "userInfo", userInfo)

	// validate that we got at least an email address
	if userInfo.Email == "" {
		hs.redirectWithError(ctx, login.ErrNoEmail)
		return
	}

	// validate that the email is allowed to login to smartems
	if !connect.IsEmailAllowed(userInfo.Email) {
		hs.redirectWithError(ctx, login.ErrEmailNotAllowed)
		return
	}

	extUser := &m.ExternalUserInfo{
		AuthModule: "oauth_" + name,
		OAuthToken: token,
		AuthId:     userInfo.Id,
		Name:       userInfo.Name,
		Login:      userInfo.Login,
		Email:      userInfo.Email,
		OrgRoles:   map[int64]m.RoleType{},
		Groups:     userInfo.Groups,
	}

	if userInfo.Role != "" {
		rt := m.RoleType(userInfo.Role)
		if rt.IsValid() {
			extUser.OrgRoles[1] = rt
		}
	}

	// add/update user in smartems
	cmd := &m.UpsertUserCommand{
		ReqContext:    ctx,
		ExternalUser:  extUser,
		SignupAllowed: connect.IsSignupAllowed(),
	}

	err = bus.Dispatch(cmd)
	if err != nil {
		hs.redirectWithError(ctx, err)
		return
	}

	// Do not expose disabled status,
	// just show incorrect user credentials error (see #17947)
	if cmd.Result.IsDisabled {
		oauthLogger.Warn("User is disabled", "user", cmd.Result.Login)
		hs.redirectWithError(ctx, login.ErrInvalidCredentials)
		return
	}

	// login
	hs.loginUserWithUser(cmd.Result, ctx)

	metrics.MApiLoginOAuth.Inc()

	if redirectTo, _ := url.QueryUnescape(ctx.GetCookie("redirect_to")); len(redirectTo) > 0 {
		ctx.SetCookie("redirect_to", "", -1, setting.AppSubUrl+"/")
		ctx.Redirect(redirectTo)
		return
	}

	ctx.Redirect(setting.AppSubUrl + "/")
}

func (hs *HTTPServer) deleteCookie(w http.ResponseWriter, name string, sameSite http.SameSite) {
	hs.writeCookie(w, name, "", -1, sameSite)
}

func (hs *HTTPServer) writeCookie(w http.ResponseWriter, name string, value string, maxAge int, sameSite http.SameSite) {
	cookie := http.Cookie{
		Name:     name,
		MaxAge:   maxAge,
		Value:    value,
		HttpOnly: true,
		Path:     setting.AppSubUrl + "/",
		Secure:   hs.Cfg.CookieSecure,
	}
	if sameSite != http.SameSiteDefaultMode {
		cookie.SameSite = sameSite
	}
	http.SetCookie(w, &cookie)
}

func hashStatecode(code, seed string) string {
	hashBytes := sha256.Sum256([]byte(code + setting.SecretKey + seed))
	return hex.EncodeToString(hashBytes[:])
}

func (hs *HTTPServer) redirectWithError(ctx *m.ReqContext, err error, v ...interface{}) {
	ctx.Logger.Error(err.Error(), v...)
	if err := hs.trySetEncryptedCookie(ctx, LoginErrorCookieName, err.Error(), 60); err != nil {
		oauthLogger.Error("Failed to set encrypted cookie", "err", err)
	}

	ctx.Redirect(setting.AppSubUrl + "/login")
}
