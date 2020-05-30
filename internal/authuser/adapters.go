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

// Seralizer user credentials
type Seralizer interface {
	MarshalAccessCred(*AccessCredential) ([]byte, error)
	UnmarshalLoginCred([]byte) (*LoginCredential, error)
}

// CredentialRepo is an interface to the cred repository
type CredentialRepo interface {
	FindLoginCred(string) (*LoginCredential, error)
	FindAccessCred(string) (*AccessCredential, error)
}

// Authenticator is an interface to a credential service
type Authenticator interface {
	WithLoginCred(*LoginCredential) (*AccessCredential, error)
}
