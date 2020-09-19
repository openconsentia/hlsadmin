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

package json

import (
	"openconsentia/hlsadmin/internal/authuser"
	"reflect"
	"testing"
)

func TestMarshalAccessCred(t *testing.T) {
	dataSet := []struct {
		input    authuser.AccessCredential
		expected []byte
		err      error
	}{
		{
			input:    authuser.AccessCredential{},
			expected: []byte(`{"id":"","displayName":"","accessToken":"","accessRole":""}`),
			err:      nil,
		},
		{
			input: authuser.AccessCredential{
				ID: "id",
			},
			expected: []byte(`{"id":"id","displayName":"","accessToken":"","accessRole":""}`),
			err:      nil,
		},
	}

	for index, dataItem := range dataSet {
		serializer := &seralizerSvc{}
		got, err := serializer.MarshalAccessCred(&dataItem.input)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.err, err)
			}
			if !reflect.DeepEqual(dataItem.expected, got) {
				t.Fatalf("Index: %v Expected: %v Got: %v", index, string(dataItem.expected), string(got))
			}
		} else {
			if dataItem.err.Error() != err.Error() {
				t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.err, err)
			}
		}
	}
}

func TestUnmarshalLoginCred(t *testing.T) {
	dataSet := []struct {
		input    []byte
		expected authuser.LoginCredential
		err      error
	}{
		{
			input: []byte(`{"id":"id","secrets":"secrets"}`),
			expected: authuser.LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			err: nil,
		},
	}
	for index, dataItem := range dataSet {
		serializer := &seralizerSvc{}
		got, err := serializer.UnmarshalLoginCred(dataItem.input)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.err, err)
			}
			if !reflect.DeepEqual(&dataItem.expected, got) {
				t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.expected, got)
			}
		}
	}
}
