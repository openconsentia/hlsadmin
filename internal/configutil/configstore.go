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

import "os"

type Config struct {
	storeExist  func(string) bool
	createStore func(string) error
}

func (s *Config) StoreExist(location string) bool {
	return s.storeExist(location)
}
func (s *Config) CreateStore(location string) error {
	return s.createStore(location)
}

var configOps Config

func init() {
	configOps = Config{
		storeExist: func(location string) bool {
			_, err := os.Stat(location)
			if os.IsExist(err) {
				return true
			}
			return false
		},

		createStore: func(location string) error {
			err := os.Mkdir(location, 0700)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func NewConfigStore(location string) (string, error) {

	if !configOps.storeExist(location) {
		err := configOps.createStore(location)
		if err != nil {
			return "", nil
		}
		return location, nil
	}

	return location, nil
}
