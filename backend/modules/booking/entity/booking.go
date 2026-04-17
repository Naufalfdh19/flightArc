package entity

import (
	flightEnt "flight/modules/schedule/entity"
	userEnt "flight/modules/user/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserId      int
	ScheduleId  uuid.UUID
	User        userEnt.User     `gorm:"foreignKey:UserId"`
	Flight      flightEnt.Flight `gorm:"foreignKey:ScheduleId"`
	Status      string
	BookingTime time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
