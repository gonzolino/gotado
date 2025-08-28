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
