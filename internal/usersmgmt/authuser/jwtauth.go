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
	"encoding/json"
	"fmt"
	"strings"
)

const bearerPrefix = "Bearer"

type bearerToken string

func addBearerPrefix(input bearerToken) string {
	return fmt.Sprintf("Bearer %s", input)
}

func removeBearerPrefix(input string) bearerToken {
	if strings.Contains(input, bearerPrefix) {
		result := strings.TrimLeft(input, bearerPrefix)
		result = strings.TrimSpace(result)
		return bearerToken(result)
	}
	return bearerToken(input)
}

var jwtAuthenticator func(id string, secrets string) (bearerToken, userInfoBytes, error)

// This is a mock implementation. Replace this with an actual one
func withJWT(id string, secrets string) (bearerToken, userInfoBytes, error) {

	token := bearerToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	info := userInfo{
		ID:          id,
		DisplayName: fmt.Sprintf("%v %v", id, "display name"),
	}
	infoBytes, err := json.Marshal(info)
	if err != nil {
		return bearerToken(""), userInfoBytes(""), err
	}

	return token, infoBytes, nil
}

func init() {
	jwtAuthenticator = withJWT
}
