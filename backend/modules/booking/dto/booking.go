package dto

import (
	scheduleDto "flight/modules/schedule/dto"
	userDto "flight/modules/user/dto"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type GetBookingReq struct {
	Id          uuid.UUID                `json:"id"`
	User        userDto.GetUserResponse  `json:"user"`
	Flight      scheduleDto.GetFlightDto `json:"flight"`
	Status      string                   `json:"status"`
	BookingTime time.Time                `json:"booking_time"`
}

type AddBookingReq struct {
	UserId     int
	ScheduleId uuid.UUID   `json:"schedule_id" binding:"required"`
	Tickets    []TicketDTO `json:"tickets" binding:"required,min=1"`
}

type TicketDTO struct {
	PassangerName string          `json:"passanger_name" binding:"required"`
	IdNumber      string          `json:"id_number" binding:"required"`
	SeatNumber    string          `json:"seat_number" binding:"required"`
	Class         string          `json:"class" binding:"required"`
	Price         decimal.Decimal `json:"price" binding:"required"`
}
