package platform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/registry"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/utils"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
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

/*
func (s *Deploy) LogsFunc() interface{} {
	// setup tailing the logs files
	return s
}

func (s *Deploy) NextLogBatch(ctx context.Context) ([]component.LogEvent, error) {

	return []component.LogEvent{{Message: "ok"}}, nil
}
*/

//func (d *Deploy) DefaultReleaserFunc() interface{} {
//	return func() *Deploy { return &release.Releaser{} }
//}

func (d *Deploy) deploy(
	ctx context.Context,
	ji *component.JobInfo,
	artifact *registry.Artifact,
	ui terminal.UI,
) (*Deployment, error) {

	st := ui.Status()
	defer st.Close()
	st.Update(fmt.Sprintf("Deploying version %s", ji.Id))

	//deps, err := hc.Deployments(ctx, &history.Lookup{FilterStatus: history.StatusSuccess})
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "Unable to query deployment history: %s", err)
	//}

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
