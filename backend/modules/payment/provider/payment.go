package provider

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	paymentConstant "flight/modules/payment/constant"
	paymentEntity "flight/modules/payment/entity"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateVirtualAccountInput struct {
	BookingId uuid.UUID
	BankCode  string
	Amount    decimal.Decimal
}

type CreateVirtualAccountOutput struct {
	Provider      string
	Method        string
	BankCode      string
	AccountNumber string
	Reference     string
	Amount        decimal.Decimal
	ExpiredAt     time.Time
	Instruction   string
}

type PaymentProvider interface {
	CreateVirtualAccount(input CreateVirtualAccountInput) (*CreateVirtualAccountOutput, error)
	BuildCallbackSignature(payment paymentEntity.Payment, status string) string
	VerifyCallbackSignature(payment paymentEntity.Payment, status, signature string) bool
}

type ManualVAProvider struct{}

func NewManualVAProvider() ManualVAProvider {
	return ManualVAProvider{}
}

func (p ManualVAProvider) CreateVirtualAccount(input CreateVirtualAccountInput) (*CreateVirtualAccountOutput, error) {
	bankCode := strings.ToUpper(strings.TrimSpace(input.BankCode))
	accountNumber := buildAccountNumber(bankCode, input.BookingId)
	reference := fmt.Sprintf("%s-%s", bankCode, strings.ToUpper(uuid.NewString()))

	return &CreateVirtualAccountOutput{
		Provider:      paymentConstant.ProviderManualVA,
		Method:        paymentConstant.MethodVirtualVA,
		BankCode:      bankCode,
		AccountNumber: accountNumber,
		Reference:     reference,
		Amount:        input.Amount,
		ExpiredAt:     time.Now().Add(15 * time.Minute),
		Instruction:   fmt.Sprintf("Transfer to %s virtual account %s before expiry, then send the payment callback to confirm settlement.", bankCode, accountNumber),
	}, nil
}

func (p ManualVAProvider) BuildCallbackSignature(payment paymentEntity.Payment, status string) string {
	secret := os.Getenv("PAYMENT_WEBHOOK_SECRET")
	if secret == "" {
		return ""
	}

	payload := fmt.Sprintf("%s|%s", payment.Reference, strings.ToUpper(status))
	return signPayload(secret, payload)
}

func (p ManualVAProvider) VerifyCallbackSignature(payment paymentEntity.Payment, status, signature string) bool {
	secret := os.Getenv("PAYMENT_WEBHOOK_SECRET")
	if secret == "" {
		return true
	}

	expected := p.BuildCallbackSignature(payment, status)
	return hmac.Equal([]byte(expected), []byte(signature))
}

func buildAccountNumber(bankCode string, bookingID uuid.UUID) string {
	cleaned := strings.ReplaceAll(bookingID.String(), "-", "")
	if len(cleaned) > 12 {
		cleaned = cleaned[:12]
	}

	prefixes := map[string]string{
		"BCA":     "014",
		"MANDIRI": "008",
		"BNI":     "009",
		"BRI":     "002",
	}

	prefix, exists := prefixes[bankCode]
	if !exists {
		prefix = "999"
	}

	return prefix + cleaned
}

func signPayload(secret, payload string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
