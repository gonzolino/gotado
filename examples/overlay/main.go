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

	// Set heating off in zone
	if err := zone.SetHeatingOff(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn off heating: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Turned off heating in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(30 * time.Second)

	// Set heating on in zone (unit for temperature matches default unit of home)
	if err := zone.SetHeatingOn(ctx, 25.0); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn heating to 25 degrees: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Turned on heating in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(30 * time.Second)

	// Turn off manual heating control in zone. Return to smart schedule
	if err := zone.ResumeSchedule(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to turn off manual heating: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Turned heating in home '%s', zone '%s' back to smart schedule\n", home.Name, zone.Name)
}
