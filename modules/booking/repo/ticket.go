package repo

import (
	"context"
	"flight/modules/booking/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/transaction"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepo interface{
	AddTickets(ctx context.Context, bookingId uuid.UUID, tickets []entity.Ticket) error 
}

type TicketRepoImpl struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepoImpl {
	return &TicketRepoImpl{
		db: db,
	}
}

func (r *TicketRepoImpl) AddTickets(ctx context.Context, bookingId uuid.UUID, tickets []entity.Ticket) error {
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