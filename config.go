package main

import (
	"flag"
	"log"
	"time"
)

var (
	httpAddress = flag.String("http-address", ":3333", "HTTP listener address")
	httpBye     = flag.String("http-bye", "This is a proxy server. Does not respond to non-proxy requests.", "")
	httpRealm   = flag.String("http-realm", "Proxy", "HTTP authentication realm")

	duoIKey = flag.String("duo-ikey", "DIKT827ALIAIZM9N31PH", "integration key provided by Duo")
	duoSKey = flag.String("duo-skey", "Z1BkT2OriakMCxEJg3ShPfxCCQpOX3o0ugEuoT1L", "secret key provided by Duo")
	duoHost = flag.String("duo-host", "api-bc24bc24.duosecurity.com", "hostname of Duo API")
	duoUA   = flag.String("duo-ua", "abcdefgh", "userAgent reported to Duo API")

	duoSkip = flag.Bool("duo-skip", false, "bypass Duo verifications")
	duoTTL  = flag.Duration("duo-ttl", 24*time.Hour, "how long to remember calls to Duo")

	githubBase = flag.String("github-base", "https://api.git.mycompany.com", "GitHub Base URL")
	githubTTL  = flag.Duration("github-ttl", 15*time.Minute, "how long to remember calls to GitHub")
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	flag.Parse()
}
