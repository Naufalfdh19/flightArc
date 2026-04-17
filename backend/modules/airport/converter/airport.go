package converter

import (
	"flight/modules/airport/dto"
	"flight/modules/airport/entity"
)

type GetAirportConverter struct{}

func (c GetAirportConverter) ToDto(schedule entity.Airport) dto.Airport {
	return dto.Airport{
		Id:            schedule.Id,
		Name: schedule.Name,
		Code: schedule.Code,
		City: schedule.City,
		Country: schedule.Country,
	}
}
