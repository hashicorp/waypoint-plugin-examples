// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/platform"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/registry"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/release"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	// sdk.Main allows you to register the components which should
	// be included in your plugin
	// Main sets up all the go-plugin requirements

	sdk.Main(sdk.WithComponents(
		// Comment out any components which are not
		// required for your plugin
		&registry.Registry{},
		&platform.Deploy{},
		&release.Releaser{},
	))
}
