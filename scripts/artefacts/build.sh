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

native_build_image=${IMAGE_BASE_NAME}/native_build_image:current

function packageContainer() {
    docker build -f ./build/package/artefacts/container.dockerfile \
            --build-arg APP_NAME=${APP_NAME} \
            --build-arg NODE_VER=${NODE_VER} \
            --build-arg GO_VER=${GO_VER} \
            --build-arg WEB_FRAMEWORK=${WEB_FRAMEWORK} \
            -t ${APP_IMAGE_NAME}:${APP_IMAGE_TAG} .
}

function packageNative(){
    docker build -f ./build/package/artefacts/native.dockerfile \
                --build-arg APP_NAME=${APP_NAME} \
                --build-arg NODE_VER=${NODE_VER} \
                --build-arg GO_VER=${GO_VER} \
                --build-arg WEB_FRAMEWORK=${WEB_FRAMEWORK} \
                -t ${native_build_image} .
    cleanNative
    id=$(docker create ${native_build_image})
    CONTAINER_ID="${id:0:12}"
    mkdir -p ./build/native/linux
    docker cp $CONTAINER_ID:/opt/build/package/linux/ ./build/native
    mkdir -p ./build/native/macOS
    docker cp $CONTAINER_ID:/opt/build/package/macOS/ ./build/native
    mkdir -p ./build/native/windows
    docker cp $CONTAINER_ID:/opt/build/package/windows/ ./build/native
    docker rm -f $id
    docker rmi -f ${native_build_image}
}

function cleanNative() {
    if [ -d ./build/native ]; then
        rm -rf ./build/native
    fi
}

function cleanImages() {
    docker rmi -f ${APP_IMAGE_NAME}:${APP_IMAGE_TAG}
    docker rmi -f ${test_build_image}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

message="$0 clean | container | native "

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "clean")
        cleanNative
        cleanImages
        ;;
    "container")
        packageContainer
        ;;
    "native")
        packageNative
        ;;
    *)
        echo $message
        ;;
esac