package gotado

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	oauth2int "github.com/gonzolino/gotado/v2/internal/oauth2"
	"golang.org/x/oauth2"
)

const (
	authURL  = "https://auth.tado.com/oauth/authorize"
	tokenURL = "https://auth.tado.com/oauth/token"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// client to access the tado° API
type client struct {
	// ClientID specifies the client ID to use for authentication
	ClientID string
	// ClientSecret specifies the client secret to use for authentication
	ClientSecret string

	http HTTPClient
}

// newClient creates a new tado° client
func newClient(clientID, clientSecret string) *client {
	return &client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		http:         http.DefaultClient,
	}
}

// WithHTTPClient configures the http client to use for tado° API interactions
func (c *client) WithHTTPClient(httpClient *http.Client) *client {
	c.http = httpClient
	return c
}

// WithCredentials sets the given credentials and scopes for the tado° API
func (c *client) WithCredentials(ctx context.Context, username, password string) (*client, error) {
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
func (c *client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to talk to tado° API: %w", err)
	}
	return resp, nil
}

// Request performs an HTTP request to the tado° API
func (c *client) Request(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create http request: %w", err)
	}
	return c.Do(req)
}

// RequestWithHeaders performs an HTTP request to the tado° API with the given map of HTTP headers
func (c *client) RequestWithHeaders(ctx context.Context, method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create http request: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return c.Do(req)
}

// get retrieves an object from the tado° API.
func (c *client) get(ctx context.Context, url string, v interface{}) error {
	resp, err := c.Request(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return fmt.Errorf("tado° API error: %w", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return nil
}

// post sends a post request to the tado° API.
func (c *client) post(ctx context.Context, url string) error {
	resp, err := c.Request(ctx, http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	if err := isError(resp); err != nil {
		return fmt.Errorf("tado° API error: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected tado° API response status: %s", resp.Status)
	}
	return nil
}

// put updates an object on the tado° API.
// If the update is successful and v is a pointer, put will decode the response
// body into the value pointed to by v. If v is not a pointer the response body
// will be ignored.
func (c *client) put(ctx context.Context, url string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("unable to marshal object: %w", err)
	}
	resp, err := c.RequestWithHeaders(ctx, http.MethodPut, url, bytes.NewReader(data),
		map[string]string{"Content-Type": "application/json;charset=utf-8"})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return fmt.Errorf("tado° API error: %w", err)
	}

	// If v is not a pointer, ignore the response body
	if rv := reflect.ValueOf(v); rv.Kind() != reflect.Ptr {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("unable to decode tado° API response: %w", err)
	}
	return nil
}

// delete deletes an object from the tado° API.
func (c *client) delete(ctx context.Context, url string) error {
	resp, err := c.Request(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if err := isError(resp); err != nil {
		return fmt.Errorf("tado° API error: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected tado° API response status: %s", resp.Status)
	}
	return nil
}
