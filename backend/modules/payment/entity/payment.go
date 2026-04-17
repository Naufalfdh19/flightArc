package entity

import (
	bookingEntity "flight/modules/booking/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	Id            uuid.UUID             `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingId     uuid.UUID             `gorm:"type:uuid;index"`
	Booking       bookingEntity.Booking `gorm:"foreignKey:BookingId"`
	Provider      string
	Method        string
	BankCode      string
	AccountNumber string
	Reference     string          `gorm:"uniqueIndex"`
	Amount        decimal.Decimal `gorm:"type:numeric"`
	Status        string
	ExpiredAt     time.Time
	PaidAt        *time.Time
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
