package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/cloudflare/golibs/lrucache"
)

var (
	githubCache = lrucache.NewMultiLRUCache(8, 64)
)

func githubAuth(token string) string {
	if v, ex := githubCache.Get(token); ex {
		if v, ok := v.(string); ok {
			return v
		}
	}

	resp, err := http.Get(*githubBase + "/user?access_token=" + token)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		githubCache.Set(token, "", time.Now().Add(*githubTTL))
		return ""
	}

	v := &githubUser{}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		log.Println(err)
		return ""
	}
	githubCache.Set(token, v.Login, time.Now().Add(*githubTTL))
	return v.Login
}

type githubUser struct {
	Login string
	Email string
}
