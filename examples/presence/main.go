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

	// Get current presence from home state
	state, err := home.GetState(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get home state: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Initial Presence: %s", state.Presence)
	if state.PresenceLocked {
		fmt.Printf(" (locked)\n")
	} else {
		fmt.Println()
	}

	// Lock presence to 'away'
	if err := home.SetPresenceAway(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set presence 'away': %v", err)
		os.Exit(1)
	}

	fmt.Println("Set presence away")
	time.Sleep(10 * time.Second)

	// Lock presence to 'at home'
	if err := home.SetPresenceHome(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set presence 'home': %v", err)
		os.Exit(1)
	}

	fmt.Println("Set presence home")
	time.Sleep(10 * time.Second)

	// Set auto presence
	if err := home.SetPresenceAuto(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set presence 'auto': %v", err)
		os.Exit(1)
	}

	fmt.Println("Set presence auto")
	time.Sleep(10 * time.Second)

	// Return to initial presence settings
	if state.PresenceLocked {
		switch state.Presence {
		case "HOME":
			err = home.SetPresenceHome(ctx)
		case "AWAY":
			err = home.SetPresenceAway(ctx)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to return to initial presence settings: %v", err)
			os.Exit(1)
		}
	}
}
