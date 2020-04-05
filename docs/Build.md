# Build process

Out-of-the-box you will find scripts to help you build sample native and container based apps based 
on the included skeleton ReactJS and Go Rest API codes. If you plan to extend from the skeletons, 
you just need to extend from these scripts. If you plan to replace the skeletons, please use this
build script as inspiration.

NOTE: All scripts are BASH based and located in the folder `scripts`

## Build for development

1. Open terminal and navigate to the root of your scaffold

2. Run the command `./scripts/dev/build.sh package` to build container images for react and rest

3. Run the command `./scripts/dev/build.sh run` to start a development environment comprising of three docker containers: router, frontend and rest server (see eaxmple below)
```
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                    NAMES
80d4b66a92e6        paulwizviz/go-rest:dev    "/usr/local/bin/go-r…"   21 seconds ago      Up 19 seconds       0.0.0.0:9000->9000/tcp   rest
6536a2b8c41c        paulwizviz/go-react:dev   "docker-entrypoint.s…"   21 seconds ago      Up 19 seconds       0.0.0.0:3030->3030/tcp   react
c6c3bfb538ce        binocarlos/noxy           "bash /run.sh"           21 seconds ago      Up 19 seconds       0.0.0.0:80->80/tcp       router
```

4. Run the command `./scripts/dev/build.sh stop` to stop you development environment.

## Build for production

1. Open terminal and navigate to the root of your scaffold

### Native macOS, Linux and Windows

2. Run the command `./scripts/production/build.sh native`

3. You will find your build arfects in the folder `./build/native/<native platform>`

### Docker container based app

2. Edit the script `./scripts/production/container/build.sh` and modify the variable `$IMAGE_NAME` to one that suits your project

3. Run the command `./scripts/production/build.sh container`, where `<image tag>` is any unique string value e.g. version number

4. Run the command `docker images` and out-of-the-box, you will see the following Docker images listed:
```
paulwizviz/go-react-container   current                 273d6b7a1b46        5 minutes ago       9.07MB
```

## Clean project

Run the command `./scripts/production/build.sh clean`
