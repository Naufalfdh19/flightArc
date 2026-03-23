package entity

import (
	"flight/modules/airport/entity"
	"time"

	"github.com/shopspring/decimal"
)

type Schedule struct {
	Id            string
	Origin        string
	Destination   string
	Status        string
	DepartureTime time.Time
	ArrivalTime   time.Time
	Price         decimal.Decimal
}

type Flight struct {
	Id            string
	Origin        entity.Airport
	Destination   entity.Airport
	Status        string
	DepartureTime time.Time
	ArrivalTime   time.Time
	Price         decimal.Decimal
}
