package dto

type AddPlaneRequest struct {
	Name             string      `json:"name" binding:"required"`
	Seats            interface{} `json:"seats" binding:"required"`
	Capacity         int         `json:"capacity" binding:"required"`
	RegistrationCode string      `json:"registration_code" binding:"required"`
	Status           string      `json:"status" binding:"required"`
	AirlineId        int         `json:"airline_id" binding:"required"`
}


