package repo

import (
	"context"
	"flight/modules/schedule/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/transaction"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeatRepo interface {
	IsSeatsAvailable(ctx context.Context, scheduleId uuid.UUID, seatNumbers []string) (bool, error) 
	UpdateSeatStatus(ctx context.Context, scheduleId uuid.UUID, seatNumbers []string, status string) error 
	GetReservedSeatNumbers(ctx context.Context, scheduleId uuid.UUID) ([]string, error) 
}

type SeatRepoImpl struct {
	db *gorm.DB
}

func NewSeatRepo(db *gorm.DB) *SeatRepoImpl {
	return &SeatRepoImpl{
		db: db,
	}
}

func (r *SeatRepoImpl) IsSeatsAvailable(ctx context.Context, scheduleId uuid.UUID, seatNumbers []string) (bool, error) {
	var count int64
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	log.Println(scheduleId, seatNumbers)

	err := db.WithContext(ctx).
		Model(&entity.Seat{}).
		Where("schedule_id = ? AND seat_number IN (?) AND status = ?",
			scheduleId, seatNumbers, "AVAILABLE").
		Set("gorm:query_option", "FOR UPDATE").
		Count(&count).Error
	if err != nil {
		return false, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return int(count) == len(seatNumbers), nil
}

func (r *SeatRepoImpl) UpdateSeatStatus(ctx context.Context, scheduleId uuid.UUID, seatNumbers []string, status string) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	err := db.WithContext(ctx).Model(entity.Seat{}).
		Where("schedule_id = ? and seat_number in (?) and deleted_at is null", scheduleId, seatNumbers).
		Updates(entity.Seat{
			Status: status,
		}).Set("gorm:query_option", "FOR UPDATE").
		Error
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r *SeatRepoImpl) GetReservedSeatNumbers(ctx context.Context, scheduleId uuid.UUID) ([]string, error) {
    var seatNumbers []string

    err := r.db.WithContext(ctx).
        Model(&entity.Seat{}).
        Where("schedule_id = ? AND status = ?", scheduleId, "RESERVED").
        // Pluck akan mengambil kolom 'seat_number' dan mengisi slice 'seatNumbers'
        Pluck("seat_number", &seatNumbers). 
        Error

    if err != nil {
        return nil, err
    }

    return seatNumbers, nil
}