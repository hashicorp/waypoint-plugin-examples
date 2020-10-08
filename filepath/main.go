// Package files contains a component for validating local files.
package main

import (
	"github.com/hashicorp/waypoint-plugin-examples/filepath/deploy"
	"github.com/hashicorp/waypoint-plugin-examples/filepath/registry"
	"github.com/hashicorp/waypoint-plugin-examples/filepath/release"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

//go:generate protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./registry/plugin.proto
//go:generate protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./deploy/plugin.proto
//go:generate protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./release/plugin.proto

func main() {
	sdk.Main(sdk.WithComponents(
		&registry.Registry{},
		&deploy.Deploy{},
		&release.Releaser{},
	))
}
