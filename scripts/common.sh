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

export IMAGE_BASE_NAME=openconsentia

export APP_IMAGE_NAME=${IMAGE_BASE_NAME}/hlsadmin-container
export APP_IMAGE_TAG=current
<<<<<<< HEAD
export APP_NAME=hlsadmin # This should match the basename ./cmd/<youur-choice>, out-of-the-box this is call goreact
=======
export APP_NAME=goreact # This should match the basename ./cmd/<youur-choice>, out-of-the-box this is call goreact
>>>>>>> 1acca45... Bump lodash from 4.17.15 to 4.17.19 in /web/reactjs (#33)

export REACT_IMAGE_NAME=${IMAGE_BASE_NAME}/hlsadmin-dev-react
export REST_IMAGE_NAME=${IMAGE_BASE_NAME}/hlsadmin-dev-rest
export IMAGE_TAG=dev

export WEB_FRAMEWORK=reactjs

export NODE_VER=13.10.1
export GO_VER=1.13.3