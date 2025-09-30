package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Schedule struct {
	Id            int
	Origin        string
	Destination   string
	Status        string
	DepartureDate time.Time
	Price         decimal.Decimal
}
