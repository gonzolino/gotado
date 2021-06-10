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

	// Show away configuration
	awayConfig, err := gotado.GetAwayConfiguration(client, home, zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get away configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Away Configuration:")
	if awayConfig.AutoAdjust {
		fmt.Printf("Comfort Level: %d\n", awayConfig.ComfortLevel)
	} else {
		fmt.Printf("Temperature: %.2f C°, %.2f F°\n", awayConfig.Setting.Temperature.Celsius, awayConfig.Setting.Temperature.Fahrenheit)
	}

	// Update comfort level
	err = gotado.SetAwayComfortLevel(client, home, zone, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get set comfort level: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Set comfort level for away mode in home '%s', zone '%s' to 'Eco'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Restore original away configuration
	if err := gotado.SetAwayConfiguration(client, home, zone, awayConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set away configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restored original away configuration in home '%s', zone '%s'\n", home.Name, zone.Name)
}
