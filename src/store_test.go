package main

import (
	"io/ioutil"
	"testing"
)

func TestLoadStore(t *testing.T) {
	// create a test store file
	f, err := ioutil.TempFile("", "datastore.json")
	if err != nil {
		t.Fatal(err)
	}

	// set contents
	payload := `{}`
	ioutil.WriteFile(
		f.Name(),
		[]byte(payload),
		0644,
	)

	LoadStore(f.Name())
}

func TestSaveStore(t *testing.T) {
	// create a test store file
	f, err := ioutil.TempFile("", "datastore.json")
	if err != nil {
		t.Fatal(err)
	}

	// set contents
	payload := `{}`
	ioutil.WriteFile(
		f.Name(),
		[]byte(payload),
		0644,
	)

	LoadStore(f.Name())

	// create a new entry
	store := *LoadedStore
	store["test-asset"] = "test-asset-hash"

	SaveStore(f.Name())
}
