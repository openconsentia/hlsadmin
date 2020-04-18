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

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func CreateVolume(name string) (types.Volume, error) {

	ctx := context.Background()

	cli, err := client.NewEnvClient()
	if err != nil {
		return types.Volume{}, err
	}

	vol, err := cli.VolumeCreate(ctx, volume.VolumesCreateBody{Name: name})
	if err != nil {
		return types.Volume{}, err
	}

	return vol, nil
}
