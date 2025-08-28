package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gonzolino/gotado/v2"
)

const (
	clientID = "1bb50063-6b0c-4d11-bd99-387f4a91cc46"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s homeName zoneName\n", os.Args[0])
		os.Exit(1)
	}
	homeName, zoneName := os.Args[1], os.Args[2]

	ctx := context.Background()

	// Authenticate to tado
	// (see https://support.tado.com/en/articles/8565472-how-do-i-authenticate-to-access-the-rest-api)
	config := gotado.AuthConfig(clientID)
	response, err := config.DeviceAuth(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain device authorization: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("To authenticate, visit %s\n", response.VerificationURIComplete)

	token, err := config.DeviceAccessToken(ctx, response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain access token: %v\n", err)
		os.Exit(1)
	}

	// Create client
	tado := gotado.New(ctx, config, token)

	// Get user info
	user, err := tado.Me(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
		os.Exit(1)
	}

	// Find the home to control
	home, err := user.GetHome(ctx, homeName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find home '%s': %v\n", homeName, err)
		os.Exit(1)
	}

	// Find zone to control
	zone, err := home.GetZone(ctx, zoneName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find zone '%s': %v\n", zoneName, err)
		os.Exit(1)
	}

	// Check if early start is currently enabled for zone
	earlyStartEnabled, err := zone.GetEarlyStart(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check if early start is enabled: %v\n", err)
		os.Exit(1)
	}
	if earlyStartEnabled {
		fmt.Println("Early start is enabled")
	} else {
		fmt.Println("Early start is disabled")
	}

	// Toggle early start setting
	if earlyStartEnabled {
		err = zone.SetEarlyStart(ctx, false)
	} else {
		err = zone.SetEarlyStart(ctx, true)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to toggle early start: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Toggled early start")
	time.Sleep(10 * time.Second)

	// Toggle early start back to original value
	if earlyStartEnabled {
		err = zone.SetEarlyStart(ctx, true)
	} else {
		err = zone.SetEarlyStart(ctx, false)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to return to initial early start settings: %v\n", err)
		os.Exit(1)
	}
}
