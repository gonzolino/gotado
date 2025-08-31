package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gonzolino/gotado/v2"
)

const (
	clientID = "1bb50063-6b0c-4d11-bd99-387f4a91cc46"
)

func main() {
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

	// Get user info and print some details
	user, err := tado.Me(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Email: %s\nUsername: %s\nName: %s\n", user.Email, user.Username, user.Name)

	// for each home: get home info and print name and address
	for _, userHome := range user.Homes {
		home, err := user.GetHome(ctx, userHome.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get user home info: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Home: %s\nAddress:\n%s\n%s %s\n", home.Name, *home.Address.AddressLine1, *home.Address.ZipCode, *home.Address.City)
	}
}
