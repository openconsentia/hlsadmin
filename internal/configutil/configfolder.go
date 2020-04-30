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

var (
	configFolderExists func(name string) bool = func(name string) bool {
		_, err := os.Stat(name)
		if os.IsExist(err) {
			return true
		}
		return false
	}

	configFolderCreate func(name string) error = func(name string) error {
		err := os.MkdirAll(name, 0700)
		if err != nil {
			return err
		}
		return nil
	}
)

func newConfigurationFolder(name string) (string, error) {

	dir := path.Join(name)
	if !configFolderExists(name) {
		err := configFolderCreate(name)
		if err != nil {
			return "", err
		}
	}

	return dir, nil
}
