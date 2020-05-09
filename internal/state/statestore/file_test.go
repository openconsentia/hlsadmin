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

package statestore

import "testing"

func TestCreateFileWhenItDoesNotExist(t *testing.T) {

	var configFolderCreateCalled = 0
	configFileExists = func(name string) bool {
		return false
	}

	configFileCreate = func(name string, content []byte) error {
		configFolderCreateCalled = 1
		return nil
	}

	NewFile("./test", "config.yaml", []byte("Hello"))

	if configFolderCreateCalled != 1 {
		t.Fatalf("Expected: 1 Got: %d", configFolderCreateCalled)
	}
}

func TestCreateFileWhenItExits(t *testing.T) {

	configFileExists = func(name string) bool {
		return true
	}
	configFileCreate = func(name string, content []byte) error {
		t.Errorf("Operations to create file should not be called")
		return nil
	}

	NewFile("./test", "config.yaml", []byte("Hello"))

}
