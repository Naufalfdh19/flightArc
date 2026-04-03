package repo

import (
	"context"
	"flight/modules/booking/entity"
	"flight/modules/booking/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/common"
	"flight/pkg/constant"
	"flight/pkg/transaction"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepo interface {
	GetBookings(ctx context.Context, userId int, queryparams queryparams.QueryParams) ([]entity.Booking, int, error)
	AddBookings(ctx context.Context, booking *entity.Booking) error
	GetBookingById(ctx context.Context, bookingId uuid.UUID) (*entity.Booking, error)
	IsBookingExists(ctx context.Context, bookingId uuid.UUID) bool
	UpdateBooking(ctx context.Context, booking entity.Booking) error 
}

type BookingRepoImpl struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) *BookingRepoImpl {
	return &BookingRepoImpl{
		db: db,
	}
}

func (r *BookingRepoImpl) GetBookings(ctx context.Context, userId int, queryparams queryparams.QueryParams) ([]entity.Booking, int, error) {
	var bookings []entity.Booking
	var total int64

	query := r.db.Model(entity.Booking{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	err := query.
		Select("bookings.id", "user_id", "schedule_id", "bookings.status", "booking_time").
		Joins("User").Joins("Flight").Joins("Flight.Origin").Joins("Flight.Destination").
		Where("user_id = ?", userId).
		Preload("User").Preload("Flight.Origin").Preload("Flight.Destination").
		Scopes(common.Paginate(queryparams.Page, queryparams.Limit, int(total))).
		Find(&bookings).Error
	if err != nil {
		return nil, 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return bookings, int(total), nil
}

func (r *BookingRepoImpl) GetBookingById(ctx context.Context, bookingId uuid.UUID) (*entity.Booking, error) {
	var booking entity.Booking

	err := r.db.Model(entity.Booking{}).Joins("User").Joins("Flight").Joins("Flight.Origin").Joins("Flight.Destination").
		Where("bookings.id = ?", bookingId).
		Preload("User").Preload("Flight.Origin").Preload("Flight.Destination").
		First(&booking).Error
	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	log.Println("booking: ", booking)

	return &booking, nil
}

func (r *BookingRepoImpl) AddBookings(ctx context.Context, booking *entity.Booking) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	err := db.WithContext(ctx).Create(booking).Scan(&booking).Error
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r *BookingRepoImpl) IsBookingExists(ctx context.Context, bookingId uuid.UUID) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM bookings 
		WHERE id = ? AND deleted_at IS NULL)`
	_ = r.db.Raw(query, bookingId).Scan(&exists)
	return exists
}

func (r *BookingRepoImpl) UpdateBooking(ctx context.Context, booking entity.Booking) error {
	db := transaction.ExtractTx(ctx)
	if db == nil {
		db = r.db
	}

	err := db.WithContext(ctx).Model(entity.Booking{}).
			Where("id = ?", booking.Id).
			Updates(booking).Error
	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}
