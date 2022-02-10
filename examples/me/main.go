package main

import (
	"context"
	"fmt"
	"os"

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

	ctx := context.Background()
	tado := gotado.New(clientID, clientSecret)

	// Get user info and print some details
	user, err := tado.Me(ctx, username, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Email: %s\nUsername: %s\nName: %s\n", user.Email, user.Username, user.Name)

	// for each home: get home info and print name and address
	// for _, userHome := range user.Homes {
	// 	home, err := gotado.GetHome(client, &userHome)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Failed to get user home info: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("Home: %s\nAddress:\n%s\n%s %s\n", home.Name, *home.Address.AddressLine1, *home.Address.ZipCode, *home.Address.City)
	// }
}
