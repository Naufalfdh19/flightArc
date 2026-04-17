package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Ticket struct {
	Id            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingId     uuid.UUID
	IdNumber      string
	PassengerName string
	SeatNumber    string
	Class         string
	Price         decimal.Decimal
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
