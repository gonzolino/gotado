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

	// Check if open window detection is currently enabled for zone
	openWindowDetection, err := zone.GetOpenWindowDetection(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check if open window detection is enabled: %v\n", err)
		os.Exit(1)
	}
	if openWindowDetection.Enabled {
		fmt.Printf("Open window detection is enabled, heating timeout: %ds\n", openWindowDetection.TimeoutInSeconds)
	} else {
		fmt.Println("Open window detection is disabled")
	}

	// Toggle open window detection
	if openWindowDetection.Enabled {
		err = zone.DisableOpenWindowDetection(ctx)
	} else {
		err = zone.EnableOpenWindowDetection(ctx, openWindowDetection.TimeoutInSeconds)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to toggle open window detection: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Toggled open window detection")
	time.Sleep(10 * time.Second)

	// Toggle open window detection back to original value
	if openWindowDetection.Enabled {
		err = zone.EnableOpenWindowDetection(ctx, openWindowDetection.TimeoutInSeconds)
	} else {
		err = zone.DisableOpenWindowDetection(ctx)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to return to initial open window detection settings: %v\n", err)
		os.Exit(1)
	}
}
