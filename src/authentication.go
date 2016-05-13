package main

import (
	"net/http"

	auth "github.com/abbot/go-http-auth"
)

var Authenticator = LoadAuthenticator()

// CheckAuth will check given a *http.Request if the header contains
// the correct user name and key.
func CheckAuth(r *http.Request) bool {
	if Authenticator.CheckAuth(r) != "" {
		return true
	} else {
		return false
	}
}

// LoadAuthenticator will return an authenticator which we can
// use to authenticate our request in.
func LoadAuthenticator() *auth.BasicAuth {
	return auth.NewBasicAuthenticator("Basic Realm", Secret)
}

// Secret will return the key (when a username is present) for
// basic http authentication. The user and key are defined by
// using environment variables
func Secret(user, realm string) string {
	if user == ResolverUser {
		return ResolverKey
	}
	return ""
}
