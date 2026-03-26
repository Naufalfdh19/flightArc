package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Booking struct {
	Id            int
	UserId        int
	FlightId      int
	PassangerName int
	SeatNumber    string
	Status        string
	BookingTime   time.Time
}

type Ticket struct {
	Id            uuid.UUID       `json:"id"`
	BookingId     uuid.UUID       `json:"booking_id"`
	IdNumber      string          `json:"id_number"`
	PassengerName string          `json:"passenger_name"`
	SeatNumber    string          `json:"seat_number"`
	Class         string          `json:"class"`
	Price         decimal.Decimal `json:"price"`
}
