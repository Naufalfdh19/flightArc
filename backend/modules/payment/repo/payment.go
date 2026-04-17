package repo

import (
	"context"
	paymentEntity "flight/modules/payment/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/transaction"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	Create(ctx context.Context, payment *paymentEntity.Payment) error
	GetByID(ctx context.Context, paymentID uuid.UUID) (*paymentEntity.Payment, error)
	GetByReference(ctx context.Context, reference string) (*paymentEntity.Payment, error)
	GetActiveByBookingID(ctx context.Context, bookingID uuid.UUID) (*paymentEntity.Payment, error)
	Update(ctx context.Context, payment paymentEntity.Payment) error
}

type PaymentRepoImpl struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepoImpl {
	return &PaymentRepoImpl{db: db}
}

func (r *PaymentRepoImpl) Create(ctx context.Context, payment *paymentEntity.Payment) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	if err := db.WithContext(ctx).Create(payment).Error; err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r *PaymentRepoImpl) GetByID(ctx context.Context, paymentID uuid.UUID) (*paymentEntity.Payment, error) {
	var payment paymentEntity.Payment

	if err := r.db.WithContext(ctx).Model(paymentEntity.Payment{}).
		Where("id = ? AND deleted_at IS NULL", paymentID).
		First(&payment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.NewErrStatusNotFound(constant.SERVER, apperror.ErrPaymentNotExists, err)
		}
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &payment, nil
}

func (r *PaymentRepoImpl) GetByReference(ctx context.Context, reference string) (*paymentEntity.Payment, error) {
	var payment paymentEntity.Payment

	if err := r.db.WithContext(ctx).Model(paymentEntity.Payment{}).
		Where("reference = ? AND deleted_at IS NULL", reference).
		First(&payment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.NewErrStatusNotFound(constant.SERVER, apperror.ErrPaymentNotExists, err)
		}
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &payment, nil
}

func (r *PaymentRepoImpl) GetActiveByBookingID(ctx context.Context, bookingID uuid.UUID) (*paymentEntity.Payment, error) {
	var payment paymentEntity.Payment

	err := r.db.WithContext(ctx).Model(paymentEntity.Payment{}).
		Where("booking_id = ? AND status = ? AND expired_at > ? AND deleted_at IS NULL", bookingID, "PENDING", time.Now()).
		Order("created_at DESC").
		First(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return &payment, nil
}

func (r *PaymentRepoImpl) Update(ctx context.Context, payment paymentEntity.Payment) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	if err := db.WithContext(ctx).Model(paymentEntity.Payment{}).
		Where("id = ?", payment.Id).
		Updates(payment).Error; err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}
