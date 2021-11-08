package gotado

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// User represents a tado° user
type User struct {
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	ID            string         `json:"id"`
	Homes         []UserHome     `json:"homes"`
	Locale        string         `json:"locale"`
	MobileDevices []MobileDevice `json:"mobileDevices"`
}

// UserHome represents a home in a user object
type UserHome struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// Home represents a home equipped with tado°
type Home struct {
	ID              int32  `json:"id"`
	Name            string `json:"name"`
	DateTimeZone    string `json:"dateTimeZone"`
	DateCreated     string `json:"dateCreated"`
	TemperatureUnit string `json:"temperatureUnit"`
	// TODO partner missing
	SimpleSmartScheduleEnabled bool                  `json:"simpleSmartScheduleEnabled"`
	AwayRadiusInmeters         float32               `json:"awayRadiusInMeters"`
	InstallationCompleted      bool                  `json:"installationCompleted"`
	IncidentDetection          HomeIncidentDetection `json:"incidentDetection"`
	AutoAssistFreeTrialEnabled bool                  `json:"autoAssistFreeTrialEnabled"`
	Skills                     []string              `json:"skills"`
	ChristmasModeEnabled       bool                  `json:"christmasModeEnabled"`
	ShowAutoAssistReminders    bool                  `json:"showAutoAssistReminders"`
	ContactDetails             HomeContactDetails    `json:"contactDetails"`
	Address                    HomeAddress           `json:"address"`
	Geolocation                HomeGeolocation       `json:"geolocation"`
	ConsentGrantSkippable      bool                  `json:"consentGrantSkippable"`
}

// HomeIncidentDetection holds incident detection options for a home
type HomeIncidentDetection struct {
	Supported bool `json:"supported"`
	Enabled   bool `json:"enabled"`
}

// HomeContactDetails holds the contact details for a home
type HomeContactDetails struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

