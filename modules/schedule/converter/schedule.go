package converter

import (
	"flight/modules/schedule/dto"
	"flight/modules/schedule/entity"
)

type GetScheduleConverter struct{}

func (c GetScheduleConverter) ToDto(schedule entity.Schedule) dto.GetScheduleDto {
	return dto.GetScheduleDto{
		Id:            schedule.Id,
		Origin:        schedule.Origin,
		Destination:   schedule.Destination,
		Status:        schedule.Status,
		DepartureTime: schedule.DepartureTime,
		ArrivalTime:   schedule.ArrivalTime,
	}
}
