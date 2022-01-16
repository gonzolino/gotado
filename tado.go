package gotado

import (
	"errors"
	"fmt"
	"strconv"
)

// GetMe returns information about the authenticated user.
func GetMe(client *Client) (*User, error) {
	me := &User{}
	if err := client.get(apiURL("me"), me); err != nil {
		return nil, err
	}
	return me, nil
}

// GetHome returns information about the given home
func GetHome(client *Client, userHome *UserHome) (*Home, error) {
	home := &Home{}
	homeID := strconv.Itoa(int(userHome.ID))
	if err := client.get(apiURL("homes/%s", homeID), home); err != nil {
		return nil, err
	}
	return home, nil
}

// GetHomeState returns the state of the given home
func GetHomeState(client *Client, userHome *UserHome) (*HomeState, error) {
	homeState := &HomeState{}
	if err := client.get(apiURL("homes/%d/state", userHome.ID), homeState); err != nil {
		return nil, err
	}
	return homeState, nil
}

// GetZones returns information about the zones in the given home
func GetZones(client *Client, userHome *UserHome) ([]*Zone, error) {
	zones := make([]*Zone, 0)
	if err := client.get(apiURL("homes/%d/zones", userHome.ID), &zones); err != nil {
		return nil, err
	}
	return zones, nil
}

// GetZoneState returns the state of the given zone
func GetZoneState(client *Client, userHome *UserHome, zone *Zone) (*ZoneState, error) {
	zoneState := &ZoneState{}
	if err := client.get(apiURL("homes/%d/zones/%d/state", userHome.ID, zone.ID), zoneState); err != nil {
		return nil, err
	}
	return zoneState, nil
}

// GetZoneCapabilities returns the capabilities of the given zone
func GetZoneCapabilities(client *Client, userHome *UserHome, zone *Zone) (*ZoneCapabilities, error) {
	zoneCapabilities := &ZoneCapabilities{}
	if err := client.get(apiURL("homes/%d/zones/%d/capabilities", userHome.ID, zone.ID), zoneCapabilities); err != nil {
		return nil, err
	}
	return zoneCapabilities, nil
}

// setZoneOverlay sets a zone overlay setting
func setZoneOverlay(client *Client, userHome *UserHome, zone *Zone, overlay *ZoneOverlay) error {
	return client.put(apiURL("homes/%d/zones/%d/overlay", userHome.ID, zone.ID), overlay)
}

// SetZoneOverlayHeatingOff turns off heating in a zone
func SetZoneOverlayHeatingOff(client *Client, userHome *UserHome, zone *Zone) error {
	overlay := &ZoneOverlay{
		Setting: &ZoneSetting{
			Type:  "HEATING",
			Power: "OFF",
		},
	}
	if err := setZoneOverlay(client, userHome, zone, overlay); err != nil {
		return err
	}

	if overlay.Type != "MANUAL" || overlay.Setting.Power != "OFF" {
		return errors.New("tado° did not accept new overlay")
	}

	return nil
}

// SetZoneOverlayHeatingOn turns on heating in a zone. The temperature should
// use the unit configured for the home. Returns the resulting overlay if successful.
func SetZoneOverlayHeatingOn(client *Client, userHome *UserHome, zone *Zone, temperature float64) (*ZoneOverlay, error) {
	home, err := GetHome(client, userHome)
	if err != nil || home == nil {
		return nil, fmt.Errorf("unable to determine temperature unit")
	}
	temperatureSetting := &ZoneSettingTemperature{}
	switch home.TemperatureUnit {
	case "CELSIUS":
		temperatureSetting.Celsius = temperature
	case "FAHRENHEIT":
		temperatureSetting.Fahrenheit = temperature
	default:
		return nil, fmt.Errorf("invalid temperature unit '%s'", home.TemperatureUnit)
	}

	overlay := &ZoneOverlay{
		Setting: &ZoneSetting{
			Type:        "HEATING",
			Power:       "ON",
			Temperature: temperatureSetting,
		},
	}
	if err := setZoneOverlay(client, userHome, zone, overlay); err != nil {
		return nil, err
	}

	if overlay.Type != "MANUAL" || overlay.Setting.Power != "ON" {
		return overlay, errors.New("tado° did not accept new overlay")
	}

	return overlay, nil
}

// DeleteZoneOverlay removes an overlay from a zone, thereby returning a zone to smart schedule
func DeleteZoneOverlay(client *Client, userHome *UserHome, zone *Zone) error {
	return client.delete(apiURL("homes/%d/zones/%d/overlay", userHome.ID, zone.ID))
}

// SetWindowOpen marks the window in a zone as open (open window must have been detected before)
func SetWindowOpen(client *Client, userHome *UserHome, zone *Zone) error {
	return client.post(apiURL("homes/%d/zones/%d/state/openWindow/activate", userHome.ID, zone.ID))
}

