package gotado

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	oauth2int "github.com/gonzolino/gotado/internal/oauth2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestWithTimeout(t *testing.T) {
	client := NewClient("test", "test")

	assert.Zero(t, client.http.Timeout)

	client.WithTimeout(1)
	assert.Equal(t, time.Duration(1), client.http.Timeout)
}

func TestWithCredentials(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	config := oauth2.Config{}
	mockConfig := oauth2int.NewMockConfigInterface(ctrl)
	oauth2int.NewConfig = func(clientID, clientSecret, authURL, tokenURL string, scopes []string) oauth2int.ConfigInterface {
		return mockConfig
	}
	forbiddenError := &url.Error{}

	token := &oauth2.Token{
		AccessToken:  "access_token",
		TokenType:    "token_type",
		RefreshToken: "refresh_token",
		Expiry:       time.Now(),
	}

	client := NewClient("test", "test")
	httpCtx := context.WithValue(ctx, oauth2.HTTPClient, client.http)

	mockConfig.EXPECT().PasswordCredentialsToken(gomock.AssignableToTypeOf(httpCtx), "username", "password").Return(token, nil)
	mockConfig.EXPECT().Client(httpCtx, token).Return(config.Client(httpCtx, token))

	_, err := client.WithCredentials(ctx, "username", "password")
	assert.NoError(t, err)

	mockConfig.EXPECT().PasswordCredentialsToken(gomock.AssignableToTypeOf(httpCtx), "username", gomock.Not("password")).Return(nil, forbiddenError)

	_, err = client.WithCredentials(ctx, "username", "wrong")
	assert.Exactly(t, fmt.Errorf("invalid credentials: %w", forbiddenError), err)
}
