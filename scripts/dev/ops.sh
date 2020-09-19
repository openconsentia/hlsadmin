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

export REST_SERVER_NAME=gorest

COMMAND="$1"

function build() {
    docker-compose -f ./deployments/dev/docker-compose.yaml build
}

function run() {
    docker-compose -f ./deployments/dev/docker-compose.yaml up -d
}

function stop(){
    docker-compose -f ./deployments/dev/docker-compose.yaml down
}

function clean(){
    docker-compose -f ./deployments/dev/docker-compose.yaml down
    docker rm -f $(docker ps -aq)
    docker rmi -f ${REACT_IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

message="Usage: $0 build | run | stop | clean"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "run")
        run
        ;;
    "stop")
        stop
        ;;
    *)
        echo $message
        ;;
esac