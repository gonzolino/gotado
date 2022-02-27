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
