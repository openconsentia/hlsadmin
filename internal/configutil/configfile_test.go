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

import "testing"

func TestWhenConfigFolderDoesNotExist(t *testing.T) {
	var configFolderCreateCalled = 0
	configurationFileOps = configurationFileOperations{
		folderExist: func(name string) bool {
			return false
		},
		folderCreate: func(name string) error {
			configFolderCreateCalled = 1
			return nil
		},
	}

	NewConfigurationFolder("./test/abc")

	if configFolderCreateCalled != 1 {
		t.Errorf("Exoected: 1 Got: %v", configFolderCreateCalled)
	}
}

func TestWhenConfigFolderExist(t *testing.T) {
	configurationFileOps = configurationFileOperations{
		folderExist: func(name string) bool {
			return true
		},
		folderCreate: func(name string) error {
			t.Error("The operations to create a folder should not be called")
			return nil
		},
	}

	NewConfigurationFolder("./test/abc")

}

func TestCreateFileWhenItDoesNotExist(t *testing.T) {

	var configFolderCreateCalled = 0

	configurationFileOps = configurationFileOperations{
		fileExist: func(name string) bool {
			return false
		},
		fileCreate: func(name string, content []byte) error {
			configFolderCreateCalled = 1
			return nil
		},
	}

	NewConfigurationFile("./test", "config.yaml", []byte("Hello"))

	if configFolderCreateCalled != 1 {
		t.Fatalf("Expected: 1 Got: %d", configFolderCreateCalled)
	}
}

func TestCreateFileWhenItExits(t *testing.T) {

	configurationFileOps = configurationFileOperations{
		fileExist: func(name string) bool {
			return true
		},
		fileCreate: func(name string, content []byte) error {
			t.Errorf("Operations to create file should not be called")
			return nil
		},
	}

	NewConfigurationFile("./test", "config.yaml", []byte("Hello"))

}
