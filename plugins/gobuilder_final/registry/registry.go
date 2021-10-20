package registry

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final/builder"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type RegistryConfig struct {
	Name    string "hcl:name"
	Version string "hcl:version"
}

type Registry struct {
	config RegistryConfig
}

// Implement Configurable
func (r *Registry) Config() interface{} {
	return &r.config
}

// Implement ConfigurableNotify
func (r *Registry) ConfigSet(config interface{}) error {
	c, ok := config.(*RegistryConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *RegisterConfig as parameter")
	}

	// validate the config
	if c.Name == "" {
		return fmt.Errorf("Name must be set to a valid directory")
	}

	return nil
}

// Implement Registry
func (r *Registry) PushFunc() interface{} {
	// return a function which will be called by Waypoint
	return r.push
}

func (r *Registry) AccessInfoFunc() interface{} {
	return r.accessInfo
}

func (r *Registry) accessInfo() (*AccessInfo, error) {
	return &AccessInfo{}, nil
}

// A PushFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet
//
// In addition to default input parameters the builder.Binary from the Build step
// can also be injected.
//
// The output parameters for PushFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (r *Registry) push(
	ctx context.Context,
	ui terminal.UI,
	log hclog.Logger,
	binary *builder.Binary,
) (*Artifact, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Pushing binary to registry")

	return &Artifact{}, nil
}

// Implement Authenticator

var _ component.Registry = (*Registry)(nil)
var _ component.RegistryAccess = (*Registry)(nil)
