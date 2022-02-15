package gotado

import "context"

// GetScheduleTimeBlocks returns all time blocks of the timetable schedule.
func (s *ScheduleTimetable) GetScheduleTimeBlocks(ctx context.Context) ([]*ScheduleTimeBlock, error) {
	blocks := make([]*ScheduleTimeBlock, 0)
	if err := s.client.get(ctx, apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks", s.zone.home.ID, s.zone.ID, s.ID), &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

func (s *ScheduleTimetable) SetScheduleTimeBlocks(ctx context.Context, schedule []*ScheduleTimeBlock) error {
	// Order schedule blocks by day types.
	// For each daytipe we want to send one put request.
	scheduleMap := map[DayType][]*ScheduleTimeBlock{}
	for _, scheduleBlock := range schedule {
		if _, ok := scheduleMap[scheduleBlock.DayType]; !ok {
			scheduleMap[scheduleBlock.DayType] = make([]*ScheduleTimeBlock, 0, 1)
		}
		scheduleMap[scheduleBlock.DayType] = append(scheduleMap[scheduleBlock.DayType], scheduleBlock)
	}

	for dayType, scheduleBlocks := range scheduleMap {
		if err := s.client.put(ctx, apiURL("homes/%d/zones/%d/schedule/timetables/%d/blocks/%s", s.zone.home.ID, s.zone.ID, s.ID, dayType), scheduleBlocks); err != nil {
			return err
		}
	}

	return nil
}
