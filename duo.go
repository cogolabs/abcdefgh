package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/golibs/lrucache"
	duoapi "github.com/duosecurity/duo_api_golang"
	"github.com/duosecurity/duo_api_golang/authapi"
)

var (
	duoAuthAuth func(factor string, options ...func(*url.Values)) (*authapi.AuthResult, error)

	duoCache = lrucache.NewMultiLRUCache(8, 64)
)

func init() {
	go func() {
		for {
			duoCache.Expire()
			time.Sleep(time.Hour)
		}
	}()

	duoAuth := authapi.NewAuthApi(*duoapi.NewDuoApi(*duoIKey, *duoSKey, *duoHost, *duoUA))
	duoAuthAuth = duoAuth.Auth
}

func duo(user string, host string, req *http.Request) error {
	if *duoSkip {
		return nil
	}

	clientIP := getClientIP(req)
	cacheKey := clientIP + user

	_, cached := duoCache.GetQuiet(cacheKey)
	if cached {
		return nil
	}

	vals := []func(*url.Values){
		authapi.AuthDevice("auto"),
		authapi.AuthUsername(user),
		authapi.AuthIpAddr(clientIP),
		authapi.AuthType(host),
		authapi.AuthPushinfo("Host=" + host + "&UserAgent=" + req.UserAgent()),
	}

	res, err := duoAuthAuth("push", vals...)
	if err != nil {
		return fmt.Errorf("[duo] %s", err.Error())
	}
	if res.Stat == "FAIL" {
		return fmt.Errorf("[duo] %s: %s", *res.StatResult.Message, *res.StatResult.Message_Detail)
	}
	if res.Response.Status == "allow" {
		duoCache.Set(cacheKey, 1, time.Now().Add(*duoTTL))
		return nil
	}
	return fmt.Errorf("[duo] %s", res.Response.Status_Msg)
}
