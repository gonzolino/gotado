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

	ctx := context.Background()

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
	fmt.Printf("Email: %s\nUsername: %s\nName: %s\n", user.Email, user.Username, user.Name)
	for _, userHome := range user.Homes {
		fmt.Printf("Home ID: %d\nHome Name: %s\n", userHome.ID, userHome.Name)
		home, err := gotado.GetHome(client, &userHome)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get user home info: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Address:\n%s\n%s %s\n", *home.Address.AddressLine1, *home.Address.ZipCode, *home.Address.City)

		zones, err := gotado.GetZones(client, &userHome)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get user home zones: %v\n", err)
			os.Exit(1)
		}
		for _, zone := range zones {
			fmt.Printf("Zone ID: %d\nZone Name: %s\n", zone.ID, zone.Name)
			zoneState, err := gotado.GetZoneState(client, &userHome, zone)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get zone state: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Zone State Mode: %s\n", zoneState.TadoMode)
			if zoneState.OpenWindow != nil {
				fmt.Printf("Open window detected at %s", zoneState.OpenWindow.DetectedTime)
			}
			if zone.Name == "Bad" {
				if err := gotado.SetWindowOpen(client, &userHome, zone); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to close window: %v", err)
				} else {
					fmt.Printf("Opened window in %s", zone.Name)
				}
			}
		}
	}
}
