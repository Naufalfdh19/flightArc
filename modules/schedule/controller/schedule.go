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

func (c ScheduleController) GetSchedules(ctx *gin.Context) {
	var queryParamsDto queryparams.QueryParamsDto

	if err := ctx.ShouldBindQuery(&queryParamsDto); err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_USERS, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}
	queryParams := queryparams.QueryParamsConverter{}.ConvertDtoToEntity(queryParamsDto)

	schedulesPagination, err := c.s.GetSchedules(ctx, queryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	schedulesPaginationDto := pagination.Converter{}.ToDto(*schedulesPagination)
	users := schedulesPagination.Data.([]entity.Schedule)
	var usersDto []dto.GetScheduleDto
	for _, user := range users {
		userDto := converter.GetScheduleConverter{}.ToDto(user)
		usersDto = append(usersDto, userDto)
	}
	schedulesPaginationDto.Data = usersDto

	ctx.JSON(http.StatusOK, wrapper.Response(schedulesPaginationDto, nil, ""))
}
