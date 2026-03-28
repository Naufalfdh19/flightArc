package dto

import (
	scheduleDto "flight/modules/schedule/dto"
	userDto "flight/modules/user/dto"
	"time"

	"github.com/google/uuid"
)

type GetBooking struct {
	Id          uuid.UUID                `json:"id"`
	User        userDto.GetUserResponse  `json:"user"`
	Flight      scheduleDto.GetFlightDto `json:"flight"`
	Status      string                   `json:"status"`
	BookingTime time.Time                `json:"booking_time"`
}
