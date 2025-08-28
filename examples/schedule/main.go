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
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s homeName zoneName\n", os.Args[0])
		os.Exit(1)
	}
	homeName, zoneName := os.Args[1], os.Args[2]

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

	// Find zone to control
	zone, err := home.GetZone(ctx, zoneName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find zone '%s': %v\n", zoneName, err)
		os.Exit(1)
	}

	// Show schedule timetables
	activeSchedule, err := zone.GetHeatingSchedule(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get active heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Active heating schedule days: %s\n", activeSchedule.ScheduleDays)

	// Get and print schedule time blocks
	fmt.Println("Heating schedule:")
	for _, block := range activeSchedule.Blocks {
		fmt.Printf("%s (%s - %s): %s %s", block.DayType, block.Start, block.End, block.Setting.Type, block.Setting.Power)
		if block.Setting.Power == "ON" {
			fmt.Printf(" (%.2f°C, %.2f°F)", block.Setting.Temperature.Celsius, block.Setting.Temperature.Fahrenheit)
		}
		fmt.Println()
	}

	// Update schedule
	newSchedule, err := zone.ScheduleMonToFriSatSun(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize new heating schedule: %v\n", err)
		os.Exit(1)
	}

	newSchedule = newSchedule.
		NewTimeBlock(ctx, gotado.DayTypeMondayToFriday, "00:00", "07:00", true, gotado.PowerOff, 0.0).
		AddTimeBlock(ctx, gotado.DayTypeMondayToFriday, "07:00", "00:00", false, gotado.PowerOn, 20.0).
		AddTimeBlock(ctx, gotado.DayTypeSaturday, "00:00", "09:00", true, gotado.PowerOff, 0.0).
		AddTimeBlock(ctx, gotado.DayTypeSaturday, "09:00", "00:00", false, gotado.PowerOn, 18.0).
		AddTimeBlock(ctx, gotado.DayTypeSunday, "00:00", "09:00", true, gotado.PowerOff, 0.0).
		AddTimeBlock(ctx, gotado.DayTypeSunday, "09:00", "00:00", false, gotado.PowerOn, 19.0)

	if err := zone.SetHeatingSchedule(ctx, newSchedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Changed heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Restore original heating schedule
	if err := zone.SetHeatingSchedule(ctx, activeSchedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restored original heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
}
