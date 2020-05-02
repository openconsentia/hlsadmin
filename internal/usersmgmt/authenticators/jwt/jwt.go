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

package jwt

import (
	"fmt"
	"hls-devkit/hlsadmin/internal/usersmgmt"
	"strings"
)

const bearerPrefix = "Bearer"

func AddBearerPrefix(input string) string {
	return fmt.Sprintf("Bearer %s", input)
}

func RemoveBearerPrefix(input string) string {
	if strings.Contains(input, bearerPrefix) {
		result := strings.TrimLeft(input, bearerPrefix)
		result = strings.TrimSpace(result)
		return result
	}
	return input
}

// Response is an representation of the response data type
// from jwt authenticator.
type Response struct {
	Token    string
	UserInfo usersmgmt.UserInfo
}

// JWTAuthenticator is an implementation of an jwt based
// authenticator. THIS IS ONLY A PLACEHOLDER IMPLEMENTATION
func Authenticator(id string, secrets string) (usersmgmt.AuthenticatorResponse, error) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	info := usersmgmt.UserInfo{
		ID:          id,
		DisplayName: fmt.Sprintf("%v %v", id, "display name"),
	}
	response := Response{
		Token:    token,
		UserInfo: info,
	}
	return response, nil
}

// MockJWTAuthenticator this is a mock implementation intended
// to support test or mock RESTFul API for UI development.
func MockAuthenticator(id string, secrets string) (usersmgmt.AuthenticatorResponse, error) {

	token := "<Token value>"
	info := usersmgmt.UserInfo{
		ID:          id,
		DisplayName: fmt.Sprintf("%v %v", id, "display name"),
	}
	response := Response{
		Token:    token,
		UserInfo: info,
	}
	return response, nil
}
