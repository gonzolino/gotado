package gotado

import (
	"context"
	"errors"
	"fmt"
)

// GetState returns the state of the zone.
func (z *Zone) GetState(ctx context.Context) (*ZoneState, error) {
	state := &ZoneState{}
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/state", z.home.ID, z.ID), state); err != nil {
		return nil, err
	}
	return state, nil
}

// GetCapabilities returns the capabilities of the zone.
func (z *Zone) GetCapabilities(ctx context.Context) (*ZoneCapabilities, error) {
	capabilities := &ZoneCapabilities{}
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/capabilities", z.home.ID, z.ID), capabilities); err != nil {
		return nil, err
	}
	return capabilities, nil
}

// GetDevices lists all devices in the zone
func (z *Zone) GetDevices(ctx context.Context) ([]*Device, error) {
	devices := make([]*Device, 0)
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/devices", z.home.ID, z.ID), &devices); err != nil {
		return nil, err
	}
	for _, device := range devices {
		device.client = z.client
	}
	return devices, nil
}

// ResumeSchedule resumes the zone's smart schedule.
func (z *Zone) ResumeSchedule(ctx context.Context) error {
	return z.client.delete(ctx, apiURL("homes/%d/zones/%d/overlay", z.home.ID, z.ID))
}

// SetHeatingOff turns off the heating in the zone.
func (z *Zone) SetHeatingOff(ctx context.Context) error {
	overlay := &ZoneOverlay{
		Setting: &ZoneSetting{
			Type:  "HEATING",
			Power: "OFF",
		},
	}
	if err := z.client.put(ctx, apiURL("homes/%d/zones/%d/overlay", z.home.ID, z.ID), overlay); err != nil {
		return err
	}

	if overlay.Type != "MANUAL" || overlay.Setting.Power != "OFF" {
		return errors.New("tado° refused to turn off heating")
	}

	return nil
}

// SetHeatingOn turns on the heating in the zone. The temperature should
// use the unit configured for the home.
func (z *Zone) SetHeatingOn(ctx context.Context, temperature float64) error {
	temperatureSetting := &ZoneSettingTemperature{}
	switch z.home.TemperatureUnit {
	case TemperatureUnitCelsius:
		temperatureSetting.Celsius = temperature
	case TemperatureUnitFahrenheit:
		temperatureSetting.Fahrenheit = temperature
	default:
		return fmt.Errorf("invalid temperature unit '%s'", z.home.TemperatureUnit)
	}

	overlay := &ZoneOverlay{
		Setting: &ZoneSetting{
			Type:        ZoneTypeHeating,
			Power:       PowerOn,
			Temperature: temperatureSetting,
		},
	}
	if err := z.client.put(ctx, apiURL("homes/%d/zones/%d/overlay", z.home.ID, z.ID), overlay); err != nil {
		return err
	}

	if overlay.Type != OverlayTypeManual || overlay.Setting.Power != PowerOn {
		return errors.New("tado° refused to set the given temperature")
	}

	return nil
}

// OpenWindow puts the zone into open window mode (open window must have been
// detected by tado° beforehand).
func (z *Zone) OpenWindow(ctx context.Context) error {
	return z.client.post(ctx, apiURL("homes/%d/zones/%d/state/openWindow/activate", z.home.ID, z.ID))
}

// CloseWindow ends open window mode in the zone.
func (z *Zone) CloseWindow(ctx context.Context) error {
	return z.client.delete(ctx, apiURL("homes/%d/zones/%d/state/openWindow", z.home.ID, z.ID))
}

// GetEarlyStart checks if early start is enabled in the zone.
func (z *Zone) GetEarlyStart(ctx context.Context) (bool, error) {
	earlyStart := &EarlyStart{}
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/earlyStart", z.home.ID, z.ID), earlyStart); err != nil {
		return false, err
	}
	return earlyStart.Enabled, nil
}

// SetEarlyStart enables or disables early start in the zone.
func (z *Zone) SetEarlyStart(ctx context.Context, earlyStart bool) error {
	return z.client.put(ctx, apiURL("homes/%d/zones/%d/earlyStart", z.home.ID, z.ID), &EarlyStart{Enabled: earlyStart})
}

// newScheduleTimetable creates a new schedule timetable linked to the zone.
func (z *Zone) newScheduleTimetable(id int32, typ TimetableType) *ScheduleTimetable {
	return &ScheduleTimetable{
		client: z.client,
		zone:   z,
		ID:     id,
		Type:   typ,
	}
}

// newHeatingSchedule creates a new heating schedule linked to the zone.
func (z *Zone) newHeatingSchedule(timetable *ScheduleTimetable, blocks []*ScheduleTimeBlock) *HeatingSchedule {
	return &HeatingSchedule{
		zone:      z,
		Timetable: timetable,
		Blocks:    blocks,
	}
}

// ScheduleMonToSun has the same schedule for all days between monday and sunday.
func (z *Zone) ScheduleMonToSun(ctx context.Context) (*HeatingSchedule, error) {
	timetable := z.newScheduleTimetable(0, TimetableOneDay)
	blocks, err := timetable.GetTimeBlocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get schedule time blocks: %w", err)
	}

	return z.newHeatingSchedule(timetable, blocks), nil
}

// TimetableTMonToFriSatSun has the same schedule for all days between monday
// and friday and different schedules for saturday and sunday.
func (z *Zone) ScheduleMonToFriSatSun(ctx context.Context) (*HeatingSchedule, error) {
	timetable := z.newScheduleTimetable(1, TimetableThreeDay)
	blocks, err := timetable.GetTimeBlocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get schedule time blocks: %w", err)
	}

	return z.newHeatingSchedule(timetable, blocks), nil
}

