package dto

import "github.com/google/uuid"

type Airport struct {
	Id      uuid.UUID    `json:"id,omitempty"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}
