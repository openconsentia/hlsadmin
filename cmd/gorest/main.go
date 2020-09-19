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

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"openconsentia/hlsadmin/internal/server"

	"github.com/gorilla/mux"
)

func main() {

	portPtr := flag.Int("port", 80, "Default port: 80")
	flag.Parse()

	router := mux.NewRouter()

	server.RESTRun(router)
	log.Printf("Starting with UI on port %v", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", *portPtr), router))

}
