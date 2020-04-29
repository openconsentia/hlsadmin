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

FROM golang:1.13.3

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download && \
    go test -v hls-devkit/hlsadmin/internal/usersmgmt/authuser && \
    go test -v hls-devkit/hlsadmin/cmd/hlsadmin/cli
