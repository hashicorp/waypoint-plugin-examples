package release

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-examples/template/registry"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// Implement the Destroyer interface
func (rm *ReleaseManager) DestroyFunc() interface{} {
	return rm.destroy
}

// A DestroyFunc does not have a strict signature, you can define the parameters
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
// In addition to default input parameters the Deployment from the DeployFunc step
// can also be injected.
//
// The output parameters for PushFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (rm *ReleaseManager) destroy(
	ctx context.Context,
	log hclog.Logger,
	ui terminal.UI,
	release *Release,
) error {
	sg := ui.StepGroup()
	defer sg.Wait()

	r := rm.resourceManager(log, nil)

	// If we don't have resource state, this state is from an older version
	// and we need to manually recreate it.
	if release.ResourceState == nil {
		r.Resource("release").SetState(&Resource_Release{
			Name: release.Name,
		})
	} else {
		// Load our set state
		if err := r.LoadState(release.ResourceState); err != nil {
			return err
		}
	}

	// Destroy
	return r.DestroyAll(ctx, log, sg, ui)
}

func (rm *ReleaseManager) resourceReleaseDestroy(
	ctx context.Context,
	log hclog.Logger,
	st terminal.Status,
	ui terminal.UI,
	artifact *registry.Artifact,
	result *Release,
) error {
	// Create your deployment resource here!

	return nil
}
