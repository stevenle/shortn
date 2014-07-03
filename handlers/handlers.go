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

package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stevenle/shortn/kvstore"
	"github.com/stevenle/web"
)

func PingHandler(ctx *web.Context) {
	web.WriteResponseString(ctx, "pong\n")
}

func GoHandler(ctx *web.Context) {
	method := ctx.Request.Method
	if !(method == "GET" || method == "HEAD") {
		web.SetStatusCode(ctx, http.StatusMethodNotAllowed)
		return
	}

	u := kvstore.Get(ctx.Params["id"])
	if u == "" {
		web.SetStatusCode(ctx, http.StatusNotFound)
		return
	}
	web.Redirect(ctx, u, http.StatusFound)
}

type PutRequest struct {
	Url string `json:"url"`
}

func GoRegisterHandler(ctx *web.Context) {
	if ctx.Request.Method != "PUT" {
		web.SetStatusCode(ctx, http.StatusMethodNotAllowed)
		return
	}

	id := ctx.Params["id"]
	if id == "" {
		web.SetStatusCode(ctx, http.StatusBadRequest)
		return
	}

	var request PutRequest
	decoder := json.NewDecoder(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	err := decoder.Decode(&request)
	if err != nil {
		web.SetStatusCode(ctx, http.StatusBadRequest)
		return
	}

	u, err := url.Parse(request.Url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		web.SetStatusCode(ctx, http.StatusBadRequest)
		return
	}

	created := kvstore.Set(id, request.Url)
	if created {
		web.SetStatusCode(ctx, http.StatusCreated)
	}
	web.WriteResponseString(ctx, request.Url)
}
