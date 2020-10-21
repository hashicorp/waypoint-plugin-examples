package main

import (
	"github.com/sshintaku/waypoint-plugin-examples/gobuilder/builder"
	"github.com/sshintaku/waypoint-plugin-examples/gobuilder/platform"
	"github.com/sshintaku/waypoint-plugin-examples/gobuilder/registry"
	"github.com/sshintaku/waypoint-plugin-examples/gobuilder/release"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	// sdk.Main allows you to register the components which should
	// be included in your plugin
	// Main sets up all the go-plugin requirements

	sdk.Main(sdk.WithComponents(
		// Comment out any components which are not
		// required for your plugin
		&builder.Builder{},
		&registry.Registry{},
		&platform.Platform{},
		&release.ReleaseManager{},
	))
}
