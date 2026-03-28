package service

import (
	"context"
	"flight/modules/booking/queryparams"
	"flight/modules/booking/repo"
	"flight/pkg/pagination"
)

type BookingService interface{
	GetBookings(ctx context.Context, userId int, queryParams queryparams.QueryParams) (*pagination.Pagination, error) 
}

type BookingServiceImpl struct {
	r repo.BookingRepo
}

func NewBookingService(r repo.BookingRepo) *BookingServiceImpl {
	return &BookingServiceImpl{
		r: r,
	}
}

func (s *BookingServiceImpl) GetBookings(ctx context.Context, userId int, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	queryparams.CheckLimit(&queryParams)

	users, totalUsers, err := s.r.GetBookings(ctx, userId, queryParams)
	if err != nil {
		return nil, err
	}

	totalPage := totalUsers / queryParams.Limit
	if totalUsers % queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	pagination := pagination.Pagination{
		Page:         queryParams.Page,
		TotalElement: totalUsers,
		TotalPage:    totalPage,
		Data:         users,
	}

	return &pagination, nil
}
