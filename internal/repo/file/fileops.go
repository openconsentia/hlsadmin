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
	"io/ioutil"
	"os"
	"path"
)

var (
	folderExist func(string) bool = func(name string) bool {
		_, err := os.Stat(name)
		if os.IsExist(err) {
			return true
		}
		return false
	}

	folderCreate func(string) error = func(name string) error {
		err := os.MkdirAll(name, 0700)
		if err != nil {
			return err
		}
		return nil
	}

	fileExist func(string) bool = func(name string) bool {
		_, err := os.Stat(name)
		if os.IsExist(err) {
			return true
		}
		return false
	}
	fileCreate func(string, []byte) error = func(name string, content []byte) error {
		err := ioutil.WriteFile(name, content, 0644)
		if err != nil {
			return err
		}
		return nil
	}
)

func Create(folder string, name string, content []byte) (string, error) {
	fldr := path.Join(folder)
	if !folderExist(fldr) {
		err := folderCreate(fldr)
		if err != nil {
			return "", err
		}
	}

	fullFilename := path.Join(fldr, name)
	if !fileExist(fullFilename) {
		err := fileCreate(fullFilename, content)
		if err != nil {
			return "", err
		}
	}

	return fullFilename, nil
}
