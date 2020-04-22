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
	"os"

	"path"

	"github.com/spf13/cobra"
)

type StartCmdBuilder struct {
	initapp func() (string, error)
}

func (s *StartCmdBuilder) cli() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "choice of features",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			configDir, err := s.initapp()
			if err != nil {
				log.Fatalf("Unable to start. Reason: %v", err)
			}
			log.Printf("Created configuration store at %v", configDir)
		},
	}
}

var startCmdBuilder = StartCmdBuilder{}

func init() {

	startCmdBuilder.initapp = func() (string, error) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configFolder := path.Join(home, ".hlsadmin")
		dir, err := configutil.NewConfigurationFolder(configFolder)
		if err != nil {
			return "", err
		}
		return dir, nil
	}

}

func initStartCmd() *cobra.Command {

	startCmd := startCmdBuilder.cli()

	uiCmd := uiCmdBuilder.cli()
	startCmd.AddCommand(uiCmd)
	uiCmd.Flags().IntVarP(&uiCmdBuilder.port, "port", "p", 80, "startup default port 80")

	noUICmd := noUICmdBuilder.cli()
	startCmd.AddCommand(noUICmd)
	noUICmd.Flags().IntVarP(&noUICmdBuilder.port, "port", "p", 8080, "startup default port 8080")

	return startCmd
}
