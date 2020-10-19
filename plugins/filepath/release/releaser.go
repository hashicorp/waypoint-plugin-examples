package release

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/platform"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/utils"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil
}

// ensure Releaser implements component.Release
func (r *Release) URL() string {
	return r.Url
}
