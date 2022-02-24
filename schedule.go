package gotado

import "context"

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