// SetWindowClosed marks the window in a zone as closed
func SetWindowClosed(client *Client, userHome *UserHome, zone *Zone) error {
	return client.delete(apiURL("homes/%d/zones/%d/state/openWindow", userHome.ID, zone.ID))
}

// GetTimetables lists available schedule timetables for the given zone
func GetTimetables(client *Client, userHome *UserHome, zone *Zone) ([]*ScheduleTimetable, error) {
	timetables := make([]*ScheduleTimetable, 0)
	if err := client.get(apiURL("homes/%d/zones/%d/schedule/timetables", userHome.ID, zone.ID), &timetables); err != nil {
		return nil, err
	}
	return timetables, nil
}

// GetActiveTimetable returns the active schedule timetable for the given zone
func GetActiveTimetable(client *Client, userHome *UserHome, zone *Zone) (*ScheduleTimetable, error) {
	timetable := &ScheduleTimetable{}
	if err := client.get(apiURL("homes/%d/zones/%d/schedule/activeTimetable", userHome.ID, zone.ID), timetable); err != nil {
		return nil, err
	}
	return timetable, nil
}

// SetActiveTimetable sets the active schedule timetable for the given zone
func SetActiveTimetable(client *Client, userHome *UserHome, zone *Zone, timetable *ScheduleTimetable) error {
	newTimetable := &ScheduleTimetable{ID: timetable.ID}

	return client.put(apiURL("homes/%d/zones/%d/schedule/activeTimetable", userHome.ID, zone.ID), newTimetable)
}

// GetSchedule returns the heating schedule of the given zone
func GetSchedule(client *Client, userHome *UserHome, zone *Zone, timetable *ScheduleTimetable) ([]*ScheduleBlock, error) {
	scheduleBlocks := make([]*ScheduleBlock, 0)
	if err := client.get(apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks", userHome.ID, zone.ID, timetable.ID), &scheduleBlocks); err != nil {
		return nil, err
	}
	return scheduleBlocks, nil
}

// SetSchedule sets the heating schedule for one block of the given zone
func SetSchedule(client *Client, userHome *UserHome, zone *Zone, timetable *ScheduleTimetable, schedule []*ScheduleBlock) error {
	// Order schedule blocks by day types.
	// For each daytipe we want to send one put request.
	scheduleMap := map[DayType][]*ScheduleBlock{}
	for _, scheduleBlock := range schedule {
		if _, ok := scheduleMap[scheduleBlock.DayType]; !ok {
			scheduleMap[scheduleBlock.DayType] = make([]*ScheduleBlock, 0, 1)
		}
		scheduleMap[scheduleBlock.DayType] = append(scheduleMap[scheduleBlock.DayType], scheduleBlock)
	}

	for dayType, scheduleBlocks := range scheduleMap {
		if err := client.put(apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks/%s", userHome.ID, zone.ID, timetable.ID, dayType), scheduleBlocks); err != nil {
			return err
		}
	}

	return nil
}

// GetAwayConfiguration returns the away configuration of the given zone
func GetAwayConfiguration(client *Client, userHome *UserHome, zone *Zone) (*AwayConfiguration, error) {
	awayConfig := &AwayConfiguration{}
	if err := client.get(apiURL("homes/%d/zones/%d/schedule/awayConfiguration", userHome.ID, zone.ID), awayConfig); err != nil {
		return nil, err
	}
	return awayConfig, nil
}

// SetAwayConfiguration sets an away configuration for the given zone
func SetAwayConfiguration(client *Client, userHome *UserHome, zone *Zone, awayConfig *AwayConfiguration) error {
	return client.put(apiURL("homes/%d/zones/%d/schedule/awayConfiguration", userHome.ID, zone.ID), *awayConfig)
}

// SetAwayTemperature sets the manual temperature for a zone when everybody leaves the house
func SetAwayTemperature(client *Client, userHome *UserHome, zone *Zone, temperature float64) error {
	home, err := GetHome(client, userHome)
	if err != nil || home == nil {
		return fmt.Errorf("unable to determine temperature unit")
	}
	temperatureSetting := &ZoneSettingTemperature{}
	switch home.TemperatureUnit {
	case "CELSIUS":
		temperatureSetting.Celsius = temperature
	case "FAHRENHEIT":
		temperatureSetting.Fahrenheit = temperature
	default:
		return fmt.Errorf("invalid temperature unit '%s'", home.TemperatureUnit)
	}

	awayConfig := &AwayConfiguration{
		Type:       "HEATING",
		AutoAdjust: false,
		Setting: &ZoneSetting{
			Type:        "HEATING",
			Power:       "ON",
			Temperature: temperatureSetting,
		},
	}

	return SetAwayConfiguration(client, userHome, zone, awayConfig)
}

// SetAwayComfortLevel sets the away configuration to auto-adjust at the given comfort level ("preheat").
// Allowed values got the comfort level are 0, 50 and 100 (Eco, Balanced, Comfort)
func SetAwayComfortLevel(client *Client, userHome *UserHome, zone *Zone, comfortLevel int32) error {
	awayConfig := &AwayConfiguration{
		Type:         "HEATING",
		AutoAdjust:   true,
		ComfortLevel: ComfortLevel(comfortLevel),
	}
	return SetAwayConfiguration(client, userHome, zone, awayConfig)
}

