package builder

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/datadir"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// Builder is a concrete implementation of the Waypoint Builder interface
type Builder struct {
	config BuildConfig
}

// Config ensures that Builder implements the method from the Configurable interface
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

func (b *Builder) ConfigurableNotify(config interface{}) error {
	c, ok := config.(*BuildConfig)
	if !ok {
		return fmt.Errorf("Expected type BuildConfig")
	}

	// validate the config
	_, err := os.Stat(c.Source)
	if err != nil {
		return fmt.Errorf("Source folder does not exist")
	}

	// config validated ok
	return nil
}

// BuildFunc is called when waypoint runs the Build step
func (b *Builder) BuildFunc() interface{} {
	return b.build
}

func (b *Builder) build(
	ctx context.Context,
	src *component.Source,
	job *component.JobInfo,
	projectDir *datadir.Project,
	appDir *datadir.App,
	componentDir *datadir.Component,
	log hclog.Logger,
	ui terminal.UI,
	labels *component.LabelSet,
) (*Binary, error) {

	st := ui.Status()
	defer st.Close()

	log.Info(
		"Start build",
		"src", src,
		"job", job,
		"projectDataDir", projectDir.DataDir(),
		"projectCacheDir", projectDir.CacheDir(),
		"appDataDir", appDir.DataDir(),
		"appCacheDir", appDir.CacheDir(),
		"componentDataDir", componentDir.DataDir(),
		"componentCacheDir", componentDir.CacheDir(),
		"labels", labels,
	)

	st.Update("Building application")

	// setup the defaults
	if b.config.OutputName == "" {
		b.config.OutputName = src.App
	}

	if b.config.Source == "" {
		b.config.Source = "./"
	}

	outputFile := filepath.Join(b.config.Source, b.config.OutputName)

	c := exec.Command(
		"go",
		"build",
		"-o",
		outputFile,
	)

	err := c.Run()

	if err != nil {
		st.Step(terminal.StatusError, "Build failed")

		return nil, err
	}

	st.Step(terminal.StatusOK, "Application built successfully "+outputFile)

	return &Binary{
		Path: outputFile,
	}, nil
}

// BuildConfig defines a struct which will hold the serialzed HCL configuration
type BuildConfig struct {
	OutputName string `hcl:"output_name,optional"`
	Source     string `hcl:"source,optional"`
}
