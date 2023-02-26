package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gonzolino/gotado/v2"
)

const (
	clientID     = "tado-web-app"
	clientSecret = "wZaRN7rpjn3FoNyF5IFuxg9uMzYJcvOoQ8QWiIqS3hfk6gLhVlG57j5YNoZL2Rtc"
)

func main() {
	// Get credentials from env vars
	username, ok := os.LookupEnv("TADO_USERNAME")
	if !ok {
		fmt.Fprintf(os.Stderr, "Variable TADO_USERNAME not set\n")
		os.Exit(1)
	}
	password, ok := os.LookupEnv("TADO_PASSWORD")
	if !ok {
		fmt.Fprintf(os.Stderr, "Variable TADO_PASSWORD not set\n")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s homeName\n", os.Args[0])
		os.Exit(1)
	}
	homeName := os.Args[1]

	ctx := context.Background()
	tado := gotado.New(clientID, clientSecret)

	user, err := tado.Me(ctx, username, password)
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
