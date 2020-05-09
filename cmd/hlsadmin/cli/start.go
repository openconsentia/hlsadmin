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
	"hls-devkit/hlsadmin/internal/state/setting"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	initHomeConfigDir func() (string, error)       = setting.SettingsHomeFolder
	initConfigStore   func(string) (string, error) = setting.InitialiseSettingsStore
	initConfigYaml    func(string) (string, error) = setting.InitialiseSettingsYaml
)

func containerSettingsYam() (string, error) {
	configDir := "/var/hlsadmin"
	store, err := initConfigStore(configDir)
	if err != nil {
		return "", err
	}
	configYaml, err := initConfigYaml(store)
	if err != nil {
		return "", err
	}
	return configYaml, nil
}

func nativeSettingsYaml() (string, error) {
	configDir, err := initHomeConfigDir()
	if err != nil {
		return "", err
	}
	store, err := initConfigStore(configDir)
	if err != nil {
		return "", err
	}
	configYaml, err := initConfigYaml(store)
	if err != nil {
		return "", err
	}
	return configYaml, nil
}

func createStartCmd() *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "choice of features",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			_, envIsSet := os.LookupEnv("CONTAINER")
			if envIsSet {
				yamlfile, err := containerSettingsYam()
				if err != nil {
					log.Fatalf("Unable to create yaml settings. Reason: %v", err)
				}
				log.Printf("Settings file created: %v", yamlfile)
			} else {
				yamlfile, err := nativeSettingsYaml()
				if err != nil {
					log.Fatalf("Unable to create yaml settings. Reason: %v", err)
				}
				log.Printf("Settings file created: %v", yamlfile)
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
