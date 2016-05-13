package main

import "github.com/gorilla/mux"

// NewRouter will return a *mux.Router. Here we'll define
// the different routes and associated handlers
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	sub := router.PathPrefix("/assets").Subrouter()

	sub.HandleFunc("/", GetAsset).
		Methods("GET").
		Queries("q", "{q}")

	sub.HandleFunc("/", GetAssets).
		Methods("GET")

	sub.HandleFunc("/", PostAssets).
		Methods("POST")

	return router
}
