package builder

import (
	"context"
	"os/exec"
	"path"

	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

// Builder is a concrete implementation of the Waypoint Builder interface
type Builder struct {
	config BuildConfig
}

// Config ensures that Builder implements the method from the Configurable interface
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// BuildFunc is called when waypoint runs the Build step
func (b *Builder) BuildFunc() interface{} {
	return b.build
}

func (b *Builder) build(
	ctx context.Context,
	ui terminal.UI,
	src *component.Source,
) (*Binary, error) {

	st := ui.Status()
	defer st.Close()

	st.Update("Building application")

	// setup the defaults
	if b.config.OutputName == "" {
		b.config.OutputName = src.App
	}

	if b.config.Source == "" {
		b.config.Source = "./"
	}

	c := exec.Command(
		"go",
		"build",
		"-o",
		b.config.OutputName,
		b.config.Source,
	)

	err := c.Run()

	if err != nil {
		st.Step(terminal.StatusError, "Build failed")

		return nil, err
	}

	st.Step(terminal.StatusOK, "Application built successfully")

	return &Binary{
		Path: path.Join(b.config.Source, b.config.OutputName),
	}, nil
}

// BuildConfig defines a struct which will hold the serialzed HCL configuration
type BuildConfig struct {
	OutputName string `hcl:"output_name,optional"`
	Source     string `hcl:"source,optional"`
}
