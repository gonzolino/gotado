package gotado

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// User represents a tado° user
type User struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Username string     `json:"username"`
	ID       string     `json:"id"`
	Homes    []UserHome `json:"homes"`
	Locale   string     `json:"locale"`
	// TODO: MobileDevices missing
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

// ZoneState represents the state of a tado° zone
type ZoneState struct {
	TadoMode            string `json:"tadoMode"`
	GeolocationOverride bool   `json:"geolocationOverride"`
	// TODO missing geolocationOverrideDisableTime
	// TODO missing preparation
	Setting ZoneStateSetting `json:"setting"`
	// TODO missing overlayType
	// TODO missing overlay
	OpenWindow *ZoneStateOpenWindow `json:"openWindow"`
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

// ZoneStateSensorDataPointsInsideTemperature holds information about the inside temperatue of a zone
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

// GetMe returns information about the authenticated user.
func GetMe(client *Client) (*User, error) {
	resp, err := client.Request(http.MethodGet, apiURL("me"), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return nil, fmt.Errorf("tado° API error: %w", err)
	}

	me := &User{}
	if err := json.NewDecoder(resp.Body).Decode(me); err != nil {
		return nil, fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return me, nil
}

// GetHome returns information about the given home
func GetHome(client *Client, userHome *UserHome) (*Home, error) {
	homeID := strconv.Itoa(int(userHome.ID))
	resp, err := client.Request(http.MethodGet, apiURL("homes/%s", homeID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return nil, fmt.Errorf("tado° API error: %w", err)
	}

	home := &Home{}
	if err := json.NewDecoder(resp.Body).Decode(home); err != nil {
		return nil, fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return home, nil
}

// GetZones returns information about the zones in the given home
func GetZones(client *Client, userHome *UserHome) ([]*Zone, error) {
	resp, err := client.Request(http.MethodGet, apiURL("homes/%d/zones", userHome.ID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return nil, fmt.Errorf("tado° API error: %w", err)
	}

	zones := make([]*Zone, 0)
	if err := json.NewDecoder(resp.Body).Decode(&zones); err != nil {
		return nil, fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return zones, nil
}

// GetZoneState returns the state of the given zone
func GetZoneState(client *Client, userHome *UserHome, zone *Zone) (*ZoneState, error) {
	resp, err := client.Request(http.MethodGet, apiURL("homes/%d/zones/%d/state", userHome.ID, zone.ID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return nil, fmt.Errorf("tado° API error: %w", err)
	}

	zoneState := &ZoneState{}
	if err := json.NewDecoder(resp.Body).Decode(&zoneState); err != nil {
		return nil, fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return zoneState, nil
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
	resp, err := client.Request(http.MethodDelete, apiURL("homes/%d/zones/%d/state/openWindow", userHome.ID, zone.ID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected tado° API response status: %s", resp.Status)
	}
	return nil
}