// setPresenceLock sets a locked presence on the given home (HOME or AWAY)
func setPresenceLock(client *Client, userHome *UserHome, presence PresenceLock) error {
	return client.put(apiURL("homes/%d/presenceLock", userHome.ID), presence)
}

// SetPresenceHome sets the geofencing presence to 'at home'
func SetPresenceHome(client *Client, userHome *UserHome) error {
	presence := PresenceLock{
		HomePresence: PresenceHome,
	}
	return setPresenceLock(client, userHome, presence)
}

// SetPresenceAway sets the geofencing presence to 'away'
func SetPresenceAway(client *Client, userHome *UserHome) error {
	presence := PresenceLock{
		HomePresence: PresenceAway,
	}
	return setPresenceLock(client, userHome, presence)
}

// SetPresenceAuto removes a locked geofencing presence and returns to auto mode
func SetPresenceAuto(client *Client, userHome *UserHome) error {
	return client.delete(apiURL("homes/%d/presenceLock", userHome.ID))
}

// IsEarlyStartEnabled returns if the given zone has turned on early start
func IsEarlyStartEnabled(client *Client, userHome *UserHome, zone *Zone) (bool, error) {
	earlyStart := &EarlyStart{}
	if err := client.get(apiURL("homes/%d/zones/%d/earlyStart", userHome.ID, zone.ID), earlyStart); err != nil {
		return false, err
	}
	return earlyStart.Enabled, nil
}

// setEarlyStart sets the early start setting for the given zone
func setEarlyStart(client *Client, userHome *UserHome, zone *Zone, earlyStart *EarlyStart) error {
	return client.put(apiURL("homes/%d/zones/%d/earlyStart", userHome.ID, zone.ID), earlyStart)
}

// EnableEarlyStart enables early start in the given zone
func EnableEarlyStart(client *Client, userHome *UserHome, zone *Zone) error {
	return setEarlyStart(client, userHome, zone, &EarlyStart{Enabled: true})
}

// DisableEarlyStart disables early start in the given zone
func DisableEarlyStart(client *Client, userHome *UserHome, zone *Zone) error {
	return setEarlyStart(client, userHome, zone, &EarlyStart{Enabled: false})
}

// GetWeather returns weather information at the given homes location
func GetWeather(client *Client, userHome *UserHome) (*Weather, error) {
	weather := &Weather{}
	if err := client.get(apiURL("homes/%d/weather", userHome.ID), weather); err != nil {
		return nil, err
	}
	return weather, nil
}

// GetDevices lists all devices in the given home
func GetDevices(client *Client, userHome *UserHome) ([]*Device, error) {
	devices := make([]*Device, 0)
	if err := client.get(apiURL("homes/%d/devices", userHome.ID), &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// GetZoneDevices lists all devices in the given home and zone
func GetZoneDevices(client *Client, userHome *UserHome, zone *Zone) ([]*Device, error) {
	devices := make([]*Device, 0)
	if err := client.get(apiURL("homes/%d/zones/%d/devices", userHome.ID, zone.ID), &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// GetInstallations lists all installations in the given home
func GetInstallations(client *Client, userHome *UserHome) ([]*Installation, error) {
	installations := make([]*Installation, 0)
	if err := client.get(apiURL("homes/%d/installations", userHome.ID), &installations); err != nil {
		return nil, err
	}
	return installations, nil
}

// GetMobileDevices lists all mobile devices linked to the given home
func GetMobileDevices(client *Client, userHome *UserHome) ([]*MobileDevice, error) {
	mobileDevices := make([]*MobileDevice, 0)
	if err := client.get(apiURL("homes/%d/mobileDevices", userHome.ID), &mobileDevices); err != nil {
		return nil, err
	}
	return mobileDevices, nil
}

// DeleteMobileDevice deletes the given mobile device
func DeleteMobileDevice(client *Client, userHome *UserHome, mobileDevice *MobileDevice) error {
	return client.delete(apiURL("homes/%d/mobileDevices/%d", userHome.ID, mobileDevice.ID))
}

// SetMobileDeviceSettings updates the given mobile device with the given settings
func SetMobileDeviceSettings(client *Client, userHome *UserHome, mobileDevice *MobileDevice, settings *MobileDeviceSettings) error {
	return client.put(apiURL("homes/%d/mobileDevices/%d/settings", userHome.ID, mobileDevice.ID), settings)
}

// GetUsers lists all users and their mobile devices linked to the given home
func GetUsers(client *Client, userHome *UserHome) ([]*User, error) {
	users := make([]*User, 0)
	if err := client.get(apiURL("homes/%d/users", userHome.ID), &users); err != nil {
		return nil, err
	}
	return users, nil
}
