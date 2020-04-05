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
	"goweb/internal/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type UICmdBuilder struct {
	port    int
	service func()
}

func (s *UICmdBuilder) cli() *cobra.Command {
	return &cobra.Command{
		Use:   "ui",
		Short: "application with ui",
		Run: func(cmd *cobra.Command, args []string) {
			s.service()
		},
	}
}

var uiCmdBuilder = UICmdBuilder{
	port: 80,
}

func init() {
	uiCmdBuilder.service = func() {
		router := mux.NewRouter()
		server.RESTRun(router)
		server.WebRun(router)
		log.Printf("Starting with UI on port %v", uiCmdBuilder.port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", uiCmdBuilder.port), router))
	}
}
