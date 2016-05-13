package main

import "net/http"

// AuthenticationMiddleware will check if authentication is correct on
// every request.
func AuthenticationMiddleware(rw http.ResponseWriter,
	r *http.Request, next http.HandlerFunc) {

	if CheckAuth(r) {
		next(rw, r)
	} else {
		rw.Header().Set("WWW-Authenticate", "This page is protected")
		rw.WriteHeader(401)
		rw.Write([]byte("401 Unauthorized\n"))
	}

}
