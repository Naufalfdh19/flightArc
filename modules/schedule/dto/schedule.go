package dto

import (
	"flight/modules/airport/dto"
	"time"

	"github.com/shopspring/decimal"
)

type GetScheduleDto struct {
	Id            string    `json:"id"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	Status        string    `json:"status"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}

type GetFlightDto struct {
	Id            string `json:"id"`
	Origin        dto.Airport
	Destination   dto.Airport
	Status        string          `json:"status"`
	DepartureTime time.Time       `json:"departure_time"`
	ArrivalTime   time.Time       `json:"arrival_time"`
	Price         decimal.Decimal `json:"price"`
}

