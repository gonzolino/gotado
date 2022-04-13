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

	// Show away configuration
	awayConfig, err := zone.GetAwayConfiguration(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get away configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Away Configuration:")
	if awayConfig.AutoAdjust {
		fmt.Printf("Comfort Level: %d\n", awayConfig.ComfortLevel)
		fmt.Printf("Temperature: %.2f C°, %.2f F°\n", awayConfig.Setting.Temperature.Celsius, awayConfig.Setting.Temperature.Fahrenheit)
	}

	// Update comfort level
	err = zone.SetAwayPreheatComfortLevel(ctx, gotado.ComfortLevelEco)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get set comfort level: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Set comfort level for away mode in home '%s', zone '%s' to 'Eco'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Restore original away configuration
	if err := zone.SetAwayConfiguration(ctx, awayConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set away configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restored original away configuration in home '%s', zone '%s'\n", home.Name, zone.Name)
}
