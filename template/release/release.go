package release

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-examples/template/registry"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/framework/resource"
	sdk "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReleaseConfig struct {
	Active bool "hcl:directory,optional"
}

type ReleaseManager struct {
	config ReleaseConfig
}

// Implement Configurable
func (rm *ReleaseManager) Config() (interface{}, error) {
	return &rm.config, nil
}

// Implement ConfigurableNotify
func (rm *ReleaseManager) ConfigSet(config interface{}) error {
	_, ok := config.(*ReleaseConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	// validate the config

	return nil
}

// Implement Builder
func (rm *ReleaseManager) ReleaseFunc() interface{} {
	// return a function which will be called by Waypoint
	return rm.release
}

func (rm *ReleaseManager) StatusFunc() interface{} {
	return rm.status
}

// This function can be implemented to return various connection info required
// to connect to your given platform for Resource Manager. It could return
// a struct with client information, what namespace to connect to, a config,
// and so on.
func (rm *ReleaseManager) getConnectContext() (interface{}, error) {
	return nil, nil
}

// Resource manager will tell the Waypoint Plugin SDK how to create and delete
// certain resources for your deployments.
//
// For example, your deployment might need to create a "container" or "load balancer".
// Your plugin could implement two resources through ResourceManager and the Waypoint
// Plugin SDK will automatically create or delete these resources as well as
// obtain the defined status for them.
//
// ResourceManager can also be implemented for Release as well.
func (rm *ReleaseManager) resourceManager(log hclog.Logger, dcr *component.DeclaredResourcesResp) *resource.Manager {
	return resource.NewManager(
		resource.WithLogger(log.Named("resource_manager")),
		resource.WithValueProvider(rm.getConnectContext),
		resource.WithDeclaredResourcesResp(dcr),
		resource.WithResource(resource.NewResource(
			resource.WithName("template_example"),
			resource.WithState(&Resource_Release{}),
			resource.WithCreate(rm.resourceReleaseCreate),
			resource.WithDestroy(rm.resourceReleaseDestroy),
			resource.WithStatus(rm.resourceReleaseStatus),
			resource.WithPlatform("template_platform"),                                         // Update this to match your plugins platform, like Kubernetes
			resource.WithCategoryDisplayHint(sdk.ResourceCategoryDisplayHint_INSTANCE_MANAGER), // This is meant for the UI to determine what kind of icon to show
		)),
		// NOTE: Add more resource funcs here if your plugin has more than 1 resource
	)
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
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet

// In addition to default input parameters the platform.Deployment from the Deploy step
// can also be injected.
//
// The output parameters for ReleaseFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (rm *ReleaseManager) release(
	ctx context.Context,
	log hclog.Logger,
	dcr *component.DeclaredResourcesResp,
	ui terminal.UI,
	artifact *registry.Artifact,
) (*Release, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Release application")

	var result *Release

	// Create our resource manager and create deployment resources
	r := rm.resourceManager(log, dcr)

	// These params must match exactly to your resource manager functions. Otherwise
	// they will not be invoked during CreateAll()
	if err := r.CreateAll(
		ctx, log, u, ui,
		artifact, &result,
	); err != nil {
		return nil, err
	}

	// Store our resource state
	result.ResourceState = r.State()

	u.Update("Application deployed")

	return result, nil
}

func (rm *ReleaseManager) status(
	ctx context.Context,
	ji *component.JobInfo,
	log hclog.Logger,
	ui terminal.UI,
	artifact *registry.Artifact,
	release *Release,
) (*sdk.StatusReport, error) {
	sg := ui.StepGroup()
	s := sg.Add("Checking the status of the release...")

	r := rm.resourceManager(log, nil)

	// If we don't have resource state, this state is from an older version
	// and we need to manually recreate it.
	if release.ResourceState == nil {
		r.Resource("release").SetState(&Resource_Release{
			Name: release.Id,
		})
	} else {
		// Load our set state
		if err := r.LoadState(release.ResourceState); err != nil {
			return nil, err
		}
	}

	// This will call the StatusReport func on every defined resource in ResourceManager
	report, err := r.StatusReport(ctx, log, sg, ui)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "resource manager failed to generate resource statuses: %s", err)
	}

	report.Health = sdk.StatusReport_UNKNOWN
	s.Update("Release is currently not implemented")
	s.Done()

	return report, nil
}

func (rm *ReleaseManager) resourceReleaseCreate(
	ctx context.Context,
	log hclog.Logger,
	st terminal.Status,
	ui terminal.UI,
	artifact *registry.Artifact,
	result *Release,
) error {
	// Create your release resource here!

	return nil
}

func (rm *ReleaseManager) resourceReleaseStatus(
	ctx context.Context,
	ui terminal.UI,
	sg terminal.StepGroup,
	artifact *registry.Artifact,
) error {
	// Determine health status of "this" resource.
	return nil
}
