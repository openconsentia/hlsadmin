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
	"path"
	"testing"
)

func TestSettingLocation(t *testing.T) {

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Unable to obtain home dir")
	}

	dataSet := []struct {
		iscontainer bool
		expected    string
	}{
		{
			iscontainer: true,
			expected:    fmt.Sprintf("/etc/%v", configFolder),
		},
		{
			iscontainer: false,
			expected:    fmt.Sprintf("%v%v.%v", home, string(os.PathSeparator), configFolder),
		},
	}

	for index, dataItem := range dataSet {
		got, err := SettingsLocation(dataItem.iscontainer)
		if err != nil {
			t.Fatalf("Index: %v Expected: no error Got: %v", index, err)
		}

		if dataItem.expected != got {
			t.Fatalf("Index: %v Expected: %v Got: %v", index, dataItem.expected, got)
		}
	}
}

func TestInitialiseSettingLoc(t *testing.T) {
	testSettingFile := "test.yaml"
	testFolder := path.Join(".", "test")
	got, err := InitSettingLoc(testFolder, testSettingFile)
	if err != nil {
		t.Fatalf("Expected: no error Got: %v", err)
	}
	defer os.RemoveAll(testFolder)
	expected := path.Join(testFolder, testSettingFile)
	if expected != got {
		t.Fatalf("Expected: %v Got: %v", expected, got)
	}
}
