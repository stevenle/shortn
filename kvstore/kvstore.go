// Copyright 2014 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kvstore is a simple in-memory data store that persists to the
// filesystem.
package kvstore

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/stevenle/logging"
)

var store map[string]string

const storePath = "/var/shortn/kvstore.json"

func init() {
	store = make(map[string]string)
	_, err := os.Stat(storePath)
	if os.IsNotExist(err) {
		return
	}

	logging.Infof("initializing kvstore from %s", storePath)

	data, err := ioutil.ReadFile(storePath)
	if err != nil {
		logging.Fatalf("failed to load kvstore: %s", err)
	}
	err = json.Unmarshal(data, &store)
	if err != nil {
		logging.Fatalf("failed to parse kvstore: %s", err)
	}
}

func Set(key, value string) bool {
	_, found := store[key]
	store[key] = value
	logging.Infof("set \"%s\" to \"%s\"", key, value)
	go saveToFile()
	return !found
}

func Get(key string) string {
	return store[key]
}

// Close saves the store to the filesystem for persistence.
func Close() {
	saveToFile()
}

func saveToFile() {
	if len(store) == 0 {
		return
	}

	file, err := os.Create(storePath)
	if err != nil {
		logging.Errorf("failed to open file %s: %s", storePath, err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(store)
	if err != nil {
		logging.Errorf("failed to save kvstore, err: %s", err)
		return
	}
	logging.Infof("saved kvstore to %s", storePath)
}
