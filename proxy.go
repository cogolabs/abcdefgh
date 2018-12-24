package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

var proxy = func() *goproxy.ProxyHttpServer {
	this := goproxy.NewProxyHttpServer()
	this.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, *httpBye, 400)
	})

	this.OnRequest().Do(
		goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			user, token, _ := getProxyAuth(ctx.Req)
			if user == "" || token == "" {
				return nil, newUnauthorizedResponse(req, *httpRealm)
			}
			if githubAuth(token) != user {
				return nil, newUnauthorizedResponse(req, *httpRealm)
			}
			if err := duo(user, req.Host, req); err != nil {
				log.Println(user, err)
				return nil, newUnauthorizedResponse(req, *httpRealm)
			}
			return req, nil
		}),
	)

	this.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		user, token, _ := getProxyAuth(ctx.Req)
		if user == "" || token == "" {
			return goproxy.RejectConnect, host
		}
		if githubAuth(token) != user {
			return goproxy.RejectConnect, host
		}
		if err := duo(user, host, ctx.Req); err != nil {
			log.Println(user, err)
			return goproxy.RejectConnect, host
		}
		return goproxy.OkConnect, host
	})

	return this
}()

func main() {
	log.Fatalln(http.ListenAndServe(*httpAddress, proxy))
}
