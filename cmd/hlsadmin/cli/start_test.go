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

package cli

import (
	"errors"
	"testing"
)

func TestStartCmdName(t *testing.T) {

	cmd := createStartCmd()

	expected := "start"
	got := cmd.Use
	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestAppInitConfigYamlNoHome(t *testing.T) {
	initHomeConfigDir = func() (string, error) {
		return "", errors.New("Failed")
	}
	initConfigStore = func(name string) (string, error) {
		t.Fatal("Expected no call to initialise home config")
		return "", nil
	}
	initConfigYaml = func(name string) (string, error) {
		t.Fatal("Expected no call to initialise config yaml")
		return "", nil
	}

	appInitConfigYaml()
}

func TestAppInitConfigYamlHomeFoundStoreFailed(t *testing.T) {
	initHomeConfigDir = func() (string, error) {
		return "./test", nil
	}
	initConfigStore = func(name string) (string, error) {
		return "", errors.New("Failed")
	}
	initConfigYaml = func(name string) (string, error) {
		t.Fatal("Expected no call to initialise config yaml")
		return "", nil
	}

	appInitConfigYaml()
}

func TestAppInitConfigYamlFound(t *testing.T) {
	expectedYaml := "./test/settings.yaml"
	initHomeConfigDir = func() (string, error) {
		return "./test", nil
	}
	initConfigStore = func(name string) (string, error) {
		return name, nil
	}
	initConfigYaml = func(name string) (string, error) {
		return expectedYaml, nil
	}

	gotYaml, err := appInitConfigYaml()
	if err != nil {
		t.Fatalf("Unable to create config yaml. Reason: %v", err)
	}
	if expectedYaml != gotYaml {
		t.Fatalf("Expected: %s Got: %s", expectedYaml, gotYaml)
	}
}
