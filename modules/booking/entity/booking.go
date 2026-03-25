package entity

import "time"

type Booking struct {
	Id            int
	UserId        int
	FlightId      int
	PassangerName int
	SeatNumber    string
	Status        string
	BookingTime   time.Time
}