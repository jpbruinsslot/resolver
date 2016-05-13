package main

import "github.com/codegangsta/negroni"

func ServerSetup() *negroni.Negroni {
	// create a new router (see routes.go)
	router := NewRouter()

	// middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(AuthenticationMiddleware),
	)

	// let negroni use our mux router
	n.UseHandler(router)

	return n
}
