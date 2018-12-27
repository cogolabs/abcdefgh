[![Build Status](https://travis-ci.org/cogolabs/abcdefgh.svg?branch=master)](https://travis-ci.org/cogolabs/abcdefgh)
[![codecov](https://codecov.io/gh/cogolabs/abcdefgh/branch/master/graph/badge.svg)](https://codecov.io/gh/cogolabs/abcdefgh)
[![Go Report Card](https://goreportcard.com/badge/github.com/cogolabs/abcdefgh)](https://goreportcard.com/report/github.com/cogolabs/abcdefgh)
[![Docker Build](https://images.microbadger.com/badges/image/cogolabs/abcdefgh.svg)](https://hub.docker.com/r/cogolabs/abcdefgh/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# abcdefgh
Another BeyondCorp proxy to enable access to systems and services beyond your perimeter network. Inspired by Google BeyondCorp research: https://research.google.com/pubs/pub45728.html

## Features
- Explicit HTTP Proxy (http://)
- Explicit HTTP CONNECT (https://)
- Authenticate via GitHub OAuth2 Tokens
- Verify user identity with Duo two-factor authentication

Implicit HTTP Access Proxy available at https://github.com/cogolabs/transcend

## Install
```
$ docker pull cogolabs/abcdefgh
```
or:
```
$ go get -u -x github.com/cogolabs/abcdefgh
```
## Usage
```
$ docker run --rm -p 80:80 cogolabs/abcdefgh abcdefgh --help
Usage of ./abcdefgh:
  -duo-host string
    	hostname of Duo API (default "api-bc24bc24.duosecurity.com")
  -duo-ikey string
    	integration key provided by Duo (default "DIKT827ALIAIZM9N31PH")
  -duo-skey string
    	secret key provided by Duo (default "Z1BkT2OriakMCxEJg3ShPfxCCQpOX3o0ugEuoT1L")
  -duo-skip
    	disable Duo enforcement
  -duo-ttl duration
    	how long to remember calls to Duo (default 24h0m0s)
  -duo-ua string
    	userAgent reported to Duo API (default "abcdefgh")
  -github-base string
    	GitHub Base URL (default "https://api.git.mycompany.com")
  -github-ttl duration
    	how long to remember calls to GitHub (default 15m0s)
  -http-address string
    	HTTP listener address (default ":3333")
  -http-bye string
    	 (default "This is a proxy server. Does not respond to non-proxy requests.")
  -http-realm string
    	HTTP authentication realm (default "Proxy")
```
