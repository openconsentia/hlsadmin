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

package file

import (
	"openconsentia/hlsadmin/internal/authuser"
)

// actual credential repo

type credentialRepo struct{}

func (c *credentialRepo) FindLoginCred(id string) (*authuser.LoginCredential, error) {
	// TODO Read from file
	return &authuser.LoginCredential{
		ID:      "id",
		Secrets: "secrets",
	}, nil
}

func (c *credentialRepo) FindAccessCred(id string) (*authuser.AccessCredential, error) {
	// TODO Read from file
	return &authuser.AccessCredential{
		ID:          "id",
		DisplayName: "Test",
		AccessToken: "Test",
		AccessRole:  "Test",
	}, nil
}

// NewCredentialRepoService instantiate a credential repo
func NewCredentialRepoService() authuser.CredentialRepo {
	return &credentialRepo{}
}

// mock Credential repo

type mockCredentialRepo struct{}

func (m *mockCredentialRepo) FindLoginCred(id string) (*authuser.LoginCredential, error) {
	return &authuser.LoginCredential{
		ID:      "id",
		Secrets: "secrets",
	}, nil
}

func (m *mockCredentialRepo) FindAccessCred(id string) (*authuser.AccessCredential, error) {
	return &authuser.AccessCredential{
		ID:          "Test",
		DisplayName: "Test",
		AccessToken: "Test",
		AccessRole:  "Test",
	}, nil
}

// NewMockCredentialRepoService instantiate a mock credential repo
func NewMockCredentialRepoService() authuser.CredentialRepo {
	return &mockCredentialRepo{}
}
