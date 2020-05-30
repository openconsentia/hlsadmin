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
	"fmt"
	"testing"
)

func TestIsValidLoginCred(t *testing.T) {
	dataSet := []struct {
		input *LoginCredential
		err   error
	}{
		{
			input: nil,
			err:   fmt.Errorf("Input loginCredential is nil"),
		},
		{
			input: &LoginCredential{},
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			input: &LoginCredential{
				ID: "",
			},
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			input: &LoginCredential{
				ID: "id",
			},
			err: &fieldError{
				FieldName: "Secrets",
			},
		},
		{
			input: &LoginCredential{
				ID:      "id",
				Secrets: "",
			},
			err: &fieldError{
				FieldName: "Secrets",
			},
		},
		{
			input: &LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			err: nil,
		},
	}

	for _, dataItem := range dataSet {
		err := IsValidLoginCred(dataItem.input)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
		} else {
			if dataItem.err.Error() != err.Error() {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
		}
	}
}

func TestIsValidAccessCred(t *testing.T) {
	dataSet := []struct {
		input *AccessCredential
		err   error
	}{
		{
			input: nil,
			err:   fmt.Errorf("Input loginCredential is nil"),
		},
		{
			input: &AccessCredential{},
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			input: &AccessCredential{
				ID: "",
			},
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			input: &AccessCredential{
				ID: "id",
			},
			err: nil,
		},
	}
	for _, dataItem := range dataSet {
		err := IsValidAccessCred(dataItem.input)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
		} else {
			if dataItem.err.Error() != err.Error() {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
		}
	}
}

func TestCompareLoginCred(t *testing.T) {
	dataSet := []struct {
		tocompare   *LoginCredential
		comparewith *LoginCredential
		expected    bool
		err         error
	}{
		{
			tocompare: &LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			comparewith: &LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			expected: true,
			err:      nil,
		},
		{
			tocompare: &LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			comparewith: &LoginCredential{
				ID:      "id",
				Secrets: "secret",
			},
			expected: false,
			err:      fmt.Errorf("Login credentials not the same"),
		},
		{
			tocompare: &LoginCredential{
				ID:      "id",
				Secrets: "secrets",
			},
			comparewith: &LoginCredential{
				ID:      "i",
				Secrets: "secrets",
			},
			expected: false,
			err:      fmt.Errorf("Login credentials not the same"),
		},
		{
			tocompare: &LoginCredential{
				ID:      "",
				Secrets: "secrets",
			},
			comparewith: &LoginCredential{
				ID:      "id",
				Secrets: "secret",
			},
			expected: false,
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			tocompare: &LoginCredential{
				ID:      "id",
				Secrets: "",
			},
			comparewith: &LoginCredential{
				ID:      "id",
				Secrets: "secret",
			},
			expected: false,
			err: &fieldError{
				FieldName: "Secrets",
			},
		},
	}

	for _, dataItem := range dataSet {
		same, err := CompareLoginCred(dataItem.tocompare, dataItem.comparewith)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
			if dataItem.expected != same {
				t.Fatalf("Expected: %v Got: %v", dataItem.expected, same)
			}
		} else {
			if dataItem.err.Error() != err.Error() {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
			if dataItem.expected != same {
				t.Fatalf("Expected: %v Got: %v", dataItem.expected, same)
			}
		}
	}
}

func TestCompareAccessCred(t *testing.T) {
	dataSet := []struct {
		tocompare   *AccessCredential
		comparewith *AccessCredential
		expected    bool
		err         error
	}{
		{
			tocompare: &AccessCredential{
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			comparewith: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			expected: false,
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			tocompare: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			comparewith: &AccessCredential{
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			expected: false,
			err: &fieldError{
				FieldName: "ID",
			},
		},
		{
			tocompare: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			comparewith: &AccessCredential{
				ID:          "id",
				DisplayName: "a",
				AccessToken: "",
				AccessRole:  "",
			},
			expected: false,
			err:      fmt.Errorf("Access credentials not the same"),
		},
		{
			tocompare: &AccessCredential{
				ID: "id",
			},
			comparewith: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			expected: true,
			err:      nil,
		},
		{
			tocompare: &AccessCredential{
				ID:          "id",
				DisplayName: "a",
				AccessToken: "b",
				AccessRole:  "c",
			},
			comparewith: &AccessCredential{
				ID:          "id",
				DisplayName: "a",
				AccessToken: "b",
				AccessRole:  "c",
			},
			expected: true,
			err:      nil,
		},
		{
			tocompare: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			comparewith: &AccessCredential{
				ID:          "id",
				DisplayName: "",
				AccessToken: "",
				AccessRole:  "",
			},
			expected: true,
			err:      nil,
		},
	}

	for _, dataItem := range dataSet {
		same, err := CompareAccessCred(dataItem.tocompare, dataItem.comparewith)
		if dataItem.err == nil {
			if dataItem.err != err {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
			if dataItem.expected != same {
				t.Fatalf("Expected: %v Got: %v", dataItem.expected, same)
			}
		} else {
			if dataItem.err.Error() != err.Error() {
				t.Fatalf("Expected: %v Got: %v", dataItem.err, err)
			}
			if dataItem.expected != same {
				t.Fatalf("Expected: %v Got: %v", dataItem.expected, same)
			}
		}
	}
}
