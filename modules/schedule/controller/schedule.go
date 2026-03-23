package controller

import (
	"flight/modules/schedule/converter"
	"flight/modules/schedule/dto"
	"flight/modules/schedule/entity"
	"flight/modules/schedule/queryparams"
	"flight/modules/schedule/service"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/pagination"
	"flight/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	s service.ScheduleService
}

func NewScheduleController(s service.ScheduleService) ScheduleController {
	return ScheduleController{
		s: s,
	}
}

func (c ScheduleController) GetFlights(ctx *gin.Context) {
	var queryParamsDto queryparams.QueryParamsDto

	if err := ctx.ShouldBindQuery(&queryParamsDto); err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_USERS, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}
	queryParams := queryparams.QueryParamsConverter{}.ConvertDtoToEntity(queryParamsDto)

	flightsPagination, err := c.s.GetFlights(ctx, queryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	flightsPaginationDto := pagination.Converter{}.ToDto(*flightsPagination)
	flights := flightsPagination.Data.([]entity.Flight)
	var flightsDto []dto.GetFlightDto
	for _, flight := range flights {
		flightDto := converter.GetFlightConverter{}.ToDto(flight)
		flightsDto = append(flightsDto, flightDto)
	}
	flightsPaginationDto.Data = flightsDto

	ctx.JSON(http.StatusOK, wrapper.Response(flightsPaginationDto, nil, ""))
}
