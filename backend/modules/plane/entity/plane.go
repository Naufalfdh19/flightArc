package entity

type Plane struct {
	Id               int
	Name             string
	Seats            interface{}
	Capacity         int
	RegistrationCode string
	Status           string
	AirlineId        int
}
