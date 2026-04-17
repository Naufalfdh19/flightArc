package entity

import "github.com/shopspring/decimal"

type Seat struct {
	SeatNumber  string          
	Class       string          
	Price       decimal.Decimal 
	IsAvailable bool
}