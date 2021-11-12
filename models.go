package gotado

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
	ID                         int32                 `json:"id"`
	Name                       string                `json:"name"`
	DateTimeZone               string                `json:"dateTimeZone"`
	DateCreated                string                `json:"dateCreated"`
	TemperatureUnit            string                `json:"temperatureUnit"`
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
	ID                  int32                   `json:"id"`
	Name                string                  `json:"name"`
	Type                string                  `json:"type"`
	DateCreated         string                  `json:"dateCreated"`
	DeviceTypes         []string                `json:"deviceTypes"`
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
	TadoMode                       string                        `json:"tadoMode"`
	GeolocationOverride            bool                          `json:"geolocationOverride"`
	GeolocationOverrideDisableTime *string                       `json:"geolocationOverrideDisableTime"`
	Setting                        ZoneStateSetting              `json:"setting"`
	OverlayType                    *string                       `json:"overlayType"`
	Overlay                        *ZoneOverlay                  `json:"overlay"`
	OpenWindow                     *ZoneStateOpenWindow          `json:"openWindow"`
	NextScheduledChange            *ZoneStateNextScheduledChange `json:"nextScheduleChange"`
	NextTimeBlock                  *ZoneStateNextTimeBlock       `json:"nextTimeBlock"`
	Link                           ZoneStateLink                 `json:"link"`
	ActivityDataPoints             *ZoneStateActivityDataPoints  `json:"activityDataPoints"`
	SensorDataPoints               *ZoneStateSensorDataPoints    `json:"sensorDataPoints"`
}

// ZoneStateSetting holds the setting of a zone
type ZoneStateSetting struct {
	Type        string                       `json:"type"`
	Power       string                       `json:"power"`
	Temperature *ZoneStateSettingTemperature `json:"temperature"`
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

// ZoneStateNextScheduledChange holds start time and settings of the next scheduled change
type ZoneStateNextScheduledChange struct {
	Start   string            `json:"start"`
	Setting *ZoneStateSetting `json:"setting"`
}

// ZoneStateNextTimeBlock holds the start time of the next time block
type ZoneStateNextTimeBlock struct {
	Start string `json:"start"`
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
	Duties           []string              `json:"duties,omitempty"`
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
