package main

import (
	"context"
	"fmt"
	"net/http"
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

	// Create authenticated tado° client
	httpClient := &http.Client{Timeout: 5 * time.Second}
	client := gotado.NewClient(clientID, clientSecret).WithHTTPClient(httpClient)
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

	// Find zone to control
	zones, err := gotado.GetZones(client, home)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get zones: %v\n", err)
		os.Exit(1)
	}
	var zone *gotado.Zone
	for _, z := range zones {
		if z.Name == zoneName {
			zone = z
			break
		}
	}
	if zone == nil {
		fmt.Fprintf(os.Stderr, "Zone '%s' not found\n", zoneName)
		os.Exit(1)
	}

	// Show schedule timetables
	timetables, err := gotado.GetTimetables(client, home, zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list available heating schedule timetables: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Available heating schedule timetables:")
	for _, timetable := range timetables {
		fmt.Printf("%s (%d)\n", timetable.Type, timetable.ID)
	}
	activeTimetable, err := gotado.GetActiveTimetable(client, home, zone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Active heating schedule timetable: %s (%d)\n", activeTimetable.Type, activeTimetable.ID)

	// Get and print schedule
	schedule, err := gotado.GetSchedule(client, home, zone, activeTimetable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Heating schedule:")
	for _, block := range schedule {
		fmt.Printf("%s (%s - %s): %s %s", block.DayType, block.Start, block.End, block.Setting.Type, block.Setting.Power)
		if block.Setting.Power == "ON" {
			fmt.Printf(" (%.2f°C, %.2f°F)", block.Setting.Temperature.Celsius, block.Setting.Temperature.Fahrenheit)
		}
		fmt.Println()
	}

	// Update schedule
	newTimetable := &gotado.ScheduleTimetable{
		ID:   0,
		Type: "ONE_DAY",
	}
	newSchedule := []*gotado.ScheduleBlock{
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
	if err := gotado.SetActiveTimetable(client, home, zone, newTimetable); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	if err := gotado.SetSchedule(client, home, zone, newTimetable, newSchedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set heating schedule: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Changed heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
	time.Sleep(10 * time.Second)

	// Restore original heating schedule
	if err := gotado.SetActiveTimetable(client, home, zone, activeTimetable); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set active heating schedule timetable: %v\n", err)
		os.Exit(1)
	}
	if err := gotado.SetSchedule(client, home, zone, activeTimetable, schedule); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set heating schedule: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Restored original heating schedule in home '%s', zone '%s'\n", home.Name, zone.Name)
}
