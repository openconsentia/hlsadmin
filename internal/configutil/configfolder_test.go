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
	"path"
	"testing"
)

func TestNewConfigFolderWhenExist(t *testing.T) {
	configFolderExists = func(name string) bool {
		return true
	}

	configFolderCreate = func(name string) error {
		t.Fatalf("Do not expect create folder to be called.")
		return nil
	}

	inputName := "./test"
	folder, _ := newConfigurationFolder(inputName)
	if folder != path.Join(inputName) {
		t.Fatalf("Expected: %s Got: %s", path.Join(inputName), folder)
	}

}
