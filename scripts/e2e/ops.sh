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

export DB_IMAGE_NAME=postgres
export DB_IMAGE_TAG=12

function start(){
    docker-compose -f ./deployments/e2e/docker-compose.yaml up -d
}

function stop(){
    docker-compose -f ./deployments/e2e/docker-compose.yaml down
}

function status(){
    docker-compose -f ./deployments/e2e/docker-compose.yaml ps
}

function clean(){
    stop
    docker rmi -f ${APP_IMAGE_TAG}:${APP_IMAGE_TAG}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "start")
        start
        ;;
    "status")
        status
        ;;
    "stop")
        stop
        ;;
    "clean")
        clean
        ;;
    *)
        echo "$0 [ start | stop | clean ]"
        ;;
esac