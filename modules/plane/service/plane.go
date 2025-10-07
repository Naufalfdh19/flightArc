package service

import (
	"context"
	"flight/modules/plane/entity"
	"flight/modules/plane/repo"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type PlaneService interface{
	AddPlane(ctx context.Context, plane entity.Plane) error 
}

type PlaneServiceImpl struct {
	r repo.PlaneRepo
}

func NewPlaneService(r repo.PlaneRepo) PlaneServiceImpl {
	return PlaneServiceImpl{
		r: r,
	}
}

func (s PlaneServiceImpl) AddPlane(ctx context.Context, plane entity.Plane) error {
	isPlaneExists := s.r.IsPlaneExistsByRegistrationCode(ctx, plane.RegistrationCode) 
	if isPlaneExists {
		return apperror.NewErrStatusBadRequest(constant.ADD_PLANE, apperror.ErrPlaneExists, apperror.ErrPlaneExists)
	}

	err := s.r.AddPlane(ctx, plane) 
	if err != nil {
		return err
	}

	return nil
}
