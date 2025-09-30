package service

import (
	"context"
	"flight/modules/schedule/queryparams"
	"flight/modules/schedule/repo"
	"flight/pkg/pagination"
)

type ScheduleService interface {
	GetSchedules(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error) 
}

type ScheduleServiceImpl struct {
	r repo.ScheduleRepo
}

func NewScheduleService(r repo.ScheduleRepo) ScheduleServiceImpl {
	return ScheduleServiceImpl{
		r: r,
	}
}

func (u ScheduleServiceImpl) GetSchedules(ctx context.Context, queryParams queryparams.QueryParams) (*pagination.Pagination, error) {
	totalSchedule, err := u.r.GetTotalSchedule(ctx)
	if err != nil {
		return nil, err
	}

	queryparams.CheckLimit(&queryParams)
	totalPage := totalSchedule / queryParams.Limit
	if totalSchedule%queryParams.Limit != 0 {
		totalPage += 1
	}
	queryparams.CheckPage(&queryParams, totalPage)

	schedules, err := u.r.GetSchedules(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	pagination := pagination.Pagination{
		Page:         queryParams.Page,
		TotalElement: totalSchedule,
		TotalPage:    totalPage,
		Data:         schedules,
	}

	return &pagination, nil
}