package gotado

import (
	"context"

	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:       "https://login.tado.com/oauth2/authorize",
	TokenURL:      "https://login.tado.com/oauth2/token",
	DeviceAuthURL: "https://login.tado.com/oauth2/device_authorize",
}

type Tado struct {
	client *client
}

// AuthConfig creates a new OAuth2 config to be used for authentication with the tado° API.
// Add `"offline_access"` as a scope if you need a refresh token that can be used to regularly create new access tokens.
func AuthConfig(clientID string, scopes ...string) *oauth2.Config {
	return &oauth2.Config{
		ClientID: clientID,
		Endpoint: Endpoint,
		Scopes:   scopes,
	}
}

// New creates a new tado client.
func New(ctx context.Context, config *oauth2.Config, token *oauth2.Token) *Tado {
	return &Tado{
		client: newClient(ctx, config, token),
	}
}

// Me returns information about the authenticated user.
func (t *Tado) Me(ctx context.Context) (*User, error) {
	me := &User{client: t.client}
	if err := t.client.get(ctx, apiURL("me"), me); err != nil {
		return nil, err
	}
	return me, nil
}
