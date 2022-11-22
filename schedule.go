package gotado

import (
	"context"
)

// GetTimeBlocks returns all time blocks of the schedule.
func (s *ScheduleTimetable) GetTimeBlocks(ctx context.Context) ([]*ScheduleTimeBlock, error) {
	blocks := make([]*ScheduleTimeBlock, 0)
	if err := s.client.get(ctx, apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks", s.zone.home.ID, s.zone.ID, s.ID), &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

// SetTimeBlocks updates the schedule with the given time blocks.
func (s *ScheduleTimetable) SetTimeBlocks(ctx context.Context, blocks []*ScheduleTimeBlock) error {
	// Order schedule blocks by day types.
	// For each daytipe we want to send one put request.
	scheduleMap := map[DayType][]*ScheduleTimeBlock{}
	for _, block := range blocks {
		if _, ok := scheduleMap[block.DayType]; !ok {
			scheduleMap[block.DayType] = make([]*ScheduleTimeBlock, 0, 1)
		}
		scheduleMap[block.DayType] = append(scheduleMap[block.DayType], block)
	}

	for dayType, scheduleBlocks := range scheduleMap {
		if err := s.client.put(ctx, apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks/%s", s.zone.home.ID, s.zone.ID, s.ID, dayType), scheduleBlocks); err != nil {
			return err
		}
	}

	return nil
}

// NewTimeBlock resets the list of time blocks in the heating schedule and adds the given time block as the first new block.
func (s *HeatingSchedule) NewTimeBlock(ctx context.Context, dayType DayType, start, end string, geolocationOverride bool, power Power, temperature float64) *HeatingSchedule {
	s.Blocks = make([]*ScheduleTimeBlock, 0)
	return s.AddTimeBlock(ctx, dayType, start, end, geolocationOverride, power, temperature)
}

// AddTimeBlock adds a time block to the heating schedule.
// Start and end parameters define when the time blocks starts and ends and are
// in the format HH:MM. GeolocationOverride specifies if the timeblock will
// override geofencing control. Power defines if heating is powered on or off
// and temperature specifies the temperature to heat to. Temperature is
// interpreted in Celsius / Fahrenheit depending on the temperature unit
// configured in the home.
func (s *HeatingSchedule) AddTimeBlock(_ context.Context, dayType DayType, start, end string, geolocationOverride bool, power Power, temperature float64) *HeatingSchedule {
	var temp *ZoneSettingTemperature = nil

	// Only set temperature if power is on.
	// If power is off, the temperature should be null
	if power == PowerOn {
		temp = &ZoneSettingTemperature{}
		switch s.zone.home.TemperatureUnit {
		case TemperatureUnitCelsius:
			temp.Celsius = temperature
		case TemperatureUnitFahrenheit:
			temp.Fahrenheit = temperature
		}
	}

	block := &ScheduleTimeBlock{
		DayType:             dayType,
		Start:               start,
		End:                 end,
		GeolocationOverride: geolocationOverride,
		Setting: &ZoneSetting{
			Type:        ZoneTypeHeating,
			Power:       power,
			Temperature: temp,
		},
	}
	s.Blocks = append(s.Blocks, block)
	return s
}
