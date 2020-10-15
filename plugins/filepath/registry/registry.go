package registry

import (
	"context"
	"fmt"

	"github.com/hashicorp/waypoint-plugin-examples/plugins/filepath/utils"
	"github.com/hashicorp/waypoint-plugin-examples/plugins/gobuilder_final/builder"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Registry struct {
	config RegistryConfig
}

type RegistryConfig struct {
	Directory string `hcl:"directory"`
}

// Config ensures that Builder implements the method from the Configurable interface
func (r *Registry) Config() (interface{}, error) {
	return &r.config, nil
}

func (r *Registry) PushFunc() interface{} {
	return r.push
}

func (r *Registry) push(
	ctx context.Context,
	img *builder.Binary,
	ui terminal.UI,
) (*Artifact, error) {

	st := ui.Status()
	defer st.Close()
	st.Update("Pushing to registry")

	path, err := utils.CopyFile(img.Location, r.config.Directory)
	if err != nil {
		st.Step(terminal.StatusError, "Unable to copy file to registry")

		return &Artifact{}, status.Errorf(
			codes.Internal,
			"Unable to copy file %s to %s: %s", img.Location, r.config.Directory, err.Error(),
		)
	}

	st.Step(terminal.StatusOK, fmt.Sprintf("Application binary pushed to registry"))
	return &Artifact{Path: path}, nil
}
