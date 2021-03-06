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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stevenle/logging"
	"github.com/stevenle/shortn/handlers"
	"github.com/stevenle/shortn/kvstore"
	"github.com/stevenle/web"
)

func main() {
	defer kvstore.Close()

	// Start the server in a goroutine listening on port 8080.
	go func() {
		router := web.NewRouter()
		router.HandleFunc("/_/ping", handlers.PingHandler)
		router.HandleFunc("/*id", handlers.GoHandler)

		s := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
		logging.Infof("starting server on port 8080")
		if err := s.ListenAndServe(); err != nil {
			logging.Fatalf("failed to start server: %s", err)
		}
	}()

	// Start the internal API, which should never be exposed publicly.
	go func() {
		router := web.NewRouter()
		router.HandleFunc("/_/ping", handlers.PingHandler)
		router.HandleFunc("/*id", handlers.GoRegisterHandler)

		s := &http.Server{
			Addr:    ":9090",
			Handler: router,
		}
		logging.Infof("starting api server on port 9090")
		if err := s.ListenAndServe(); err != nil {
			logging.Fatalf("failed to start api server: %s", err)
		}
	}()

	// Exit cleanly.
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
}