// HomeAddress holds the address of a home
type HomeAddress struct {
	AddressLine1 *string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	ZipCode      *string `json:"zipCode"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
}

// HomeGeolocation holds the coordinates of a home
type HomeGeolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// HomeState represents the state of a tado° home
type HomeState struct {
	Presence       string `json:"presence"`
	PresenceLocked bool   `json:"presenceLocked"`
}

// Zone represents a tado° zone
type Zone struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	DateCreated string   `json:"dateCreated"`
	DeviceTypes []string `json:"deviceTypes"`
	// TODO devices missing
	ReportAvailable     bool                    `json:"reportAvailable"`
	SupportsDazzle      bool                    `json:"supportsDazzle"`
	DazzleEnabled       bool                    `json:"dazzleEnabled"`
	DazzleMode          ZoneDazzleMode          `json:"dazzleMode"`
	OpenWindowDetection ZoneOpenWindowDetection `json:"openWindowDetection"`
}

// ZoneDazzleMode holds information about dazzle mode in a zone
type ZoneDazzleMode struct {
	Supported bool `json:"supported"`
	Enabled   bool `json:"enabled"`
}

// ZoneOpenWindowDetection holds information about open window detection in a zone
type ZoneOpenWindowDetection struct {
	Supported        bool  `json:"supported"`
	Enabled          bool  `json:"enabled"`
	TimeoutInSeconds int32 `json:"timeoutInSeconds"`
}

// ZoneCapabilitiesstores the capabilities of a zone, such as the supported
// min/max temperatures
type ZoneCapabilities struct {
	Type              string                        `json:"type"`
	CanSetTemperature *bool                         `json:"canSetTemperature,omitempty"`
	Temperatures      *ZoneCapabilitiesTemperatures `json:"temperatures,omitempty"`
}

// ZoneCapabilitiesTemperatures holds the temperature related capabilities of a zone
type ZoneCapabilitiesTemperatures struct {
	Celsius    *ZoneCapabilitiesTemperatureValues `json:"celsius,omitempty"`
	Fahrenheit *ZoneCapabilitiesTemperatureValues `json:"fahrenheit,omitempty"`
}

// ZoneCapabilitiesTemperatureValues holds the numeric values of temperature
// related capabilities of a zone
type ZoneCapabilitiesTemperatureValues struct {
	Min  int32   `json:"min"`
	Max  int32   `json:"max"`
	Step float32 `json:"step"`
}

// ZoneState represents the state of a tado° zone
type ZoneState struct {
	TadoMode            string `json:"tadoMode"`
	GeolocationOverride bool   `json:"geolocationOverride"`
	// TODO missing geolocationOverrideDisableTime
	// TODO missing preparation
	Setting     ZoneStateSetting     `json:"setting"`
	OverlayType *string              `json:"overlayType"`
	Overlay     *ZoneOverlay         `json:"overlay"`
	OpenWindow  *ZoneStateOpenWindow `json:"openWindow"`
	// TODO missing nextScheduleChange
	// TODO missing nextTimeBlock
	Link ZoneStateLink `json:"link"`
	// TODO missing activityDataPoints
	SensorDataPoints *ZoneStateSensorDataPoints `json:"sensorDataPoints"`
}

// ZoneStateSetting holds the setting of a zone
type ZoneStateSetting struct {
	Type        string                      `json:"type"`
	Power       string                      `json:"power"`
	Temperature ZoneStateSettingTemperature `json:"temperature"`
}

// ZoneStateSettingTemperature holds the temperature of a zone state setting
type ZoneStateSettingTemperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// ZoneOverlay holds overlay information of a zone
type ZoneOverlay struct {
	Type        string                  `json:"type,omitempty"`
	Setting     ZoneOverlaySetting      `json:"setting"`
	Termination *ZoneOverlayTermination `json:"termination,omitempty"`
}

// ZoneOverlaySetting holds the setting of a zone overlay
type ZoneOverlaySetting struct {
	Type        string                         `json:"type"`
	Power       string                         `json:"power"`
	Temperature *ZoneOverlaySettingTemperature `json:"temperature,omitempty"`
}

// ZoneOverlaySettingTemperature holds the temperature of a zone state setting
type ZoneOverlaySettingTemperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// ZoneOverlayTermination holdes the termination information of a zone overlay
type ZoneOverlayTermination struct {
	Type                   string  `json:"type"`
	TypeSkillBasedApp      string  `json:"typeSkillBasedApp"`
	DurationInSeconds      int32   `json:"durationInSeconds,omitempty"`
	Expiry                 string  `json:"expiry,omitempty"`
	RemainingTimeInSeconds int32   `json:"remainingTimeInSeconds,omitempty"`
	ProjectedExpiry        *string `json:"projectedExpiry"`
}

// ZoneStateOpenWindow holds the information about an open window of a zone state
type ZoneStateOpenWindow struct {
	DetectedTime           string `json:"detectedTime"`
	DurationInSeconds      int32  `json:"durationInSeconds"`
	Expiry                 string `json:"expiry"`
	RemainingTimeInSeconds int32  `json:"remainingTimeInSeconds"`
}

// ZoneStateLink holds the link information of a tado zone
type ZoneStateLink struct {
	State string `json:"state"`
}

// ZoneStateActivityDataPoints holds activity data points for a zone
type ZoneStateActivityDataPoints struct {
	HeatingPower *ZoneStateActivityDataPointsHeatingPower `json:"heatingPower"`
}

// ZoneStateActivityDataPointsHeatingPower holds information about the heating power in a zone
type ZoneStateActivityDataPointsHeatingPower struct {
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Timestamp  string  `json:"timestamp"`
}

// ZoneStateSensorDataPoints holds sensor data points for a zone
type ZoneStateSensorDataPoints struct {
	InsideTemperature *ZoneStateSensorDataPointsInsideTemperature `json:"insideTemperature"`
	Humidity          *ZoneStateSensorDataPointsHumidity          `json:"humidity"`
}

// ZoneStateSensorDataPointsInsideTemperature holds information about the inside temperature of a zone
type ZoneStateSensorDataPointsInsideTemperature struct {
	Celsius    float64                                             `json:"celsius"`
	Fahrenheit float64                                             `json:"fahrenheit"`
	Timestamp  string                                              `json:"timestamp"`
	Type       string                                              `json:"type"`
	Precision  ZoneStateSensorDataPointsInsideTemperaturePrecision `json:"precision"`
}

// ZoneStateSensorDataPointsInsideTemperaturePrecision holds the precision of inside temperature of a zone
type ZoneStateSensorDataPointsInsideTemperaturePrecision struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// ZoneStateSensorDataPointsHumidity holds humidity information of a zone
type ZoneStateSensorDataPointsHumidity struct {
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Timestamp  string  `json:"timestamp"`
}

// ScheduleTimetable is the type of a tado° schedule timetable
type ScheduleTimetable struct {
	ID   int32  `json:"id"`
	Type string `json:"type,omitempty"`
}

// ScheduleBlock is a block in a tado° schedule
type ScheduleBlock struct {
	DayType             string               `json:"dayType"`
	Start               string               `json:"start"`
	End                 string               `json:"end"`
	GeolocationOverride bool                 `json:"geolocationOverride"`
	Setting             ScheduleBlockSetting `json:"setting"`
}

// ScheduleBlockSetting holds the setting of a schedule block
type ScheduleBlockSetting struct {
	Type        string                           `json:"type"`
	Power       string                           `json:"power"`
	Temperature *ScheduleBlockSettingTemperature `json:"temperature,omitempty"`
}

// ZoneOverlaySettingTemperature holds the temperature of a schedule block setting
type ScheduleBlockSettingTemperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// AwayConfiguration holds the settings to use when everybody leaves the house
type AwayConfiguration struct {
	Type       string `json:"type"`
	AutoAdjust bool   `json:"autoAdjust"`
	// Comfort Level must be 0 (Eco), 50 (Balanced) or 100 (Comfort)
	ComfortLevel int32                     `json:"comfortLevel"`
	Setting      *AwayConfigurationSetting `json:"setting"`
}

// AwayConfigurationSetting holds the setting of an away configuration
type AwayConfigurationSetting struct {
	Type        string                               `json:"type"`
	Power       string                               `json:"power"`
	Temperature *AwayConfigurationSettingTemperature `json:"temperature,omitempty"`
}

// AwayConfigurationSettingTemperature holds the temperature of an away configuration setting
type AwayConfigurationSettingTemperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// PresenceLock holds a locked presence setting for a home
type PresenceLock struct {
	HomePresence string `json:"homePresence"`
}

// EarlyStart controls whether tado° ensures that a set temperature is reached
// at the start of a block.
type EarlyStart struct {
	Enabled bool `json:"enabled"`
}

// Weather holds weather information from the home's location
type Weather struct {
	SolarIntensity     *WeatherSolarIntensity     `json:"solarIntensity"`
	OutsideTemperature *WeatherOutsideTemperature `json:"outsideTemperature"`
	WeatherState       *WeatherState              `json:"weatherState"`
}

// WeatherSolarIntensity holds the solar intensity at the home's location as a percentage
type WeatherSolarIntensity struct {
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
	Timestamp  string  `json:"timestamp"`
}

// WeatherOutsideTemperature holds the temperature outside of the home
type WeatherOutsideTemperature struct {
	Celsius    float64                            `json:"celsius"`
	Fahrenheit float64                            `json:"fahrenheit"`
	Timestamp  string                             `json:"timestamp"`
	Type       string                             `json:"type"`
	Precision  WeatherOutsideTemperaturePrecision `json:"precision"`
}

// WeatherOutsideTemperaturePrecision holds the precision of the home's outside temperature
type WeatherOutsideTemperaturePrecision struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// WeatherState stores the state of the weather, e.g. rain, sunny, foggy...
type WeatherState struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

// Device represents a tado° device such as a thermostat or a bridge
type Device struct {
	DeviceType       string                `json:"deviceType"`
	SerialNo         string                `json:"serialNo"`
	ShortSerialNo    string                `json:"shortSerialNo"`
	CurrentFwVersion string                `json:"currentFwVersion"`
	ConnectionState  DeviceConnectionState `json:"connectionState"`
	Characteristics  DeviceCharacteristics `json:"characteristics"`
	InPairingMode    *bool                 `json:"inPairingMode,omitempty"`
	MountingState    *DeviceMountingState  `json:"mountingState,omitempty"`
	BatteryState     *string               `json:"batteryState,omitempty"`
	ChildLockEnabled *bool                 `json:"childLockEnabled,omitempty"`
	GatewayOperation *string               `json:"gatewayOperation,omitempty"`
}

// DeviceConnectionState specifies if the device is connected or not
type DeviceConnectionState struct {
	Value     bool   `json:"value"`
	Timestamp string `json:"timestamp"`
}

// DeviceCharacteristics lists the capabilities of a device
type DeviceCharacteristics struct {
	Capabilities []string `json:"characteristics"`
}

// DeviceMountingState holds the mounting state of a device, e.g. if it is calibrated
type DeviceMountingState struct {
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

// Installation holds information about a tado° hardware installation
type Installation struct {
	ID       int32    `json:"id"`
	Type     string   `json:"type"`
	Revision int32    `json:"revision"`
	State    string   `json:"state"`
	Devices  []Device `json:"devices"`
}

// MobileDevice represents a mobile device with the tado° app installed
type MobileDevice struct {
	Name           string                `json:"name"`
	ID             int32                 `json:"id"`
	Settings       MobileDeviceSettings  `json:"settings"`
	Location       *MobileDeviceLocation `json:"location"`
	DeviceMetadata MobileDeviceMetadata  `json:"deviceMetadata"`
}

// MobileDeviceSettings holds the settings of a mobile device
type MobileDeviceSettings struct {
	GeoTrackingEnabled bool                                   `json:"geoTrackingEnabled"`
	PushNotifications  *MobileDeviceSettingsPushNotifications `json:"pushNotifications,omitempty"`
}

// MobileDeviceSettingsPushNotifications holds the push notification settings
type MobileDeviceSettingsPushNotifications struct {
	LowBatteryReminder          bool `json:"lowBatteryReminder"`
	AwayModeReminder            bool `json:"awayModeReminder"`
	HomeModeReminder            bool `json:"homeModeReminder"`
	OpenWindowReminder          bool `json:"openWindowReminder"`
	EnergySavingsReportReminder bool `json:"energySavingsReportReminder"`
	IncidentDetection           bool `json:"incidentDetection"`
}

// MobileDeviceLocation holds information regarding the current location of  mobile device
type MobileDeviceLocation struct {
	Stale                         bool                                `json:"stale"`
	AtHome                        bool                                `json:"atHome"`
	BearingFromHome               MobileDeviceLocationBearingFromHome `json:"bearingFromHome"`
	RelativeDistanceFromHomeFence float64                             `json:"relativeDistanceFromHomeFence"`
}

// MobileDeviceLocationBearingFromHome holds the current bearing of a mobile device from the home
type MobileDeviceLocationBearingFromHome struct {
	Degrees float64 `json:"degrees"`
	Radians float64 `json:"radians"`
}

// MobileDeviceMetadata holds some general metadata about a mobile device
type MobileDeviceMetadata struct {
	Platform  string `json:"platform"`
	OSVersion string `json:"osVersion"`
	Model     string `json:"model"`
	Locale    string `json:"locale"`
}

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
		Setting: ZoneOverlaySetting{
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
	temperatureSetting := &ZoneOverlaySettingTemperature{}
	switch home.TemperatureUnit {
	case "CELSIUS":
		temperatureSetting.Celsius = temperature
	case "FAHRENHEIT":
		temperatureSetting.Fahrenheit = temperature
	default:
		return nil, fmt.Errorf("invalid temperature unit '%s'", home.TemperatureUnit)
	}

	overlay := &ZoneOverlay{
		Setting: ZoneOverlaySetting{
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
	resp, err := client.Request(http.MethodPost, apiURL("homes/%d/zones/%d/state/openWindow/activate", userHome.ID, zone.ID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected tado° API response status: %s", resp.Status)
	}
	return nil
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

	if err := client.put(apiURL("homes/%d/zones/%d/schedule/activeTimetable", userHome.ID, zone.ID), newTimetable); err != nil {
		return err
	}

	return nil
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
	scheduleMap := map[string][]*ScheduleBlock{}
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
	temperatureSetting := &AwayConfigurationSettingTemperature{}
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
		Setting: &AwayConfigurationSetting{
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
		ComfortLevel: comfortLevel,
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
		HomePresence: "HOME",
	}
	return setPresenceLock(client, userHome, presence)
}

// SetPresenceAway sets the geofencing presence to 'away'
func SetPresenceAway(client *Client, userHome *UserHome) error {
	presence := PresenceLock{
		HomePresence: "AWAY",
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
