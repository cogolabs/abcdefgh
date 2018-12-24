package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	clientAnon = &http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(testServer.URL)
		},
	}}
	clientInvalid = &http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			u, err := url.Parse(testServer.URL)
			u.User = url.UserPassword(githubTestLogin, "foobar")
			return u, err
		},
	}}
	clientSuccess = &http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			u, err := url.Parse(testServer.URL)
			u.User = url.UserPassword(githubTestLogin, githubTestToken)
			return u, err
		},
	}}

	testServer = httptest.NewServer(proxy)
)

func TestProxyAnonHTTP(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientAnon.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 407, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, statusProxyAuthReq, string(body))
}

func TestProxyAnonHTTPS(t *testing.T) {
	request, err := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientAnon.Do(request)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestProxyFailDirect(t *testing.T) {
	resp, err := http.Get(testServer.URL)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "400 Bad Request", resp.Status)
}

func TestProxyInvalidAuth(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	request.Header.Set("Proxy-Authorization", "token invalid")
	assert.Nil(t, err)
	response, err := clientInvalid.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 407, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, statusProxyAuthReq, string(body))
}

func TestProxyInvalidHTTP(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientInvalid.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 407, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, statusProxyAuthReq, string(body))
}

func TestProxyInvalidHTTPS(t *testing.T) {
	request, err := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientInvalid.Do(request)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestProxySuccess0(t *testing.T) {
	duoAuthAuth = duoAuthOK
}

func TestProxySuccess1(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	request.Header.Set("X-Forwarded-For", "127.127.127.127")
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\n  \"origin\"", strings.Split(string(body), ":")[0])
}

func TestProxySuccess2(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\n  \"origin\"", strings.Split(string(body), ":")[0])
}

func TestProxySuccess3(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\n  \"origin\"", strings.Split(string(body), ":")[0])
}

func TestProxySuccessFinal(t *testing.T) {
	duoAuthAuth = duoAuthDeny
	*duoSkip = true
}

func TestProxySuccessHTTPS(t *testing.T) {
	request, err := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\n  \"origin\"", strings.Split(string(body), ":")[0])
}
