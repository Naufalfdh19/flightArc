package service

import (
	"context"
	"flight/modules/plane/entity"
	"flight/modules/plane/repo"
	airlineRepo "flight/modules/airline/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type PlaneService interface{
	AddPlane(ctx context.Context, plane entity.Plane) error 
	UpdateSeats(ctx context.Context, plane entity.Plane) error 
}

type PlaneServiceImpl struct {
	pr repo.PlaneRepo
	ar airlineRepo.AirlineRepo
}

func NewPlaneService(pr repo.PlaneRepo, ar airlineRepo.AirlineRepo) PlaneServiceImpl {
	return PlaneServiceImpl{
		pr: pr,
		ar: ar,
	}
}

func (s PlaneServiceImpl) AddPlane(ctx context.Context, plane entity.Plane) error {
	isPlaneExists := s.pr.IsPlaneExistsByRegistrationCode(ctx, plane.RegistrationCode) 
	if isPlaneExists {
		return apperror.NewErrStatusBadRequest(constant.ADD_PLANE, apperror.ErrPlaneExists, apperror.ErrPlaneExists)
	}
	
	isAirlineExists := s.ar.IsAirlineExistsById(ctx, plane.AirlineId)
	if !isAirlineExists {
		return apperror.NewErrStatusBadRequest(constant.ADD_PLANE, apperror.ErrAirlineNotExists, apperror.ErrAirlineNotExists)
	}

	err := s.pr.AddPlane(ctx, plane) 
	if err != nil {
		return err
	}

	return nil
}

func (s PlaneServiceImpl) UpdateSeats(ctx context.Context, plane entity.Plane) error {
	isPlaneExists := s.pr.IsPlaneExistsById(ctx, plane.Id) 
	if !isPlaneExists {
		return apperror.NewErrStatusBadRequest(constant.UPDATE_SEATS, apperror.ErrPlaneNotExists, apperror.ErrPlaneNotExists)
	}

	err := s.pr.UpdateSeats(ctx, plane) 
	if err != nil {
		return err
	}

	return nil
}



