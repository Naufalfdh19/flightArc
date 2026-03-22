package dto

type Airport struct {
	Id      int    `json:"id,omitempty"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}
