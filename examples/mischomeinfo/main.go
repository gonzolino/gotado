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

	// Get weather
	weather, err := home.GetWeather(ctx)
	fmt.Println("Weather:")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get weather: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Solar Intensity: %.2f\n", weather.SolarIntensity.Percentage)
	fmt.Printf("Outside Temperature: %.2f °C\n", weather.OutsideTemperature.Celsius)
	fmt.Printf("Weather State: %s\n", weather.WeatherState.Value)

	// Get Devices
	devices, err := home.GetDevices(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get devices: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("Type: %s, Firmware: %s\n", device.DeviceType, device.CurrentFwVersion)
	}

	// Get Installations
	installations, err := home.GetInstallations(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get installations: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Installations:")
	for _, installation := range installations {
		fmt.Printf("Type: %s, State: %s, Devices: %d\n", installation.Type, installation.State, len(installation.Devices))
	}

	// Get mobile Devices
	mobileDevices, err := home.GetMobileDevices(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get mobile devices: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Mobile Devices:")
	for _, mobileDevice := range mobileDevices {
		fmt.Printf("Name: %s, Device: %s, OS: %s (%s)\n",
			mobileDevice.Name,
			mobileDevice.DeviceMetadata.Model,
			mobileDevice.DeviceMetadata.Platform,
			mobileDevice.DeviceMetadata.OSVersion)
	}

	// Get Users
	users, err := home.GetUsers(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get users: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)
	}
}
