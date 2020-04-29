// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authuser

import (
	"encoding/json"
	"hls-devkit/hlsadmin/internal"
	"io/ioutil"
	"net/http"
)

type credentials struct {
	ID      string `json:"id"`
	Secrets string `json:"secrets"`
}

func parseAuthBody(req *http.Request) (*credentials, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	var cred credentials
	err = json.Unmarshal(body, &cred)
	if err != nil {
		return nil, err
	}
	return &cred, nil

}

func Handler(rw http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.URL.Path != internal.URLAuthPath {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	cred, err := parseAuthBody(req)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, userInfo, err := jwtAuthenticator(cred.ID, cred.Secrets)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	rw.Header().Set(internal.HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(internal.HTTPHeaderContentType, "application/json")
	rw.Header().Set(internal.HTTPHeaderAuthorization, addBearerPrefix(token))

	rw.WriteHeader(http.StatusOK)
	rw.Write(userInfo)
}
