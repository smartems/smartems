package api

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/setting"
	"github.com/smartems/smartems/pkg/util"
)

var smartemsComProxyTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

func ReverseProxyGnetReq(proxyPath string) *httputil.ReverseProxy {
	url, _ := url.Parse(setting.GrafanaComUrl)

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host

		req.URL.Path = util.JoinURLFragments(url.Path+"/api", proxyPath)

		// clear cookie headers
		req.Header.Del("Cookie")
		req.Header.Del("Set-Cookie")
		req.Header.Del("Authorization")
	}

	return &httputil.ReverseProxy{Director: director}
}

func ProxyGnetRequest(c *m.ReqContext) {
	proxyPath := c.Params("*")
	proxy := ReverseProxyGnetReq(proxyPath)
	proxy.Transport = smartemsComProxyTransport
	proxy.ServeHTTP(c.Resp, c.Req.Request)
	c.Resp.Header().Del("Set-Cookie")
}
