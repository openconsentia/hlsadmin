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
	"fmt"
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

func TestStartCmdNoHomeDir(t *testing.T) {
	initHomeDir = func() (string, error) {
		return "", fmt.Errorf("No home")
	}
	initConfigStore = func(string) (string, error) {
		t.Fatal("There should be no attempt to initialise store")
		return "", nil
	}
	appInit()
}

func TestStartCmdAppInit(t *testing.T) {
	initOpsCall := 0
	initHomeDir = func() (string, error) {
		return "test", nil
	}

	initConfigStore = func(string) (string, error) {
		initOpsCall += 1
		return "", nil
	}
	appInit()
	if initOpsCall != 1 {
		t.Fatalf("Expected: 1 Got: %d", initOpsCall)
	}
}