// ScheduleAllDays has a different schedule for each day of the week.
func (z *Zone) ScheduleAllDays(ctx context.Context) (*HeatingSchedule, error) {
	timetable := z.newScheduleTimetable(2, TimetableSevenDay)
	blocks, err := timetable.GetTimeBlocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get schedule time blocks: %w", err)
	}

	return z.newHeatingSchedule(timetable, blocks), nil
}

// GetActiveScheduleTimetable returns the active schedule timetable for the zone.
func (z *Zone) GetActiveScheduleTimetable(ctx context.Context) (*ScheduleTimetable, error) {
	timetable := &ScheduleTimetable{}
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/schedule/activeTimetable", z.home.ID, z.ID), timetable); err != nil {
		return nil, err
	}
	timetable.client = z.client
	timetable.zone = z
	return timetable, nil
}

// SetActiveScheduleTimetable sets the active schedule timetable for the zone.
// Should be one of TimetableMonToSun(), TimetableMonToFriSatSun() or TimetableAllDays(),
func (z *Zone) SetActiveScheduleTimetable(ctx context.Context, timetable *ScheduleTimetable) error {
	newTimetable := &ScheduleTimetable{ID: timetable.ID}
	return z.client.put(ctx, apiURL("homes/%d/zones/%d/schedule/activeTimetable", z.home.ID, z.ID), newTimetable)
}

// GetHeatingSchedule gets the whole active schedule for the zone, including active timetable and time blocks.
func (z *Zone) GetHeatingSchedule(ctx context.Context) (*HeatingSchedule, error) {
	timetable, err := z.GetActiveScheduleTimetable(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get active schedule timetable: %w", err)
	}
	blocks, err := timetable.GetTimeBlocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get time blocks: %w", err)
	}

	var scheduleDays ScheduleDays
	switch timetable.Type {
	case TimetableOneDay:
		scheduleDays = ScheduleDaysMonToSun
	case TimetableThreeDay:
		scheduleDays = ScheduleDaysMonToFriSatSun
	case TimetableSevenDay:
		scheduleDays = ScheduleDaysMonTueWedThuFriSatSun
	default:
		return nil, errors.New("unknown schedule timetable type")
	}

	return &HeatingSchedule{
		zone:         z,
		ScheduleDays: scheduleDays,
		Timetable:    timetable,
		Blocks:       blocks,
	}, nil
}

// SetHeatingSchedule sets the whole active schedule for the zone, including active timetable and time blocks.
func (z *Zone) SetHeatingSchedule(ctx context.Context, schedule *HeatingSchedule) error {
	if err := z.SetActiveScheduleTimetable(ctx, schedule.Timetable); err != nil {
		return fmt.Errorf("unable to set active schedule timetable: %w", err)
	}
	if err := schedule.Timetable.SetTimeBlocks(ctx, schedule.Blocks); err != nil {
		return fmt.Errorf("unable to set time blocks: %w", err)
	}
	return nil
}

// GetAwayConfiguration returns the away configuration of the zone.
func (z *Zone) GetAwayConfiguration(ctx context.Context) (*AwayConfiguration, error) {
	awayConfig := &AwayConfiguration{}
	if err := z.client.get(ctx, apiURL("homes/%d/zones/%d/schedule/awayConfiguration", z.home.ID, z.ID), awayConfig); err != nil {
		return nil, err
	}
	return awayConfig, nil
}

// SetAwayConfiguration updates the away configuration of the zone.
func (z *Zone) SetAwayConfiguration(ctx context.Context, awayConfig *AwayConfiguration) error {
	return z.client.put(ctx, apiURL("homes/%d/zones/%d/schedule/awayConfiguration", z.home.ID, z.ID), *awayConfig)
}

// SetAwayMinimumTemperature sets the minimum temperature for away mode in the zone.
func (z *Zone) SetAwayMinimumTemperature(ctx context.Context, temperature float64) error {
	awayConfig, err := z.GetAwayConfiguration(ctx)
	if err != nil {
		return fmt.Errorf("unable to get current away configuration: %w", err)
	}

	switch z.home.TemperatureUnit {
	case TemperatureUnitCelsius:
		awayConfig.Setting.Temperature.Celsius = temperature
	case TemperatureUnitFahrenheit:
		awayConfig.Setting.Temperature.Fahrenheit = temperature
	default:
		return fmt.Errorf("invalid temperature unit '%s'", z.home.TemperatureUnit)
	}

	return z.SetAwayConfiguration(ctx, awayConfig)
}

// SetAwayPreheatOff turns off preheat before arrival. Tado° will only start
// heating after arrival. To turn preheating back on, use SetAwayPreheatComfortLevel().
func (z *Zone) SetAwayPreheatOff(ctx context.Context) error {
	awayConfig, err := z.GetAwayConfiguration(ctx)
	if err != nil {
		return fmt.Errorf("unable to get current away configuration: %w", err)
	}

	awayConfig.AutoAdjust = false

	return z.SetAwayConfiguration(ctx, awayConfig)
}

// SetAwayPreheatComfortLevel sets the comfort level for preheating before arrival.
func (z *Zone) SetAwayPreheatComfortLevel(ctx context.Context, comfortLevel ComfortLevel) error {
	awayConfig, err := z.GetAwayConfiguration(ctx)
	if err != nil {
		return fmt.Errorf("unable to get current away configuration: %w", err)
	}

	awayConfig.AutoAdjust = true
	awayConfig.ComfortLevel = comfortLevel

	return z.SetAwayConfiguration(ctx, awayConfig)
}
