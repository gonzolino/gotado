package api

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"

	"github.com/gonzolino/gotado/api/client"
	"github.com/gonzolino/gotado/api/client/home"
	oauth2int "github.com/gonzolino/gotado/api/internal/oauth2"
)

//go:generate swagger generate client -f ../swagger.yaml

const (
	clientIDEnvVar     = "TADO_CLIENT_ID"
	clientSecretEnvVar = "TADO_CLIENT_SECRET"
	usernameEnvVar     = "TADO_USERNAME"
	passwordEnvVar     = "TADO_PASSWORD"
	authURL            = "https://auth.tado.com/oauth/authorize"
	tokenURL           = "https://auth.tado.com/oauth/token"
)

var (
	scopes = []string{"home.user"}
)

// API is the interface for the Tado API
type API struct {
	client       *client.TadoAPI
	clientID     string
	clientSecret string

	Home        home.ClientService
	Transport   runtime.ClientTransport
	BearerToken *runtime.ClientAuthInfoWriter
}

// NewAPI creates a new API client with the given ID and secret
func NewAPI(_ context.Context, clientID, clientSecret string) *API {
	api := &API{
		client:       client.Default,
		clientID:     clientID,
		clientSecret: clientSecret,
		BearerToken:  nil,
	}
	api.Home = api.client.Home
	api.Transport = api.client.Transport
	return api
}

func (api *API) WithAuthentication(ctx context.Context, username, password string) error {
	config := oauth2int.NewConfig(api.clientID, api.clientSecret, authURL, tokenURL, scopes)
	httpContext := context.WithValue(ctx, oauth2.HTTPClient, http.DefaultClient)
	token, err := config.PasswordCredentialsToken(httpContext, username, password)
	if err != nil {
		return err
	}
	bearerToken := httptransport.BearerToken(token.AccessToken)
	api.BearerToken = &bearerToken
	return nil
}
