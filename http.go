package main

import (
	"encoding/base64"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const (
	headerProxyAuth = "Proxy-Authorization"
	headerXFF       = "X-Forwarded-For"

	statusProxyAuthReq = "407 Proxy Authentication Required"
)

func newUnauthorizedResponse(req *http.Request, realm string) *http.Response {
	return &http.Response{
		StatusCode: 407,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Header: http.Header{
			"Proxy-Authenticate": []string{"Basic realm=" + realm},
			"Proxy-Connection":   []string{"close"},
		},
		Body:          ioutil.NopCloser(strings.NewReader(statusProxyAuthReq)),
		ContentLength: int64(len(statusProxyAuthReq)),
	}
}

func getProxyAuth(req *http.Request) (username, password string, ok bool) {
	auth := req.Header.Get(headerProxyAuth)
	if auth == "" {
		return
	}

	if strings.HasPrefix(strings.ToLower(auth), "basic ") {
		c, err := base64.StdEncoding.DecodeString(auth[6:])
		if err == nil {
			cs := string(c)
			s := strings.IndexByte(cs, ':')
			if s >= 0 {
				return cs[:s], cs[s+1:], true
			}
		}
	}
	return
}

func getClientIP(req *http.Request) string {
	field := strings.Split(req.Header.Get(headerXFF), ",")
	if field[0] != "" {
		return field[0]
	}
	x, _, z := net.SplitHostPort(req.RemoteAddr)
	if z == nil {
		return x
	}
	return ""
}
