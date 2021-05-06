package gotado

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	oauth2int "github.com/gonzolino/gotado/internal/oauth2"
	"golang.org/x/oauth2"
)

const (
	authURL  = "https://auth.tado.com/oauth/authorize"
	tokenURL = "https://auth.tado.com/oauth/token"
)

// Client to access the tado° API
type Client struct {
	// ClientID specifies the client ID to use for authentication
	ClientID string
	// ClientSecret specifies the client secret to use for authentication
	ClientSecret string

	http *http.Client
}

// NewClient creates a new tado° client
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		http:         http.DefaultClient,
	}
}

// WithTimeout configures the tado° object with the given timeout for HTTP requests.
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.http.Timeout = timeout
	return c
}

// WithCredentials sets the given credentials and scopes for the tado° API
func (c *Client) WithCredentials(ctx context.Context, username, password string) (*Client, error) {
	config := oauth2int.NewConfig(c.ClientID, c.ClientSecret, authURL, tokenURL, []string{"home.user"})

	httpContext := context.WithValue(ctx, oauth2.HTTPClient, c.http)
	token, err := config.PasswordCredentialsToken(httpContext, username, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}
	authClient := config.Client(httpContext, token)

	c.http = authClient

	return c, nil
}

// Do sends the given HTTP request to the tado° API.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to talk to tado° API: %w", err)
	}
	return resp, nil
}

// Request performs an HTTP request to the tado° API
func (c *Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create http request: %w", err)
	}
	return c.Do(req)
}
