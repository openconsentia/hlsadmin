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
	"goweb/internal"
	"goweb/internal/usermgmt"
	"net/http"

	"github.com/gorilla/mux"
)

func authUserHandler(rw http.ResponseWriter, req *http.Request) {

	userInfo := usermgmt.UserInfo{
		ID:          "john",
		DisplayName: "John Doe",
	}
	userInfoInBytes, _ := userInfo.Marshal()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	authUser := usermgmt.AuthUser{}
	authUser.Authenticator = func(id string, secret string) (string, []byte, error) {
		return token, userInfoInBytes, nil
	}
	authUser.Handler(rw, req)
}

func RESTRun(router *mux.Router) {
	router.HandleFunc(internal.URLAuthPath, authUserHandler)
	router.Use(mux.CORSMethodMiddleware(router))
}
