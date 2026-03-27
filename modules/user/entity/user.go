package entity

import "gorm.io/gorm"

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
