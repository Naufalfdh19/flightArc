package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Seat struct {
	Id uuid.UUID
	ScheduleId uuid.UUID
	Class string
	Price decimal.Decimal
	SeatNumber string
	Status string
}