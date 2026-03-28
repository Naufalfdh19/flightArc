package repo

import (
	"context"
	"flight/modules/booking/entity"
	"flight/modules/booking/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/common"
	"flight/pkg/constant"

	"gorm.io/gorm"
)

type BookingRepo interface{
	GetBookings(ctx context.Context, userId int, queryparams queryparams.QueryParams) ([]entity.Booking, int, error) 
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
        return nil, 0, err
    }

	err := query.
			Select("bookings.id","user_id","schedule_id","bookings.status","booking_time").
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