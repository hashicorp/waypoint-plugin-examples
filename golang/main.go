// Package files contains a component for validating local files.
package main

import (
	"github.com/hashicorp/waypoint-plugin-examples/golang/builder"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

//go:generate protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./builder/plugin.proto

func main() {
	sdk.Main(sdk.WithComponents(
		&builder.Builder{},
	))
}
