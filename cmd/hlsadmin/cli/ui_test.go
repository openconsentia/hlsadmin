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

import "testing"

func TestUICmdName(t *testing.T) {
	builder := UICmdBuilder{}
	builder.service = func() {}

	cmd := builder.cli()
	expected := "ui"
	got := cmd.Use
	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestUICmdPortFlag(t *testing.T) {
	builder := UICmdBuilder{
		port: 80,
	}
	builder.service = func() {
		expected := 1000
		got := builder.port
		if expected != got {
			t.Errorf("Expected: %v Got: %v", expected, got)
		}
	}
	cmd := builder.cli()
	// The following simulates port sets to 1000
	cmd.Flags().IntVarP(&builder.port, "port", "p", 1000, "startup default port 80")
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unable to execute command. Reason: %v", err)
	}
}
