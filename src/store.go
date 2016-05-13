package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

type Store map[string]string

// LoadStore will load the datastore file
func LoadStore(path string) *Store {
	s := new(Store)

	if _, err := os.Stat(path); err != nil {
		payload := fmt.Sprint("{}")
		err = ioutil.WriteFile(path, []byte(payload), 0755)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(file).Decode(&s); err != nil {
		log.Fatal(err)
	}

	return s
}

// SaveStore will save/update the datastore when new entries have
// been added
func SaveStore(path string) {
	// truncate existing file if it's present
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	store := *LoadedStore
	b, err := json.MarshalIndent(store, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}
}
