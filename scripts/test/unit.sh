#!/bin/bash

# Copyright 2020 Paul Sitoh
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

COMMAND="$1"

go_unit_test_image=hls_devkit/go_unit_test_image:current
react_unit_test_image=hls_devkit/react_unit_test_image:current

case $COMMAND in
    "go")
        docker build -f ./build/package/test/unit/go.dockerfile -t ${go_unit_test_image} .
        ;;
    "react")
        docker build -f ./build/package/test/unit/react.dockerfile -t ${react_unit_test_image} .
        ;;
    "clean")
        docker rmi -f ${react_unit_test_image}
        docker rmi -f ${go_unit_test_image}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "$0 [ go | react ]"
        ;;
esac