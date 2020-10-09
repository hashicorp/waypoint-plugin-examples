# Waypoint Plugin Template

This folder contains an example plugin structure which can be used when building your own plugins.

## Steps

1. To scaffold a new plugin use the `./clone.sh` script passing the destination folder and the Go package
for your new plugin as parameters

```shell
./clone.sh ../destination_folder github.com/myorg/mypackage
```

2. You can then run the Makefile to compile the new plugin

```shell
cd ../destination_folder

make
```

```shell
Build Protos
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./builder/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./registry/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./platform/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./release/output.proto

Compile Plugin
go build -o ./bin/waypoint-plugin-template ./main.go
```