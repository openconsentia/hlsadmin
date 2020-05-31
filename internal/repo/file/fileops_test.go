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
	"os"
	"path"
	"testing"
)

func TestCreateFileOps(t *testing.T) {

	t.Run("Folder does not exists", func(t *testing.T) {
		folderCreateCall := 0

		folderExist = func(name string) bool {
			return false
		}

		folderCreate = func(name string) error {
			folderCreateCall++
			return nil
		}

		fileExist = func(name string) bool {
			return true
		}

		fileCreate = func(name string, content []byte) error {
			return nil
		}

		Create("./test", "test", []byte{})
		if folderCreateCall != 1 {
			t.Fatalf("Expected: 1 Got: %d", folderCreateCall)
		}
	})

	t.Run("Folder exist", func(t *testing.T) {
		folderExist = func(name string) bool {
			return true
		}

		folderCreate = func(name string) error {
			t.Fatal("Create Folder should not be called")
			return nil
		}

		t.Run("File not exist", func(t *testing.T) {
			fileCreateCalled := 0
			fileExist = func(name string) bool {
				return false
			}

			fileCreate = func(name string, content []byte) error {
				fileCreateCalled++
				return nil
			}
			Create("./test", "test", []byte{})
			if fileCreateCalled != 1 {
				t.Fatalf("Expected: 1 Got: %d", fileCreateCalled)
			}
		})

		t.Run("File exist", func(t *testing.T) {
			fileExist = func(name string) bool {
				return true
			}

			fileCreate = func(name string, content []byte) error {
				t.Fatal("File create operations should not be called")
				return nil
			}
			Create("./test", "test", []byte{})
		})
	})

}

func TestCreateOps(t *testing.T) {
	testFolder := path.Join(".", "test")
	settings, err := Create(testFolder, "settings.yaml", []byte{})
	if err != nil {
		t.Fatalf("Expected: no error Got: %v", err)
	}
	defer os.RemoveAll(testFolder)
	if settings == "" {
		t.Fatal("Settings file expected")
	}
}
