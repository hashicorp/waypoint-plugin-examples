package builder

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type BuildConfig struct {
	OutputName string  `hcl:"output_name,optional"`
	Source     string  `hcl:"source,optional"`
	Lambda     *Lambda `hcl:"lambda,block"`
}
type Lambda struct {
	Amd64 bool `hcl:"amd64,optional"`
	Linux bool `hcl:"linux,optional"`
}

type Builder struct {
	config BuildConfig
}

// Implement Configurable
func (b *Builder) Config() interface{} {
	return &b.config
}

// Implement ConfigurableNotify
func (b *Builder) ConfigSet(config interface{}) error {
	c, ok := config.(*BuildConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *BuildConfig as parameter")
	}

	// validate the config
	_, err := os.Stat(c.Source)
	if err != nil {
		return fmt.Errorf("Source folder does not exist")
	}
	return nil
}

// Implement Builder
func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.build
}

// A BuildFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - *datadir.Project
// - *datadir.App
// - *datadir.Component
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet
//
// The output parameters for BuildFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (b *Builder) build(ctx context.Context, ui terminal.UI) (*Binary, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

	if b.config.OutputName == "" {
		b.config.OutputName = "app"
	}

	if b.config.Source == "" {
		b.config.Source = "./"
	}
	c := BuildCommand(b)
	err := c.Run()
	if err != nil {
		u.Step(terminal.StatusError, "Build failed")
		return nil, err
	}
	u.Step(terminal.StatusOK, "Application")
	return &Binary{
		Location: path.Join(b.config.Source, b.config.OutputName),
	}, nil
}
