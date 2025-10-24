package gotado

import (
	"testing"
	"time"

	"golang.org/x/oauth2"
)

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
