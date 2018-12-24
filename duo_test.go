package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	duoapi "github.com/duosecurity/duo_api_golang"
	"github.com/duosecurity/duo_api_golang/authapi"
	"github.com/stretchr/testify/assert"
)

var (
	testDuoFailMessage       = "Invalid request parameters"
	testDuoFailMessageDetail = "username"
)

func duoAuthDeny(factor string, options ...func(*url.Values)) (*authapi.AuthResult, error) {
	r := &authapi.AuthResult{StatResult: duoapi.StatResult{Stat: "OK"}}
	r.Response.Result = "deny"
	r.Response.Status = "deny"
	r.Response.Status_Msg = "Login request denied."
	return r, nil
}

func duoAuthError(factor string, options ...func(*url.Values)) (*authapi.AuthResult, error) {
	return nil, http.ErrHijacked
}

func duoAuthFail(factor string, options ...func(*url.Values)) (*authapi.AuthResult, error) {
	r := &authapi.AuthResult{StatResult: duoapi.StatResult{Stat: "FAIL"}}
	r.StatResult.Message = &testDuoFailMessage
	r.StatResult.Message_Detail = &testDuoFailMessageDetail
	return r, nil
}

func duoAuthOK(factor string, options ...func(*url.Values)) (*authapi.AuthResult, error) {
	r := &authapi.AuthResult{StatResult: duoapi.StatResult{Stat: "OK"}}
	r.Response.Result = "allow"
	r.Response.Status = "allow"
	r.Response.Status_Msg = "Success. Logging you in..."
	return r, nil
}

func TestDuoDeny(t *testing.T) {
	duoAuthAuth = duoAuthDeny

	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 407, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, statusProxyAuthReq, string(body))
}

func TestDuoError(t *testing.T) {
	duoAuthAuth = duoAuthError

	request, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 407, response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, statusProxyAuthReq, string(body))
}

func TestDuoFail(t *testing.T) {
	duoAuthAuth = duoAuthFail

	request, err := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	assert.Nil(t, err)
	response, err := clientSuccess.Do(request)
	assert.Error(t, err)
	assert.Nil(t, response)
}
