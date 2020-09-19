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

package server

import (
	"io/ioutil"
	"net/http"
	"openconsentia/hlsadmin/internal/authuser"
	"openconsentia/hlsadmin/internal/repo/file"
	jsonserializer "openconsentia/hlsadmin/internal/serializer/json"
	"os"
)

var (
	authenticate authuser.Authenticator
	serializer   authuser.Seralizer
	credRepo     authuser.CredentialRepo
)

func init() {
	serializer = jsonserializer.NewSerializer()
	_, exists := os.LookupEnv("DEV")
	if exists {
		credRepo = file.NewMockCredentialRepoService()
		authenticate = authuser.NewMockAuthenticateService(credRepo)
	} else {
		credRepo = file.NewCredentialRepoService()
		authenticate = authuser.NewAuthenticationService(credRepo)
	}
}

func authUserHdler(rw http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.URL.Path != URLAuthPath {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer req.Body.Close()

	loginCred, err := serializer.UnmarshalLoginCred(body)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessCred, err := authenticate.WithLoginCred(loginCred)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO create security services
	// Add jwt token to accessCred token field

	rw.Header().Set(HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(HTTPHeaderContentType, "application/json")

	accessCredBytes, err := serializer.MarshalAccessCred(accessCred)
	if err != nil {
		rw.WriteHeader(http.StatusNoContent)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(accessCredBytes)
}
