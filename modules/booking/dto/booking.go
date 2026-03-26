package dto

import (
	scheduleDto "flight/modules/schedule/dto"
	userDto "flight/modules/user/dto"
	"time"
)

type GetBookingsDto struct {
	Id            int                      `json:"id"`
	User          userDto.GetUserResponse  `json:"user"`
	Flight        scheduleDto.GetFlightDto `json:"flight"`
	Status        string                   `json:"status"`
	BookingTime   time.Time                `json:"booking_time"`
}

