package gotado

import (
	"sync"

	"golang.org/x/oauth2"
)

// TokenRefreshCallback is called whenever a token is automatically refreshed.
// The callback receives the new token and should persist it to storage.
//
// The token passed to the callback contains the essential fields needed for
// persistence (AccessToken, RefreshToken, TokenType, Expiry, ExpiresIn).
// Extra fields stored via WithExtra() are intentionally not included to avoid
// potential race conditions from sharing references to the original token's
// internal data structures.
//
// IMPORTANT: The callback is called synchronously and should return quickly.
// If heavy processing is needed, consider sending the token to a channel
// or queue for asynchronous processing.
type TokenRefreshCallback func(token *oauth2.Token)

// callbackTokenSource wraps an oauth2.TokenSource and calls a callback
// whenever Token() returns a different token than the previous call.
// This is useful for persisting refreshed OAuth2 tokens to disk or other storage.
type callbackTokenSource struct {
	src       oauth2.TokenSource
	callback  TokenRefreshCallback
	mu        sync.Mutex
	lastToken *oauth2.Token
}

// Token implements the oauth2.TokenSource interface.
// It retrieves a token from the underlying source and invokes the callback
// if the token has changed (either access token or refresh token is different).
func (cts *callbackTokenSource) Token() (*oauth2.Token, error) {
	cts.mu.Lock()
	defer cts.mu.Unlock()

	newToken, err := cts.src.Token()
	if err != nil {
		return nil, err
	}

	// Check if token has changed (different access token or refresh token)
	// We check both because:
	// - Access tokens expire frequently (every 10 minutes for tado°)
	// - Refresh tokens may rotate (tado° uses refresh token rotation)
	tokenChanged := false
	if cts.lastToken == nil {
		tokenChanged = true
	} else {
		// Compare access tokens
		if cts.lastToken.AccessToken != newToken.AccessToken {
			tokenChanged = true
		}
		// Compare refresh tokens (important for token rotation)
		if cts.lastToken.RefreshToken != newToken.RefreshToken {
			tokenChanged = true
		}
	}

	// Update lastToken if token changed
	if tokenChanged {
		cts.lastToken = copyToken(newToken)

		// Invoke callback if provided
		if cts.callback != nil {
			// Make a copy of the token to pass to the callback
			// This prevents the callback from modifying the token
			tokenCopy := copyToken(newToken)
			cts.callback(tokenCopy)
		}
	}

	return newToken, nil
}

// NewCallbackTokenSource creates a TokenSource that invokes the provided
// callback whenever the underlying token is refreshed.
//
// This is particularly useful for the tado° API which:
//   - Has short-lived access tokens (10 minutes)
//   - Uses refresh token rotation (old refresh token is invalidated when new one is issued)
//   - Requires offline_access scope for refresh tokens
//
// Example usage:
//
//	config := gotado.AuthConfig(clientID, "offline_access")
//	token, _ := config.DeviceAccessToken(ctx, deviceAuth)
//
//	callback := func(newToken *oauth2.Token) {
//	    // Save token to encrypted storage
//	    log.Println("Token refreshed, saving to disk")
//	    SaveTokenToFile(newToken)
//	}
//
//	tado := gotado.NewWithTokenRefreshCallback(ctx, config, token, callback)
func NewCallbackTokenSource(src oauth2.TokenSource, callback TokenRefreshCallback) oauth2.TokenSource {
	return &callbackTokenSource{
		src:      src,
		callback: callback,
	}
}

// copyToken creates a copy of an oauth2.Token.
// This creates a new token with the same field values as the source token.
//
// Note: Extra fields stored via WithExtra() are intentionally not copied.
// The purpose of the callback is to persist the access and refresh tokens
// for later use. Extra fields are not needed for token storage and excluding
// them avoids potential race conditions from sharing references to the
// original token's internal data structures.
func copyToken(src *oauth2.Token) *oauth2.Token {
	if src == nil {
		return nil
	}

	return &oauth2.Token{
		AccessToken:  src.AccessToken,
		TokenType:    src.TokenType,
		RefreshToken: src.RefreshToken,
		Expiry:       src.Expiry,
		ExpiresIn:    src.ExpiresIn,
	}
}
