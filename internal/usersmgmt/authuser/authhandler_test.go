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
	"bytes"
	"encoding/json"
	"hls-devkit/hlsadmin/internal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidAuthHandlerMethod(t *testing.T) {
	req, err := http.NewRequest("GET", internal.URLAuthPath, nil)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func TestInvalidAuthHandlerPath(t *testing.T) {
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusUnauthorized, rr.Code)
	}
}

func mockAuthUserRequestBody(t *testing.T) []byte {

	reqBody := struct {
		ID      string `json:"id"`
		Secrets string `json:"secrets"`
	}{
		ID:      "John",
		Secrets: "1we5d",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Unable to mock request body %v", err)
	}

	return bodyBytes
}

func TestJWTAutheticator(t *testing.T) {

	// Mocking
	reqBody := mockAuthUserRequestBody(t)
	expectedBearerToken := bearerToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

	jwtAuthenticator = func(id string, secrets string) (bearerToken, userInfoBytes, error) {
		tokenBytes := expectedBearerToken
		info := userInfo{
			ID:          id,
			DisplayName: id,
		}
		userInfoBytes, err := json.Marshal(info)
		if err != nil {
			t.Fatalf("Unable to marshal authenticator response. Reason: %v", err)
		}
		return tokenBytes, userInfoBytes, nil
	}

	// Initiate http test
	req, err := http.NewRequest("POST", internal.URLAuthPath, bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Unable to create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(Handler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status: %v Got status: %v", http.StatusOK, rr.Code)
	}

	resHeader := rr.Header()
	bearer := resHeader.Get(internal.HTTPHeaderAuthorization)
	actualToken := removeBearerPrefix(bearer)
	if expectedBearerToken != actualToken {
		t.Errorf("Expected: %v Got: %v", expectedBearerToken, actualToken)
	}

	resBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Unable to extract response body. Reason: %v", err)
	}
	var info userInfo
	err = json.Unmarshal(resBody, &info)
	if err != nil {
		t.Fatalf("Unable to unmarshal user info. Reason: %v", err)
	}
	if info.ID != "John" {
		t.Errorf("Expected: John Got: %v", info.ID)
	}

	if info.DisplayName != "John" {
		t.Errorf("Expected: John Got: %v", info.DisplayName)
	}
}
