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
	"reflect"
)

// Error handlers

const errorPrefix = "Missing:"

type fieldError struct {
	FieldName string
}

func (e *fieldError) Error() string {
	return fmt.Sprintf("%v %v", errorPrefix, e.FieldName)
}

func isEmptyString(value string) bool {
	if len(value) == 0 {
		return true
	}
	return false
}

// LoginCredential associated with a person presented for authentication
type LoginCredential struct {
	// ID assigned to person holding this credential
	ID string `json:"id"`
	// Secrets provided by the person with this credential
	Secrets string `json:"secrets"`
}

func IsValidLoginCred(cred *LoginCredential) error {

	if cred == nil {
		return fmt.Errorf("Input loginCredential is nil")
	}

	if isEmptyString(cred.ID) {
		return &fieldError{
			FieldName: "ID",
		}
	}
	if isEmptyString(cred.Secrets) {
		return &fieldError{
			FieldName: "Secrets",
		}
	}
	return nil
}

func CompareLoginCred(cred *LoginCredential, compared *LoginCredential) (bool, error) {
	err := IsValidLoginCred(cred)
	if err != nil {
		return false, err
	}

	err = IsValidLoginCred(compared)
	if err != nil {
		return false, err
	}
	if !reflect.DeepEqual(cred, compared) {
		return false, fmt.Errorf("Login credentials not the same")
	}
	return true, nil
}

// AccessCredential contains information associated with a person
type AccessCredential struct {
	// ID is a unique string value associated with a user.
	ID string `json:"id"`
	// DisplayName is an optional representation of a user name
	// that can be read by a person
	DisplayName string `json:"displayName"`
	// JWTToken is an optional representation of a JWT Token.
	AccessToken string `json:"accessToken"`
	// AccessRole represents role
	AccessRole string `json:"accessRole"`
}

func IsValidAccessCred(cred *AccessCredential) error {

	if cred == nil {
		return fmt.Errorf("Input loginCredential is nil")
	}

	if isEmptyString(cred.ID) {
		return &fieldError{
			FieldName: "ID",
		}
	}
	return nil
}

func CompareAccessCred(cred *AccessCredential, compared *AccessCredential) (bool, error) {
	err := IsValidAccessCred(cred)
	if err != nil {
		return false, err
	}

	err = IsValidAccessCred(compared)
	if err != nil {
		return false, err
	}
	if !reflect.DeepEqual(cred, compared) {
		return false, fmt.Errorf("Access credentials not the same")
	}
	return true, nil
}
