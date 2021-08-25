package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gonzolino/gotado/api"
	"github.com/gonzolino/gotado/api/client/home"
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

	ctx := context.Background()

	// Create authenticated tado° client
	client := api.NewAPI(ctx, clientID, clientSecret)
	if err := client.WithAuthentication(ctx, username, password); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to authenticate: %v\n", err)
		os.Exit(1)
	}

	// Get user info and print some details
	resp, err := client.Home.ShowUser(home.NewShowUserParams(), *client.BearerToken)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
		os.Exit(1)
	}
	user := resp.Payload
	fmt.Printf("Email: %s\nUsername: %s\nName: %s\n", *user.Email, *user.Username, *user.Name)

	// for each home: get home info and print name and address
	for _, userHome := range user.Homes {
		resp, err := client.Home.ShowHome(home.NewShowHomeParams().WithHomeID(int64(userHome.ID)), *client.BearerToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get user home info: %v\n", err)
			os.Exit(1)
		}
		home := resp.Payload
		fmt.Printf("Home: %s\nLocation:\n%f %f\n", *home.Name, *home.Geolocation.Latitude, *home.Geolocation.Longitude)
	}
}
