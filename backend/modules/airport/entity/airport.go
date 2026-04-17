package entity

import "github.com/google/uuid"

type Airport struct {
	Id      uuid.UUID
	Code    string
	Name    string
	City    string
	Country string
}
