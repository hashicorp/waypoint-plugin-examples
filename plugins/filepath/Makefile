PLUGIN_NAME=filepath

all: protos build

protos:
	@echo ""
	@echo "Build Protos"

	protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./registry/output.proto
	protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./platform/output.proto
	protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./release/output.proto

build:
	@echo ""
	@echo "Compile Plugin"

	go build -o ./bin/waypoint-plugin-${PLUGIN_NAME} ./main.go

install:
	@echo ""
	@echo "Installing Plugin"

	cp ./bin/waypoint-plugin-${PLUGIN_NAME} ${HOME}/.config/waypoint/plugins/
