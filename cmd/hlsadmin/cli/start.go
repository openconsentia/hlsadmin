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
	"hls-devkit/hlsadmin/internal/configutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	initConfigHomeDir func() (string, error)       = configutil.HomeConfigFolder
	initConfigStore   func(string) (string, error) = configutil.InitialiseConfigStore
	initSettingsFile  func(string) (string, error) = configutil.InitialiseSettingsFile
)

func appInit() error {
	home, err := initConfigHomeDir()
	if err != nil {
		return err
	}

	configDir, err := initConfigStore(home)
	if err != nil {
		return err
	}
	log.Printf("Created configuration store at %v", configDir)
	return nil
}

func createStartCmd() *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "choice of features",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := appInit()
			if err != nil {
				log.Fatalf("Unable to initialise app. Reason: %v", err)
			}
		},
	}

	uiCmd := uiCmdBlder.cli()
	startCmd.AddCommand(uiCmd)
	uiCmd.Flags().IntVarP(&uiCmdBlder.port, "port", "p", 80, "startup default port 80")

	noUICmd := noUICmdBlder.cli()
	startCmd.AddCommand(noUICmd)
	noUICmd.Flags().IntVarP(&noUICmdBlder.port, "port", "p", 8080, "startup default port 8080")

	return startCmd
}
