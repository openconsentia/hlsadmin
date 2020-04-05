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

package usermgmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goweb/internal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockAuthUserRequestBody(t *testing.T) []byte {

	reqBody := struct {
		ID      string `json:"id"`
		Secrets string `json:"secrets"`
	}{
		ID:      "Paul",
		Secrets: "1%we5d",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Unable to mock request body %v", err)
	}

	return bodyBytes
}

func mockResponse() (string, UserInfo) {

	info := UserInfo{
		ID:          "john",
		DisplayName: "John Doe",
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	return token, info
}

func mockAuthorization(token string, info UserInfo) (string, []byte, error) {
	infoInBytes, err := info.Marshal()
	if err != nil {
		return "", nil, fmt.Errorf("Unable to marshal user info, %v", err)
	}

	return token, infoInBytes, nil
}

func TestAuthUserInvalidMethod(t *testing.T) {

	expectedToken, expectedUserInfo := mockResponse()

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) (string, []byte, error) {
		return mockAuthorization(expectedToken, expectedUserInfo)
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("GET", internal.URLAuthPath, bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestAuthUserInvalidPath(t *testing.T) {

	expectedToken, expectedUserInfo := mockResponse()

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) (string, []byte, error) {
		return mockAuthorization(expectedToken, expectedUserInfo)
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", "/users", bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestAuthUserPayload(t *testing.T) {

	expectedToken, expectedUserInfo := mockResponse()

	authUser := AuthUser{}
	authUser.Authenticator = func(id string, secrets string) (string, []byte, error) {
		return mockAuthorization(expectedToken, expectedUserInfo)
	}

	bodyBytes := mockAuthUserRequestBody(t)
	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", internal.URLAuthPath, bytesReader)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(authUser.Handler).ServeHTTP(rr, req)

	// Verify header
	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}

	header := rr.Header()
	authorization := header.Get(internal.HTTPHeaderAuthorization)
	if strings.Compare(string(expectedToken), authorization) == 0 {
		t.Fatalf("Expected authorization: %v Got authorization: %v", string(expectedToken), authorization)
	}

	// Verify body
	resBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Unable to read response body: %v", err)
	}

	expectedUserInfoInByte, err := expectedUserInfo.Marshal()
	if err != nil {
		t.Fatalf("Unable to marshal user info: %v", err)
	}
	if bytes.Compare(resBody, expectedUserInfoInByte) != 0 {
		t.Fatalf("Expected: %v Got: %v", string(expectedUserInfoInByte), string(resBody))
	}
}
