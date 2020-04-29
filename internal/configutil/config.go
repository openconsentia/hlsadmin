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

package configutil

import (
	"os"
	"path"
)

const DefaultFolderName = ".hlsadmin"

func HomePathToConfigFolder() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	folder := path.Join(home, DefaultFolderName)
	return folder, nil
}

func NewConfigurationFolder(name string) (string, error) {
	return newConfigurationFolder(name)
}

func NewSettingsFile(foldername string) (string, error) {
	return newConfigurationFile(foldername, "settings.yaml", []byte{})
}
