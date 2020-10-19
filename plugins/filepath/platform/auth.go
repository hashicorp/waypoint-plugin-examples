package platform

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// ValidateAuthFunc satisfies the Authenticator interface
func (p *Deploy) ValidateAuthFunc() interface{} {
	return p.validateAuth
}

// AuthFunc satisfies the Authenticator interface
func (p *Deploy) AuthFunc() interface{} {
	return p.authenticate
}

func (p *Deploy) validateAuth(
	ctx context.Context,
	log hclog.Logger,
	ui terminal.UI,
) error {
	s := ui.Status()
	defer s.Close()

	s.Update("Validate authentication")

	// returning an error from ValidateAuthFunc causes Waypoint
	// to call AuthFunc
	return nil
}

func (p *Deploy) authenticate(
	ctx context.Context,
	log hclog.Logger,
	ui terminal.UI,
) (*component.AuthResult, error) {

	ui.Output("Describe the manual authentication steps here")

	return &component.AuthResult{Authenticated: true}, nil
}
