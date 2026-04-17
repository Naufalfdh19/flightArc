package converter

import (
	"flight/modules/plane/dto"
	"flight/modules/plane/entity"
)

type AddPlaneConverter struct{}

func (c AddPlaneConverter) ToEntity(planeDto dto.AddPlaneRequest) entity.Plane {
	return entity.Plane{
		Name: planeDto.Name,
		Seats: planeDto.Seats,
		Capacity: planeDto.Capacity,
		RegistrationCode: planeDto.RegistrationCode,
		Status: planeDto.Status,
		AirlineId: planeDto.AirlineId,
	}
}

type UpdateSeatsConverter struct{}

func (c UpdateSeatsConverter) ToEntity(seatsDto dto.UpdateSeats) entity.Plane {
	return entity.Plane{
		Seats: seatsDto.Seats,
	}
}


