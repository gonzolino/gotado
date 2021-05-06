package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gonzolino/gotado"
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

	// Create authenticated tado° client
	client := gotado.NewClient(clientID, clientSecret).WithTimeout(5 * time.Second)
	client, err := client.WithCredentials(ctx, username, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Authentication failed: %v\n", err)
		os.Exit(1)
	}

	user, err := gotado.GetMe(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
		os.Exit(1)
	}

	// Find the home to control
	var home *gotado.UserHome
	for _, h := range user.Homes {
		if h.Name == homeName {
			home = &h
			break
		}
	}
	if home == nil {
		fmt.Fprintf(os.Stderr, "Home '%s' not found\n", homeName)
		os.Exit(1)
	}

	// Find zone to control
	zones, err := gotado.GetZones(client, home)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get zones: %v\n", err)
		os.Exit(1)
	}
	var zone *gotado.Zone
	for _, z := range zones {
		if z.Name == zoneName {
			zone = z
			break
		}
	}
	if zone == nil {
		fmt.Fprintf(os.Stderr, "Zone '%s' not found\n", zoneName)
		os.Exit(1)
	}

	// Set heating off in zone
	if err := gotado.SetZoneOverlayHeatingOff(client, home, zone); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn off heating: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Turned off heating in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(30 * time.Second)

	// Set heating on in zone (unit for temperature matches default unit of home)
	overlay, err := gotado.SetZoneOverlayHeatingOn(client, home, zone, 25)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn heating to 25 degrees: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Turned heating in home '%s', zone '%s' to %f°C\n", home.Name, zone.Name, overlay.Setting.Temperature.Celsius)
	time.Sleep(30 * time.Second)

	// Turn off manual heating control in zone. Return to smart schedule
	if err := gotado.DeleteZoneOverlay(client, home, zone); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn off manual heating: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Turned heating in home '%s', zone '%s' back to smart schedule\n", home.Name, zone.Name)
}
