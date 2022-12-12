// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package platform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/registry"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/utils"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	sdk "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type Deploy struct {
	config DeployConfig
}

type DeployConfig struct {
	Directory string `hcl:"directory"`
}

// Config ensures that Builder implements the method from the Configurable interface
func (d *Deploy) Config() (interface{}, error) {
	return &d.config, nil
}

func (d *Deploy) DeployFunc() interface{} {
	return d.deploy
}

func (d *Deploy) DestroyFunc() interface{} {
	return d.destroy
}

func (d *Deploy) StatusFunc() interface{} {
	return d.status
}

func (d *Deploy) deploy(
	ctx context.Context,
	ji *component.JobInfo,
	artifact *registry.Artifact,
	ui terminal.UI,
) (*Deployment, error) {

	st := ui.Status()
	defer st.Close()
	st.Update(fmt.Sprintf("Deploying version %s", ji.Id))

	// get the current deployment number
	deployCount := utils.DeploymentCount(d.config.Directory)

	// create the deployment filename
	deploy := fmt.Sprintf(
		"%s.%d.deployment",
		utils.Filename(artifact.Path),
		deployCount+1,
	)

	deploy = filepath.Join(d.config.Directory, deploy)

	err := utils.CreateSimlink(artifact.Path, deploy)
	if err != nil {
		return nil, err
	}

	st.Step(terminal.StatusOK, fmt.Sprintf("Created deployment %s", deploy))

	return &Deployment{Path: deploy}, nil
}

func (d *Deploy) destroy(
	ctx context.Context,
	ui terminal.UI,
	deployment *Deployment,
) error {
	st := ui.Status()
	defer st.Close()

	err := os.RemoveAll(d.config.Directory)
	if err != nil {
		st.Step(terminal.ErrorStyle, fmt.Sprintf("Unable to remove deployments %s", d.config.Directory))
		return err
	}

	st.Step(terminal.StatusOK, fmt.Sprintf("Removed deployments %s", d.config.Directory))
	return nil
}

func (d *Deploy) status(
	ctx context.Context,
	ji *component.JobInfo,
	deploy *Deployment,
	ui terminal.UI,
) (*sdk.StatusReport, error) {
	sg := ui.StepGroup()
	s := sg.Add("Checking the status of the file...")

	report := &sdk.StatusReport{}
	if _, err := os.Stat(deploy.Path); err == nil {
		s.Update("File is ready!")
		report.Health = sdk.StatusReport_READY
	} else {
		st := ui.Status()
		defer st.Close()
		st.Step(terminal.StatusError, "File is missing!")
		s.Status(terminal.StatusError)

		report.Health = sdk.StatusReport_MISSING
	}
	s.Done()

	return report, nil
}
