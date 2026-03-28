package entity

import (
	flightEnt "flight/modules/schedule/entity"
	userEnt "flight/modules/user/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Booking struct {
	Id          uuid.UUID
	UserId      int
	ScheduleId    uuid.UUID
	User        userEnt.User    `gorm:"foreignKey:UserId"`
    Flight      flightEnt.Flight  `gorm:"foreignKey:ScheduleId"`
	Status      string
	BookingTime time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Ticket struct {
	Id            uuid.UUID
	BookingId     uuid.UUID
	IdNumber      string
	PassengerName string
	SeatNumber    string
	Class         string
	Price         decimal.Decimal
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
