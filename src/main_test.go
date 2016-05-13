package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
)

var (
	Server *httptest.Server
)

func init() {
	// set necessary env variables
	os.Setenv("RESOLVER_USER", "admin")
	os.Setenv("RESOLVER_KEY", "$apr1$btdc3WXb$dcl1QeZUA.M6xVTjTKj5a/")

	// Spin up test server
	Server = httptest.NewServer(ServerSetup())
}

// define here the setUp and tearDown functions
func TestMain(m *testing.M) {
	// setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func tearDown() {
	os.Remove("./datastore.json")
}

// request will be a helper function to quickly make request against the
// test server
func request(method, url string, payload []byte, auth bool) *http.Response {
	// create new request
	reqURL := fmt.Sprintf("%s/%s", Server.URL, url)
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	// set request headers
	if auth {
		req.SetBasicAuth("admin", "password")
	}
	req.Header.Set("Content-Type", "application/json")

	// setup client with request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	return resp
}

// An authorized client should be able to GET `/assets/`
func TestGetAssetsAuthorized(t *testing.T) {
	resp := request("GET", "assets/", nil, true)

	if resp.StatusCode != http.StatusOK {
		t.Error("not able to GET /assets")
	}
}

// An unauthorized client shouldn't be able to GET `/assets/`
func TestGetAssetsUnauthorized(t *testing.T) {
	resp := request("GET", "assets/", nil, false)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Unauthorized request was allowed")
	}
}

// An authorized client should be able to POST `/assets/`
func TestPostAssetsAuthorized(t *testing.T) {
	var payload = []byte(`{"test-asset": "test-asset-hash"}`)
	resp := request("POST", "assets/", payload, true)

	if resp.StatusCode != http.StatusCreated {
		t.Error("not able to POST /assets")
	}
}

// An unauthorized client shouldn't be able to POST `/assets/`
func TestPostAssetsUnauthorized(t *testing.T) {
	var payload = []byte(`{"test-asset": "test-asset-hash"}`)
	resp := request("POST", "assets/", payload, false)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Unauthorized request was allowed")
	}
}

// An authorized client should be able to GET `/assets/{asset}`
func TestGetAssetAuthorized(t *testing.T) {
	// first create an asset
	request("POST", "assets/", []byte(`{"test-asset": "test-asset-hash"}`), true)

	// get the created asset
	resp := request("GET", "assets/test-asset", nil, true)

	if resp.StatusCode != http.StatusOK {
		t.Error("not able to GET /assets/{asset}")
	}
}

// An unauthorized client shouldn't be able to GET `/assets/{asset}`
func TestGetAssetUnauthorized(t *testing.T) {
	// first create an asset
	request("POST", "assets/", []byte(`{"test-asset": "test-asset-hash"}`), true)

	// get the created asset
	resp := request("GET", "assets/test-asset", nil, false)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Unauthorized request was allowed")
	}
}
