package gotado

import (
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// staticTokenSource returns the same token every time.
type staticTokenSource struct {
	token *oauth2.Token
}

func (s *staticTokenSource) Token() (*oauth2.Token, error) {
	return s.token, nil
}

func TestCallbackTokenSource(t *testing.T) {
	t.Run("NoCallbackWhenTokenUnchanged", func(t *testing.T) {
		token := &oauth2.Token{
			AccessToken:  "access",
			RefreshToken: "refresh",
		}

		var callCount atomic.Int32
		callback := func(newToken *oauth2.Token) {
			callCount.Add(1)
		}

		src := NewCallbackTokenSource(&staticTokenSource{token: token}, callback, token)

		// First call should NOT trigger callback since token matches initialToken
		_, err := src.Token()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if callCount.Load() != 0 {
			t.Errorf("callback should not fire when token unchanged, got %d calls", callCount.Load())
		}
	})

	t.Run("CallbackWhenNilInitialToken", func(t *testing.T) {
		token := &oauth2.Token{
			AccessToken:  "access",
			RefreshToken: "refresh",
		}

		var callCount atomic.Int32
		callback := func(newToken *oauth2.Token) {
			callCount.Add(1)
		}

		src := NewCallbackTokenSource(&staticTokenSource{token: token}, callback, nil)

		// First call should trigger callback since initialToken is nil
		_, err := src.Token()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if callCount.Load() != 1 {
			t.Errorf("callback should fire when initialToken is nil, got %d calls", callCount.Load())
		}
	})

	t.Run("CallbackWhenTokenChanges", func(t *testing.T) {
		initialToken := &oauth2.Token{
			AccessToken:  "access1",
			RefreshToken: "refresh1",
		}
		newToken := &oauth2.Token{
			AccessToken:  "access2",
			RefreshToken: "refresh2",
		}

		var callCount atomic.Int32
		callback := func(token *oauth2.Token) {
			callCount.Add(1)
		}

		src := NewCallbackTokenSource(&staticTokenSource{token: newToken}, callback, initialToken)

		// Should trigger callback since source returns a different token
		_, err := src.Token()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if callCount.Load() != 1 {
			t.Errorf("callback should fire when token changes, got %d calls", callCount.Load())
		}
	})
}

func TestCopyToken(t *testing.T) {
	t.Run("NilToken", func(t *testing.T) {
		copied := copyToken(nil)
		if copied != nil {
			t.Error("Expected nil for nil input")
		}
	})

	t.Run("BasicFields", func(t *testing.T) {
		expiry := time.Now().Add(1 * time.Hour)
		original := &oauth2.Token{
			AccessToken:  "test_access_token",
			TokenType:    "Bearer",
			RefreshToken: "test_refresh_token",
			Expiry:       expiry,
		}

		copied := copyToken(original)

		// Verify it's a different object
		if original == copied {
			t.Error("Expected different pointer, got same object")
		}

		// Verify all fields are copied
		if copied.AccessToken != original.AccessToken {
			t.Errorf("AccessToken mismatch: got %v, want %v", copied.AccessToken, original.AccessToken)
		}
		if copied.TokenType != original.TokenType {
			t.Errorf("TokenType mismatch: got %v, want %v", copied.TokenType, original.TokenType)
		}
		if copied.RefreshToken != original.RefreshToken {
			t.Errorf("RefreshToken mismatch: got %v, want %v", copied.RefreshToken, original.RefreshToken)
		}
		if !copied.Expiry.Equal(original.Expiry) {
			t.Errorf("Expiry mismatch: got %v, want %v", copied.Expiry, original.Expiry)
		}
	})

	t.Run("ExpiresInField", func(t *testing.T) {
		original := &oauth2.Token{
			AccessToken: "test_access_token",
			ExpiresIn:   600,
		}

		copied := copyToken(original)

		if copied.ExpiresIn != original.ExpiresIn {
			t.Errorf("ExpiresIn mismatch: got %v, want %v", copied.ExpiresIn, original.ExpiresIn)
		}
	})

	t.Run("WithExtraFields", func(t *testing.T) {
		extraMap := map[string]interface{}{
			"scope":        "read write",
			"custom_field": "custom_value",
			"number":       float64(123),
		}

		original := (&oauth2.Token{
			AccessToken:  "test_access_token",
			TokenType:    "Bearer",
			RefreshToken: "test_refresh_token",
		}).WithExtra(extraMap)

		copied := copyToken(original)

		// Verify extra fields are NOT copied (limitation of simple copy)
		if copied.Extra("scope") != nil {
			t.Errorf("Extra 'scope' should be nil (not copied), got %v", copied.Extra("scope"))
		}
		if copied.Extra("custom_field") != nil {
			t.Errorf("Extra 'custom_field' should be nil (not copied), got %v", copied.Extra("custom_field"))
		}
		if copied.Extra("number") != nil {
			t.Errorf("Extra 'number' should be nil (not copied), got %v", copied.Extra("number"))
		}
	})

	t.Run("IndependentCopy", func(t *testing.T) {
		original := &oauth2.Token{
			AccessToken:  "original_access",
			TokenType:    "Bearer",
			RefreshToken: "original_refresh",
		}

		copied := copyToken(original)

		// Modify the original
		original.AccessToken = "modified_access"
		original.RefreshToken = "modified_refresh"

		// Verify the copy is not affected
		if copied.AccessToken != "original_access" {
			t.Errorf("Copy was affected by original modification: got %v, want %v", copied.AccessToken, "original_access")
		}
		if copied.RefreshToken != "original_refresh" {
			t.Errorf("Copy was affected by original modification: got %v, want %v", copied.RefreshToken, "original_refresh")
		}
	})
}
