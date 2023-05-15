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

	// Get current termination condition for manual control
	terminationCondition, err := zone.GetManualControlTerminationCondition(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get termination condition for manual control: %v\n", err)
		os.Exit(1)
	}

	// Set condition to manual control
	if err := zone.ManualControlUntilUserEnd(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set manual control termination to manual: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Manual controlled until ended by user in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Set condition to timer
	if err := zone.ManualControlTimer(ctx, 900); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set manual control termination to timer: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Manual controlled until timer runs out in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Set condition to automatic control
	if err := zone.ManualControlUntilAutoChange(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set manual control termination to automatic: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Manual controlled until ended by next automatic change in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Return to original manual control termination condition
	switch terminationCondition.Type {
	case gotado.OverlayTypeManual:
		err = zone.ManualControlUntilUserEnd(ctx)
	case gotado.OverlayTypeTimer:
		err = zone.ManualControlTimer(ctx, terminationCondition.DurationInSeconds)
	case gotado.OverlayTypeAuto:
		err = zone.ManualControlUntilAutoChange(ctx)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to restore original manual control termination condition: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Manual control termination condition home '%s', zone '%s' restored\n", home.Name, zone.Name)
}
