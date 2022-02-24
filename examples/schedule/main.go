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
	fmt.Println("Available heating schedule timetables:")
	for _, timetable := range []*gotado.ScheduleTimetable{zone.ScheduleMonToSun(), zone.ScheduleMonToFriSatSun(), zone.ScheduleAllDays()} {
		fmt.Printf("%s (%d)\n", timetable.Type, timetable.ID)
	}
	activeSchedule, err := zone.GetActiveScheduleTimetable(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Active heating schedule timetable: %s (%d)\n", activeSchedule.Type, activeSchedule.ID)

	// Get and print schedule time blocks
	timeBlocks, err := activeSchedule.GetTimeBlocks(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Heating schedule:")
	for _, block := range timeBlocks {
		fmt.Printf("%s (%s - %s): %s %s", block.DayType, block.Start, block.End, block.Setting.Type, block.Setting.Power)
		if block.Setting.Power == "ON" {
			fmt.Printf(" (%.2f°C, %.2f°F)", block.Setting.Temperature.Celsius, block.Setting.Temperature.Fahrenheit)
		}
		fmt.Println()
	}

	// Update schedule
	newSchedule := zone.ScheduleMonToSun()
	newTimeBlocks := []*gotado.ScheduleTimeBlock{
		{
			DayType:             "MONDAY_TO_SUNDAY",
			Start:               "00:00",
			End:                 "12:00",
			GeolocationOverride: false,
			Setting: &gotado.ZoneSetting{
				Type:  "HEATING",
				Power: "ON",
				Temperature: &gotado.ZoneSettingTemperature{
					Celsius: 18,
				},
			},
		},
		{
			DayType:             "MONDAY_TO_SUNDAY",
			Start:               "12:00",
			End:                 "14:00",
			GeolocationOverride: false,
			Setting: &gotado.ZoneSetting{
				Type:  "HEATING",
				Power: "OFF",
			},
		},
		{
			DayType:             "MONDAY_TO_SUNDAY",
			Start:               "14:00",
			End:                 "00:00",
			GeolocationOverride: false,
			Setting: &gotado.ZoneSetting{
				Type:  "HEATING",
				Power: "ON",
				Temperature: &gotado.ZoneSettingTemperature{
					Celsius: 20,
				},
			},
		},
	}
	if err := zone.SetActiveScheduleTimetable(ctx, newSchedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	if err := newSchedule.SetTimeBlocks(ctx, newTimeBlocks); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set heating schedule: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Changed heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Restore original heating schedule
	if err := zone.SetActiveScheduleTimetable(ctx, activeSchedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	if err := activeSchedule.SetTimeBlocks(ctx, timeBlocks); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restored original heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
}
