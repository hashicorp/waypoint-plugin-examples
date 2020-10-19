PLUGIN_NAME=gobuilder

all: protos build

protos:
	@echo ""
	@echo "Build Protos"

	protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./builder/output.proto

build:
	@echo ""
	@echo "Compile Plugin"

	go build -o ./bin/waypoint-plugin-${PLUGIN_NAME} ./main.go 

install:
	@echo ""
	@echo "Installing Plugin"

	cp ./bin/waypoint-plugin-${PLUGIN_NAME} ${HOME}/.config/waypoint/plugins/   