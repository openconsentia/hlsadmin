# Overview

This a template to help you scaffold a project to create macOS, Linux and Windows native app or Docker base app that combine Web UI framework (e.g. ReactJS, Vue, etc) in a single package.

This project is **NOT** a Go module or library where you would `go get` packages for inclusion into your project.

The scaffold generated from this template includes a sample app. The app is intended to demonstrates steps involve in building functional apps. The sample app does **NOT** represent any specific use case but you can extend it into a fully functional production release or deployment. You may also elect to replace the sample app with one configured to work within the context of the scaffold.

## How do I scaffold a project from this template?

At the top this page, you should see a green coloured button named `Use this template`. Click on the button and follow the instructions from Github.

## What do I get out-of-the-box when I create a scaffold?

You'll get the following:

1. A ReactJS sub-project

2. Go codes to manage app configuration, startup sequence and network operations

3. Go implementation of a webserver and RESTful server

4. Build scripts -- Docker based -- to create native (macOS, Linux and Windows) app or Docker image app.

5. A locally deployable development environment

## Prerequisite

In order to build apps based on this template ensure you have the following items installed:

1. Docker and Docker compose

2. [Go toolkit version 1.11 onwards](https://blog.golang.org/)

3. Node

## How do I see the sample app in action?

Follow these steps:

1. Build the sample app (please refer to [doc](./docs/Build.md) for instruction 

2. Open an terminal and navigate to `./build/native/<platform of your choice>`

3. Start app `goreact` (for macOS and Linux) or `goreact.exe` (for windows)
NOTE: By default the app will run require port 80 to be available. If you need the app to use other ports type the command `goreact --help`.

4. Assuming you have no problem with STEP 3, run the command `goreact frontend` open a browser with url to `localhost`.

## How do I extend/modify my scaffold?

Please refer to:

* [Architecture](./docs/Arch.md)

* [Build](./docs/Build.md)
 
