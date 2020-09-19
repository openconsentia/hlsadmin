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
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"openconsentia/hlsadmin/internal/authuser"
	"openconsentia/hlsadmin/internal/repo/file"
	"testing"
)

func TestInvalidAuthHandlerMethodAndPath(t *testing.T) {

	dataSet := []struct {
		method     string
		path       string
		httpstatus int
	}{
		{
			method:     "GET",
			path:       URLAuthPath,
			httpstatus: http.StatusUnauthorized,
		},
		{
			method:     "PUT",
			path:       URLAuthPath,
			httpstatus: http.StatusUnauthorized,
		},
		{
			method:     "POST",
			path:       "/users",
			httpstatus: http.StatusUnauthorized,
		},
	}

	for _, dataItem := range dataSet {
		req, err := http.NewRequest(dataItem.method, dataItem.path, nil)
		if err != nil {
			t.Fatalf("Unable to create a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		http.HandlerFunc(authUserHdler).ServeHTTP(rr, req)

		if rr.Code != dataItem.httpstatus {
			t.Fatalf("Expected status: %v Got status: %v", dataItem.httpstatus, rr.Code)
		}
	}
}

func mockRequestBody(t *testing.T) []byte {
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

func TestEmptyRequestBody(t *testing.T) {

	credRepo = file.NewMockCredentialRepoService()
	authenticate = authuser.NewMockAuthenticateService(credRepo)

	reqBody := mockRequestBody(t)
	dataSet := []struct {
		body       *bytes.Reader
		httpstatus int
	}{
		{
			body:       bytes.NewReader([]byte("")),
			httpstatus: http.StatusUnauthorized,
		},
		{
			body:       bytes.NewReader([]byte(reqBody)),
			httpstatus: http.StatusOK,
		},
	}

	for index, dataItem := range dataSet {

		req, err := http.NewRequest("POST", URLAuthPath, dataItem.body)
		if err != nil {
			t.Fatalf("Index: %v Unable to create request", index)
		}

		rr := httptest.NewRecorder()
		http.HandlerFunc(authUserHdler).ServeHTTP(rr, req)

		if dataItem.httpstatus != rr.Code {
			t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.httpstatus, rr.Code)
		}
	}
}
