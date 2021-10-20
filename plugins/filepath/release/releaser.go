package release

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/platform"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/utils"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	sdk "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type ReleaseConfig struct {
}

type Releaser struct {
	config ReleaseConfig
}

// Config ensures that Builder implements the method from the Configurable interface
func (d *Releaser) Config() (interface{}, error) {
	return &d.config, nil
}

func (r *Releaser) ReleaseFunc() interface{} {
	return r.release
}

func (r *Releaser) DestroyFunc() interface{} {
	return r.destroy
}

func (r *Releaser) StatusFunc() interface{} {
	return r.status
}

func (r *Releaser) release(
	ctx context.Context,
	ui terminal.UI,
	artifact *platform.Deployment,
) (*Release, error) {

	st := ui.Status()
	defer st.Close()

	od := utils.Directory(artifact.Path)
	rf := filepath.Join(od, "release")
	err := utils.CreateSimlink(artifact.Path, rf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create release: %s", err)
	}

	st.Step(terminal.StatusOK, fmt.Sprintf("Created release %s", rf))

	return &Release{Url: rf}, nil
}

func (r *Releaser) destroy(
	ctx context.Context,
) error {
	// TODO(briancain): write me
	return nil
}

func (r *Releaser) status(
	ctx context.Context,
	ji *component.JobInfo,
	release *Release,
	ui terminal.UI,
) (*sdk.StatusReport, error) {
	sg := ui.StepGroup()
	s := sg.Add("Checking the status of the file...")

	report := &sdk.StatusReport{}
	if _, err := os.Stat(release.Url); err == nil {
		s.Update("Symlink exists!")
		report.Health = sdk.StatusReport_READY
	} else {
		st := ui.Status()
		defer st.Close()
		st.Step(terminal.StatusError, "Symlink to File is missing!")
		report.Health = sdk.StatusReport_MISSING
	}
	s.Done()

	return report, nil
}

// ensure Releaser implements component.Release
func (r *Release) URL() string {
	return r.Url
}
