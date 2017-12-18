Proof of concept for a simple gRPC setup with servers and clients written in multiple languages

# Contents
* [requirements](#requirements)
* [build](#build)
* [run](#run)

# Requirements
## Go SDK
[Go](https://golang.org/) >= 1.6 has to be installed in your machine

`go env GOPATH` and `go env GOBIN` must point to directories in your filesystem. `go env GOBIN` must be part of your `$PATH`.

Easiest way to get this working is to set the following environment variables in your shell init scripts (`.basrc`, `.zshrc`, ...)

```bash
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

## Go packages
### dep
Provides dependency management for the project
```
go get -u github.com/golang/dep/cmd/dep
```

### grpc
Provides go grpc package, Protocol Buffers compiler (protoc) and php code generator source code
```
go get -u google.golang.org/grpc
```
### proto-gen-go
Providers go code generator for grpc services and clients
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

### protoc-gen-grpc-gateway
Provides go code generator for grpc gateway
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

### protoc-gen-swagger
Provides swagger files generator for grpc-gateway
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
# PHP
[php](http://php.net/) >= 5.6 has to be installed in your machine


# Docker
[docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) have to be installed in your machine

# Build
```
make all
```
- Compiles php code generator plugin
- Generates Go and php code from `.proto` files
- Compiles Go `server`, `client` and `gw` applications
- Build docker images to run php and nodejs code

Inspect the `Makefile` for more fine grained tasks

# Run
Run one or the other...
```
# starts a GO server and a JSON gateway
docker-compose up

# starts a NODE server and a JSON gateway
docker-compose -f docker-compose-node.yml up
```

In a different terminal, experiment with the following operations...

## add recording
```
# go
bin/client --command=add --recording "<name of the recording>"
# php
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:php php bin/add.php "<name of the recording>"
# node
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:node node add.js "name of the recording"
# gateway
curl -X POST http://localhost:8080/v1/recordings -d '{"recording": {"name": "name of the recording", "author": { "name": "HTTP/JSON"}}}'

```

## list recordings
```
# go
bin/client --command=list
# php
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:php php bin/list.php
# node
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:node node list.js
# gateway
curl http://localhost:8080/v1/recordings
```

## list recordings stream (1000ms delay) (delay only if go server, nodejs sleep requires python)
```
# go
bin/client --command=stream
# php
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:php php bin/stream.php
# node
docker run -ti --rm --network grpcspike_default kobaltmusic/grpc-spike:node node stream.js
```
