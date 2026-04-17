package repo

import (
	"context"
	"flight/modules/booking/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/transaction"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TicketRepo interface {
	AddTickets(ctx context.Context, tickets []entity.Ticket) error
	GetTotalPriceByBookingID(ctx context.Context, bookingID uuid.UUID) (decimal.Decimal, error)
}

type TicketRepoImpl struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepoImpl {
	return &TicketRepoImpl{
		db: db,
	}
}

func (r *TicketRepoImpl) AddTickets(ctx context.Context, tickets []entity.Ticket) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	err := db.WithContext(ctx).Model(entity.Ticket{}).
		Create(tickets).Error
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r *TicketRepoImpl) GetTotalPriceByBookingID(ctx context.Context, bookingID uuid.UUID) (decimal.Decimal, error) {
	var total decimal.Decimal

	if err := r.db.WithContext(ctx).
		Model(entity.Ticket{}).
		Select("COALESCE(SUM(price), 0)").
		Where("booking_id = ? AND deleted_at IS NULL", bookingID).
		Scan(&total).Error; err != nil {
		return decimal.Zero, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return total, nil
}
