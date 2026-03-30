package dto

type GetUserResponse struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        string `json:"role,omitempty"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AddUserRequest struct {
	Name     string `json:"name" binding:"required,min=7,max=15"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Role     string `json:"role"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type ContactDetail struct {
	Email string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number,omitempty"`
}