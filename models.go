package gotado

import "time"

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

// TemperatureUnit defines the unit in which a temperature is measured
type TemperatureUnit string

const (
	TemperatureUnitCelsius    TemperatureUnit = "CELSIUS"
	TemperatureUnitFahrenheit TemperatureUnit = "FAHRENHEIT"
)

// Home represents a home equipped with tado°
type Home struct {
	ID                         int32                 `json:"id"`
	Name                       string                `json:"name"`
	DateTimeZone               string                `json:"dateTimeZone"`
	DateCreated                time.Time             `json:"dateCreated"`
	TemperatureUnit            TemperatureUnit       `json:"temperatureUnit"`
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

// Presence defines if somebody is present at home
type Presence string

const (
	PresenceHome Presence = "HOME"
	PresenceAway Presence = "AWAY"
)

// HomeState represents the state of a tado° home
type HomeState struct {
	Presence       Presence `json:"presence"`
	PresenceLocked bool     `json:"presenceLocked"`
}

// ZoneType defines the type of a zone
type ZoneType string

const (
	ZoneTypeHeating  = "HEATING"
	ZoneTypeHotWater = "HOT_WATER"
)

// DeviceType defines the type of a device
type DeviceType string

const (
	DeviceTypeInternetBridge               DeviceType = "IB01"
	DeviceTypeExtensionKit01               DeviceType = "BU01"
	DeviceTypeExtensionKit02               DeviceType = "BU02"
	DeviceTypeSmartACControl01             DeviceType = "WR01"
	DeviceTypeSmartACControl02             DeviceType = "WR02"
	DeviceTypeSmartThermostat01            DeviceType = "RU01"
	DeviceTypeSmartThermostat02            DeviceType = "RU02"
	DeviceTypeSmartRadiatorThermostat01    DeviceType = "VA01"
	DeviceTypeSmartRadiatorThermostat02    DeviceType = "VA02"
	DeviceTypeWirelessTemperatureSensor01  DeviceType = "SU02"
	DeviceTypeWirelessReceiverProgrammer01 DeviceType = "BP02"
	DeviceTypeWirelessReceiverBoiler01     DeviceType = "BR02"
)

// Zone represents a tado° zone
type Zone struct {
	ID                  int32                   `json:"id"`
	Name                string                  `json:"name"`
	Type                ZoneType                `json:"type"`
	DateCreated         time.Time               `json:"dateCreated"`
	DeviceTypes         []DeviceType            `json:"deviceTypes"`
	Devices             []Device                `json:"devices"`
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

// ZoneCapabilities stores the capabilities of a zone, such as the supported
// min/max temperatures
type ZoneCapabilities struct {
	Type              ZoneType                      `json:"type"`
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

// OverlayType specifies the type of an overlay
type OverlayType string

const (
	OverlayTypeManual OverlayType = "MANUAL"
)

// ZoneState represents the state of a tado° zone
type ZoneState struct {
	TadoMode                       string                        `json:"tadoMode"`
	GeolocationOverride            bool                          `json:"geolocationOverride"`
	GeolocationOverrideDisableTime *string                       `json:"geolocationOverrideDisableTime"`
	Setting                        ZoneSetting                   `json:"setting"`
	OverlayType                    *OverlayType                  `json:"overlayType"`
	Overlay                        *ZoneOverlay                  `json:"overlay"`
	OpenWindow                     *ZoneStateOpenWindow          `json:"openWindow"`
	NextScheduledChange            *ZoneStateNextScheduledChange `json:"nextScheduleChange"`
	NextTimeBlock                  *ZoneStateNextTimeBlock       `json:"nextTimeBlock"`
	Link                           ZoneStateLink                 `json:"link"`
	ActivityDataPoints             *ZoneStateActivityDataPoints  `json:"activityDataPoints"`
	SensorDataPoints               *ZoneStateSensorDataPoints    `json:"sensorDataPoints"`
}

// Power specifies is something is powered on or off
type Power string

const (
	PowerOn  Power = "ON"
	PowerOff Power = "OFF"
)

// ZoneSetting holds the setting of a zone
type ZoneSetting struct {
	Type        ZoneType                `json:"type"`
	Power       Power                   `json:"power"`
	Temperature *ZoneSettingTemperature `json:"temperature"`
}

// ZoneSettingTemperature holds the temperature of a zone state setting
type ZoneSettingTemperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// ZoneOverlay holds overlay information of a zone
type ZoneOverlay struct {
	Type        OverlayType             `json:"type,omitempty"`
	Setting     *ZoneSetting            `json:"setting"`
	Termination *ZoneOverlayTermination `json:"termination,omitempty"`
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

// ZoneStateNextScheduledChange holds start time and settings of the next scheduled change
type ZoneStateNextScheduledChange struct {
	Start   time.Time    `json:"start"`
	Setting *ZoneSetting `json:"setting"`
}

// ZoneStateNextTimeBlock holds the start time of the next time block
type ZoneStateNextTimeBlock struct {
	Start time.Time `json:"start"`
}

// LinkState holds the state of a link (online or offline)
type LinkState string

const (
	LinkStateOnline  LinkState = "ONLINE"
	LinkStateOffline LinkState = "OFFLINE"
)

// ZoneStateLink holds the link information of a tado zone
type ZoneStateLink struct {
	State  string              `json:"state"`
	Reason ZoneStateLinkReason `json:"reason,omitempty"`
}

// ZoneStateLinkReason holds the reason why a link is offline
type ZoneStateLinkReason struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// ZoneStateActivityDataPoints holds activity data points for a zone
type ZoneStateActivityDataPoints struct {
	HeatingPower *PercentageMeasurement `json:"heatingPower"`
}

// ZoneStateSensorDataPoints holds sensor data points for a zone
type ZoneStateSensorDataPoints struct {
	InsideTemperature *TemperatureMeasurement `json:"insideTemperature"`
	Humidity          *PercentageMeasurement  `json:"humidity"`
}

// MeasurementType specifies teh type of a measurement
type MeasurementType string

const (
	MeasurementTypeTemperature MeasurementType = "TEMPERATURE"
	MeasurementTypePercentage  MeasurementType = "PERCENTAGE"
	MeasurementTypeWeather     MeasurementType = "WEATHER_STATE"
)

// Measurement measures a value at a certain point in time.
// See MeasurementType for available types of measurements.
type Measurement struct {
	Type      MeasurementType `json:"type"`
	Timestamp time.Time       `json:"timestamp"`
}

// TemperatureMeasurement holds a measured temperature
type TemperatureMeasurement struct {
	Measurement
	Celsius    float64                         `json:"celsius"`
	Fahrenheit float64                         `json:"fahrenheit"`
	Precision  TemperatureMeasurementPrecision `json:"precision"`
}

// TemperatureMeasurementPrecision holds the precision of a temperature measurement
type TemperatureMeasurementPrecision struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// PercentageMeasurement holds a measured percentage
type PercentageMeasurement struct {
	Measurement
	Percentage float64 `json:"percentage"`
}

// WeatherMeasurement holds a measurement of the weather state
type WeatherMeasurement struct {
	Measurement
	Value string `json:"value"`
}

type TimetableType string

const (
	TimetableTypeOneDay   TimetableType = "ONE_DAY"
	TimetableTypeThreeDay TimetableType = "THREE_DAY"
	TimetableTypeSevenDay TimetableType = "SEVEN_DAY"
)

// ScheduleTimetable is the type of a tado° schedule timetable
type ScheduleTimetable struct {
	ID   int32         `json:"id"`
	Type TimetableType `json:"type,omitempty"`
}

// DayType specifies the type of day for a heating schedule block
type DayType string

const (
	DayTypeMondayToSunday DayType = "MONDAY_TO_SUNDAY"
	DayTypeMondayToFriday DayType = "MONDAY_TO_FRIDAY"
	DayTypeMonday         DayType = "MONDAY"
	DayTypeTuesday        DayType = "TUESDAY"
	DayTypeWednesday      DayType = "WEDNESDAY"
	DayTypeThursday       DayType = "THURSDAY"
	DayTypeFriday         DayType = "FRIDAY"
	DayTypeSaturday       DayType = "SATURDAY"
	DayTypeSunday         DayType = "SUNDAY"
)

// ScheduleBlock is a block in a tado° schedule
type ScheduleBlock struct {
	DayType             DayType      `json:"dayType"`
	Start               string       `json:"start"`
	End                 string       `json:"end"`
	GeolocationOverride bool         `json:"geolocationOverride"`
	Setting             *ZoneSetting `json:"setting"`
}

// ComfortLevel defines how a zone is preheated before arrival
type ComfortLevel int32

const (
	// ComfortLevelEco will not preheat the zone too early before arrival and only reach the target temperature after arrival
	ComfortLevelEco = 0
	// ComfortLevelBalance will find the best trade-off between comfort and savings
	ComfortLevelBalance = 50
	// ComfortLevelComfort ensures that the desired home temperature is reached shortly before arrival
	ComfortLevelComfort = 100
)

// AwayConfiguration holds the settings to use when everybody leaves the house
type AwayConfiguration struct {
	Type         ZoneType     `json:"type"`
	AutoAdjust   bool         `json:"autoAdjust"`
	ComfortLevel ComfortLevel `json:"comfortLevel"`
	Setting      *ZoneSetting `json:"setting"`
}

// PresenceLock holds a locked presence setting for a home
type PresenceLock struct {
	HomePresence Presence `json:"homePresence"`
}

// EarlyStart controls whether tado° ensures that a set temperature is reached
// at the start of a block.
type EarlyStart struct {
	Enabled bool `json:"enabled"`
}

// Weather holds weather information from the home's location
type Weather struct {
	SolarIntensity     *PercentageMeasurement  `json:"solarIntensity"`
	OutsideTemperature *TemperatureMeasurement `json:"outsideTemperature"`
	WeatherState       *WeatherMeasurement     `json:"weatherState"`
}

// Device represents a tado° device such as a thermostat or a bridge
type Device struct {
	DeviceType       DeviceType            `json:"deviceType"`
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
	Duties           []string              `json:"duties,omitempty"`
}

// DeviceConnectionState specifies if the device is connected or not
type DeviceConnectionState struct {
	Value     bool      `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

// DeviceCharacteristics lists the capabilities of a device
type DeviceCharacteristics struct {
	Capabilities []string `json:"capabilities"`
}

// DeviceMountingState holds the mounting state of a device, e.g. if it is calibrated
type DeviceMountingState struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
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

// TemperatureOffset holds the current temperature offsets for a tado° device
type TemperatureOffset struct {
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
}
