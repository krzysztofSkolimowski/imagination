# imagination
Simple service, capable of performing basic image transformations.
### Project is currently under development

## Installation and running

### Requirements

1. docker
1. docker-compose
1. Go - it is necessary to download dep and wire for now however project uses it's own version. In the future the only requirements will be docker-compose and docker.
1. dep - `go get -u github.com/golang/dep/cmd/dep`
1. wire - `go get github.com/google/go-cloud/wire/cmd/wire`

### Running for the first time

1. Ensure you have proper set environment variables on your system
```go env```
```GOBIN```
```GOENV```
```GOPATH```
1. Run ```go get github.com/krzysztofSkolimowski/imagination``` it will clone the repository, you can clone the repository, however remember to clone it at correct path: ```$GOPATH/src/github.com/krzysztofSkolimowski/imagination```, otherwise go will not find relative paths
1. image will work on dependencies provided by go dep, which will be located in vendor directory after `$GOBIN/dep ensure` in the project directory
1. ```Make up``` should start the project with config loaded from ```.env```
#### Dependency injection
1. Project uses `go-cloud/wire` dependency injection generation tool. Check if you have wire in you ```$GOBIN```, if not run ```go get github.com/google/go-cloud/wire/cmd/wire```.
1. To regenerate `wire_gen.go` run:
```cd $GOPATH/src/github.com/krzysztofSkolimowski/cmd/modules```
```$GOBIN/wire```


## Application structure
In order to minimize number of external dependencies, project follows rules of clean architecture and consists of three layers: 
1. ```app```  - application - layer responsible for conducting actual processing and place of bussiness logic
1. ```infrastructure``` Layer responsible for accessing every external services and types of storage from application
1. ```interfaces``` Layer responsible for providing access to our service e.g. place to declare all kind of API

Outside of this layers there exists common, which is a place to cover all external libraries with proper abstraction. At this point there still remains some of them (from earlier stage of development - e. g. `logrus.Loggger`), but all of external libraries should be moved to common in the future. 

Moreover outside of layer also exists package cmd, which serves the purpose of providing necessary commands to start the application (and probably conduct scheduled jobs in the future).