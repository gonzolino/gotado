package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gonzolino/gotado/v2"
)

const (
	clientID = "1bb50063-6b0c-4d11-bd99-387f4a91cc46"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s homeName\n", os.Args[0])
		os.Exit(1)
	}
	homeName := os.Args[1]

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

	// List all devices
	devices, err := home.GetDevices(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list devices: %v\n", err)
		os.Exit(1)
	}
	for _, device := range devices {
		fmt.Printf("Device %s (type %s, connected %t)\n", device.SerialNo, device.DeviceType, device.ConnectionState.Value)
	}

	// Get current temperature offset from a device
	initialOffset, err := devices[1].GetTemperatureOffset(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get temperature offset: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Initial temperature offset: %.2f °C\n", initialOffset)

	// Set temperature offset on a device
	if err := devices[1].SetTemperatureOffset(ctx, 2.5); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set temperature offset: %v\n", err)
		os.Exit(1)
	}
	updatedOffset, err := devices[1].GetTemperatureOffset(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get temperature offset: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Updated temperature offset to %.2f °C\n", updatedOffset)

	// Restore to original offset
	if err := devices[1].SetTemperatureOffset(ctx, initialOffset); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set temperature offset: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restied temperature offset of %.2f °C\n", initialOffset)
}
