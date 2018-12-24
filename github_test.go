package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	githubTestLogin = "user1"
	githubTestToken = "932928c0a4edf9878ee0257a1d8f4d06adaaffee"

	testGithubServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("access_token") == "invalid" {
			io.WriteString(w, "{")
			return
		}
		if r.FormValue("access_token") != githubTestToken {
			w.WriteHeader(403)
			return
		}
		json.NewEncoder(w).Encode(githubUser{Login: githubTestLogin})
	}))
)

func init() {
	*githubBase = testGithubServer.URL
}

func TestGithubErrors(t *testing.T) {
	old := *githubBase
	*githubBase = "https://foo.bar?"
	assert.Equal(t, "", githubAuth("error"))

	*githubBase = old
	assert.Equal(t, "", githubAuth("invalid"))
}

func TestGithubSuccess(t *testing.T) {
	assert.Equal(t, githubTestLogin, githubAuth(githubTestToken))
	assert.Equal(t, githubTestLogin, githubAuth(githubTestToken))
}
