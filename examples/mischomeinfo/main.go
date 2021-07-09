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

	// Create authenticated tado° client
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

	// Find the home to control
	var home *gotado.UserHome
	for _, h := range user.Homes {
		if h.Name == homeName {
			home = &h
			break
		}
	}
	if home == nil {
		fmt.Fprintf(os.Stderr, "Home '%s' not found\n", homeName)
		os.Exit(1)
	}

	// Get weather
	weather, err := gotado.GetWeather(client, home)
	fmt.Println("Weather:")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get weather: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Solar Intensity: %.2f\n", weather.SolarIntensity.Percentage)
	fmt.Printf("Outside Temperature: %.2f °C\n", weather.OutsideTemperature.Celsius)
	fmt.Printf("Weather State: %s\n", weather.WeatherState.Value)

	// Get Devices
	devices, err := gotado.GetDevices(client, home)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get devices: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("Type: %s, Firmware: %s\n", device.DeviceType, device.CurrentFwVersion)
	}

	// Get Installations
	installations, err := gotado.GetInstallations(client, home)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get installations: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Installations:")
	for _, installation := range installations {
		fmt.Printf("Type: %s, State: %s, Devices: %d\n", installation.Type, installation.State, len(installation.Devices))
	}

	// Get mobile Devices
	mobileDevices, err := gotado.GetMobileDevices(client, home)
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
}
