package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s homeName zoneName\n", os.Args[0])
		os.Exit(1)
	}
	homeName, zoneName := os.Args[1], os.Args[2]

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

	// Find zone to control
	zone, err := home.GetZone(ctx, zoneName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find zone '%s': %v\n", zoneName, err)
		os.Exit(1)
	}

	// Check if dazzle mode is currently enabled for zone
	dazzleMode, err := zone.GetDazzleMode(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check if dazzle mode is enabled: %v\n", err)
		os.Exit(1)
	}
	if dazzleMode.Enabled {
		fmt.Println("Dazzle mode is enabled")
	} else {
		fmt.Println("Dazzle mode is disabled")
	}

	// Toggle dazzle mode
	if dazzleMode.Enabled {
		err = zone.DisableDazzleMode(ctx)
	} else {
		err = zone.EnableDazzleMode(ctx)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to toggle dazzle mode: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Toggled dazzle mode")
	time.Sleep(10 * time.Second)

	// Toggle dazzle mode back to original value
	if dazzleMode.Enabled {
		err = zone.EnableDazzleMode(ctx)
	} else {
		err = zone.DisableDazzleMode(ctx)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to return to initial dazzle mode settings: %v\n", err)
		os.Exit(1)
	}
}
