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
	"io/ioutil"
	"os"
	"path"
)

type configurationFileOperations struct {
	folderExist  func(name string) bool
	folderCreate func(name string) (string, error)
	fileExist    func(name string) bool
	fileCreate   func(name string, content []byte) error
}

var configurationFileOps configurationFileOperations

func configFolderExist(name string) bool {
	_, err := os.Stat(name)
	if os.IsExist(err) {
		return true
	}
	return false
}

func configFolderCreate(name string) (string, error) {
	dir := path.Join(name)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func configFileExist(name string) bool {
	_, err := os.Stat(name)
	if os.IsExist(err) {
		return true
	}
	return false
}

func configFileCreate(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	configurationFileOps.folderExist = configFolderExist
	configurationFileOps.folderCreate = configFolderCreate
	configurationFileOps.fileExist = configFileExist
	configurationFileOps.fileCreate = configFileCreate
}

func NewConfigurationFile(foldername string, filename string, content []byte) (string, error) {

	var dir string
	var err error
	if !configurationFileOps.folderExist(foldername) {
		dir, err = configurationFileOps.folderCreate(foldername)
		if err != nil {
			return "", err
		}
	}

	filenameFull := path.Join(dir, filename)
	if !configurationFileOps.fileExist(filenameFull) {
		configurationFileOps.fileCreate(filenameFull, content)
	}

	return filenameFull, nil
}
