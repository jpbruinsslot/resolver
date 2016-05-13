package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// GET /assets/
//
// GetAssets will return all the available assets present in the
// datastore.
func GetAssets(w http.ResponseWriter, r *http.Request) {
	// load the store
	store := *LoadedStore

	// return data
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(store); err != nil {
		log.Fatal(err)
	}
}

// POST /assets/
//
// PostAssets will update the datastore with the new assets with their
// hashed counterparts.
func PostAssets(w http.ResponseWriter, r *http.Request) {
	// get body, and limit what we can receive
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}

	// close the body
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}

	// load the body in Store s
	var s Store
	if err := json.Unmarshal(body, &s); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	// load the LoadedStore and update it with the received Store s
	store := *LoadedStore
	for k, v := range s {
		store[k] = v
	}

	// persist the store
	SaveStore(ResolverStore)

	// return success message
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(store); err != nil {
		log.Fatal(err)
	}
}

// GET /assets/{key}
//
// GetAsset will return the hashed value of the requested asset.
// This uses the `?q=` query param because it isn't possible to add full
// hashes as a key.
func GetAsset(w http.ResponseWriter, r *http.Request) {
	// get arguments
	vars := mux.Vars(r)
	key := vars["q"]

	// load the store
	store := *LoadedStore

	// get the value with the key
	value, ok := store[key]

	// check if key was present
	var statusCode int
	if ok {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}

	// return data
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Fatal(err)
	}
}
