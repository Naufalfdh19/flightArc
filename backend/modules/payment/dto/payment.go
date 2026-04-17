package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreatePaymentRequest struct {
	BookingId uuid.UUID `json:"booking_id" binding:"required"`
	BankCode  string    `json:"bank_code" binding:"required"`
}

type PaymentWebhookRequest struct {
	Reference string `json:"reference" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

type PaymentResponse struct {
	Id                uuid.UUID       `json:"id"`
	BookingId         uuid.UUID       `json:"booking_id"`
	Provider          string          `json:"provider"`
	Method            string          `json:"method"`
	BankCode          string          `json:"bank_code"`
	AccountNumber     string          `json:"account_number"`
	Reference         string          `json:"reference"`
	Amount            decimal.Decimal `json:"amount"`
	Status            string          `json:"status"`
	ExpiredAt         time.Time       `json:"expired_at"`
	PaidAt            *time.Time      `json:"paid_at,omitempty"`
	Instruction       string          `json:"instruction"`
	CallbackSignature string          `json:"callback_signature,omitempty"`
}
