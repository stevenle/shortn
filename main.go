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

package main

import (
	"log"
	"net/http"

	// TODO(stevenle): Implement.
	// "github.com/stevenle/shortn/handlers"
	"github.com/stevenle/web"
)

func pingHandler(ctx *web.Context) {
	web.WriteResponseString(ctx, "pong\n")
}

func main() {
	router := web.NewRouter()
	router.HandleFunc("/ping", pingHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}