# Architecture

This project has been architect to support you, as developers, create and package your web UI and RESTful interfaces into a single macOS, Linux and Windows app or a single Docker container.

## Project layout

The layout of your scaffold is based on the [standard Go project layout](https://github.com/golang-standards/project-layout).

Out-of-the-box, you will find the following folders and these are as follows:

### `./build`

This folder includes artefacts to help you build development and production artefacts.

* [./build/package/dev](../build/package/dev) folder includes Dockerfiles to create a container of the sample ReactJS UI found in this project and a container of Go based REST server to support development activities.

* [./build/package/production](../build/package/production) folder containers Dockerfiles to produce native macOS, Linux and Windows app or a Docker image for a containerisable version of your app.

* [./build/package/go-rice.sh](../build/package/go-rice.sh) script to pre-generate a Go source file embedding the ReactJS artefacts in byte code form.

A folder named `package`, not version controlled, will be generated containing native apps for macOS, Linux and Windows platform.

### `./cmd`

This folder includes source codes to create two types of cli executables:

* `goreact` to activate a combined ReactJS and REST app;

* `gorest` to activate a standalone RESTFul server.

These executables is based on the following Go frameworks [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)

### `./deployments`

This folder contains a `docker-compose.yaml` use to deploy a local docker network comprising of three containers to support development work. The containers are:

* An nginx-styled reverse proxy known as noxy based on this [implementation](https://github.com/binocarlos/noxy)

* A react container built from this project see [./build/dev/react/Dockerfile](../build/dev/react/Dockerfile)

* A Go based RESTFul container built from this project see [./build/dev/rest/Dockerfile](../build/dev/rest/Dockerfile)

### `./internal`

This folder container source codes to build an embedded webserver package and RESTFul endpoints. The codes are packaged as follows:

* `server` containing codes for a RESTful and web servers;

* `usermgmt` containing codes for user management handler.

The name of this folder has a special impact on project scoping. Please refer to this [doc](https://blog.learngoprogramming.com/special-packages-and-directories-in-go-1d6295690a6b) for explanation before you modify it.

### `./web`

This folder contains a demonstration RectJS sub project where you have all the necessary artefacts to create a web UI. 

## Modifying/extending scaffold

TO-DO
