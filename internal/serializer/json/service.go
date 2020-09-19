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
	jsonstl "encoding/json"
	"openconsentia/hlsadmin/internal/authuser"
)

type seralizerSvc struct{}

func (s *seralizerSvc) MarshalAccessCred(accessCred *authuser.AccessCredential) ([]byte, error) {
	return jsonstl.Marshal(accessCred)
}
func (s *seralizerSvc) UnmarshalLoginCred(cred []byte) (*authuser.LoginCredential, error) {

	var loginCred authuser.LoginCredential

	err := jsonstl.Unmarshal(cred, &loginCred)
	if err != nil {
		return nil, err
	}

	return &loginCred, nil
}

func NewSerializer() authuser.Seralizer {
	return &seralizerSvc{}
}
