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

package setting

import (
	"hls-devkit/hlsadmin/internal/state/statestore"
	"os"
	"path"
)

const (
	defaultHomeConfigBaseName = ".hlsadmin"
	defaultSettingsYAML       = "settings.yaml"
)

func SettingsHomeFolder() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	folder := path.Join(home, defaultHomeConfigBaseName)
	return folder, nil
}

func InitialiseSettingsStore(folder string) (string, error) {

	dir, err := statestore.NewFolder(folder)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func InitialiseSettingsYaml(foldername string) (string, error) {
	return statestore.NewFile(foldername, defaultSettingsYAML, []byte{})
}
