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
	"fmt"
	"os"
)

const configFolder = "hlsadmin"

func SettingsLocation(isCOntainer bool) (string, error) {
	if isCOntainer {
		return fmt.Sprintf("/etc/%v", configFolder), nil
	}
	location, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v%v.%v", location, string(os.PathSeparator), configFolder), nil
}

func InitSettingLoc(location string, filename string) (string, error) {
	return Create(location, filename, []byte{})
}
