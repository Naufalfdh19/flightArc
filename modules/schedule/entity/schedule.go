package entity

import (
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
