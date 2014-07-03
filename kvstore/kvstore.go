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

package kvstore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var store map[string]string

const storePath = "/var/tmp/shortn.json"

func init() {
	store = make(map[string]string)
	_, err := os.Stat(storePath)
	if os.IsNotExist(err) {
		return
	}

	log.Printf("Initializing from %q", storePath)

	data, err := ioutil.ReadFile(storePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &store)
	if err != nil {
		log.Fatal(err)
	}
}

func Set(key, value string) bool {
	_, found := store[key]
	store[key] = value
	return !found
}

func Get(key string) string {
	return store[key]
}

// Close saves the store to the filesystem for persistence.
func Close() {
	if len(store) == 0 {
		return
	}

	file, err := os.Create(storePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(store)
}
