package controller

import (
	"flight/modules/plane/converter"
	"flight/modules/plane/dto"
	"flight/modules/plane/service"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/wrapper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaneController struct {
	s service.PlaneService
}

func NewPlaneController(s service.PlaneService) PlaneController {
	return PlaneController{
		s: s,
	}
}

func (c PlaneController) AddPlane(ctx *gin.Context) {
	var planeDto dto.AddPlaneRequest
	err := ctx.ShouldBindJSON(&planeDto)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.ADD_PLANE, apperror.ErrBindingRequest, err)
		ctx.Error(err)
		return
	}

	plane := converter.AddPlaneConverter{}.ToEntity(planeDto)

	err = c.s.AddPlane(ctx, plane)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, wrapper.Response(nil, nil, "add plane success"))
}

func (c PlaneController) UpdateSeats(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.UPDATE_SEATS, apperror.ErrConvertingType, err)
		ctx.Error(err)
		return
	}

	var seatsDto dto.UpdateSeats
	err = ctx.ShouldBindJSON(&seatsDto)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.UPDATE_SEATS, apperror.ErrBindingRequest, err)
		ctx.Error(err)
		return
	}

	plane := converter.UpdateSeatsConverter{}.ToEntity(seatsDto)
	plane.Id = id

	err = c.s.UpdateSeats(ctx, plane)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, wrapper.Response(nil, nil, "update seats success"))
}
