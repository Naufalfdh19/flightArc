package converter

import (
	airportConverter "flight/modules/airport/converter"
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

type GetFlightConverter struct{}

func (c GetFlightConverter) ToDto(schedule entity.Flight) dto.GetFlightDto {
	return dto.GetFlightDto{
		Id:            schedule.Id,
		Origin:        airportConverter.GetAirportConverter{}.ToDto(schedule.Origin),
		Destination:   airportConverter.GetAirportConverter{}.ToDto(schedule.Destination),
		Status:        schedule.Status,
		DepartureTime: schedule.DepartureTime,
		ArrivalTime:   schedule.ArrivalTime,
	}
}
