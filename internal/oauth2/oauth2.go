package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

//go:generate mockgen -source oauth2.go -package oauth2 -destination oauth2_mock.go

// ConfigInterface codifies the interface of an oauth2.Config object.
// This is relevant for testing purposes to mock the OAuth2 authentication flow.
type ConfigInterface interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	PasswordCredentialsToken(ctx context.Context, username, password string) (*oauth2.Token, error)
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Client(ctx context.Context, t *oauth2.Token) *http.Client
	TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource
}

type oauth2Config struct {
	config *oauth2.Config
}

// NewConfig creates a new oauth2 config object
var NewConfig = func(clientID, clientSecret, authURL, tokenURL string, scopes []string) ConfigInterface {
	return &oauth2Config{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   authURL,
				TokenURL:  tokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
			Scopes: scopes,
		},
	}
}

func (oc oauth2Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return oc.config.AuthCodeURL(state, opts...)
}

func (oc oauth2Config) PasswordCredentialsToken(ctx context.Context, username, password string) (*oauth2.Token, error) {
	return oc.config.PasswordCredentialsToken(ctx, username, password)
}

func (oc oauth2Config) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return oc.config.Exchange(ctx, code, opts...)
}

func (oc oauth2Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oc.config.Client(ctx, t)
}

func (oc oauth2Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return oc.config.TokenSource(ctx, t)
}
