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

. ./scripts/common.sh

COMMAND="$1"

go_unit_image=${IMAGE_BASE_NAME}/go_unit_image:current
react_unit_image=${IMAGE_BASE_NAME}/react_unit_image:current

case $COMMAND in
    "go")
        docker run --rm \
            -w="/opt" \
            -v ${PWD}/internal:/opt/internal \
            -v ${PWD}/cmd:/opt/cmd \
            -v ${PWD}/go.mod:/opt/go.mod \
            -v ${PWD}/go.sum:/opt/go.sum \
            -v ${PWD}/test/unit/go-test.sh:/opt/go-test.sh \
            golang:${GO_VER} /opt/go-test.sh
        ;;
    "react")
        docker build -f ./test/unit/react.dockerfile \
            --build-arg NODE_VER=${NODE_VER} \
            --build-arg WEB_FRAMEWORK=${WEB_FRAMEWORK} \
            -t ${react_unit_image} .
        ;;
    "clean")
        docker rmi -f ${go_unit_image}
        docker rmi -f ${react_unit_image}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "$0 [ go | react | clean ]"
        ;;
esac