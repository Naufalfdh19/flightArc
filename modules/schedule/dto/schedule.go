package dto

import "time"

type GetScheduleDto struct {
	Id            string    `json:"id"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	Status        string    `json:"status"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}
