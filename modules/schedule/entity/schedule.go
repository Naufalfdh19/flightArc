package entity

import (
	"flight/modules/airport/entity"
	"time"
)

type Schedule struct {
	Id            string
	Origin        string
	Destination   string
	Status        string
	DepartureTime time.Time
	ArrivalTime   time.Time
}

type Flight struct {
	Id              string
	OriginCode      string         `gorm:"column:origin_code"`
	DestinationCode string         `gorm:"column:destination_code"`
	Origin          entity.Airport `gorm:"foreignKey:OriginCode;references:Code"`
	Destination     entity.Airport `gorm:"foreignKey:DestinationCode;references:Code"`
	Status          string
	DepartureTime   time.Time
	ArrivalTime     time.Time
}

func (Flight) TableName() string {
	return "schedules"
}
