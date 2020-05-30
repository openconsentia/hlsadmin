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

import "fmt"

// actaul authenticator

type authenticateSvc struct {
	credRepo CredentialRepo
}

func (a *authenticateSvc) WithLoginCred(loginCred *LoginCredential) (*AccessCredential, error) {

	err := IsValidLoginCred(loginCred)
	if err != nil {
		return nil, err
	}

	storedLoginCred, err := a.credRepo.FindLoginCred(loginCred.ID)
	if err != nil {
		return nil, err
	}
	if storedLoginCred.ID != loginCred.ID {
		return nil, fmt.Errorf("Login credential ID not the same")
	}
	if storedLoginCred.Secrets != loginCred.Secrets {
		return nil, fmt.Errorf("Login credential secrets not the same")
	}

	accessCred, err := a.credRepo.FindAccessCred(loginCred.ID)
	if err != nil {
		return nil, fmt.Errorf("Access credentials not found")
	}

	return accessCred, nil
}

func NewAuthenticationService(credRepo CredentialRepo) Authenticator {
	return &authenticateSvc{
		credRepo: credRepo,
	}
}

// Mock authenticator

type mockAuthenticateSvc struct {
	credRepo CredentialRepo
}

func (m *mockAuthenticateSvc) WithLoginCred(loginCred *LoginCredential) (*AccessCredential, error) {

	return &AccessCredential{
		ID:          "id",
		DisplayName: "Display Name",
		AccessToken: "Access Token",
		AccessRole:  "Access Role",
	}, nil

}

// NewMockAuthenticateService instantiate a new autheticate service
func NewMockAuthenticateService(credRepo CredentialRepo) Authenticator {
	return &mockAuthenticateSvc{
		credRepo: credRepo,
	}
}
