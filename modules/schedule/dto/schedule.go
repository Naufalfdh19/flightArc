package dto

import "time"

type GetScheduleDto struct {
	Id            int       `json:"id"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	Status        string    `json:"status"`
	DepartureDate time.Time `json:"departure_date"`
}
