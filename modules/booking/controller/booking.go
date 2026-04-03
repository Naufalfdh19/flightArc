package controller

import (
	"flight/modules/booking/converter"
	"flight/modules/booking/dto"
	"flight/modules/booking/entity"
	"flight/modules/booking/queryparams"
	"flight/modules/booking/service"
	"flight/pkg/apperror"
	"flight/pkg/common"
	"flight/pkg/constant"
	"flight/pkg/pagination"
	"flight/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingController struct {
	s service.BookingService
}

func NewBookingController(s service.BookingService) BookingController {
	return BookingController{
		s: s,
	}
}

func (c BookingController) GetBookings(ctx *gin.Context) {
	id, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var queryParamsDto queryparams.QueryParamsDto

	if err := ctx.ShouldBindQuery(&queryParamsDto); err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_BOOKINGS, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}
	queryParams := queryparams.QueryParamsConverter{}.ConvertDtoToEntity(queryParamsDto)

	bookingsPagination, err := c.s.GetBookings(ctx, id, queryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	bookingsPaginationDto := pagination.Converter{}.ToDto(*bookingsPagination)
	bookings := bookingsPagination.Data.([]entity.Booking)
	var bookingsDto []dto.GetBookingReq
	for _, booking := range bookings {
		bookingDto := converter.GetBookingsConverter{}.ToDto(booking)
		bookingsDto = append(bookingsDto, bookingDto)
	}
	bookingsPaginationDto.Data = bookingsDto

	ctx.JSON(http.StatusOK, wrapper.Response(bookingsPaginationDto, nil, ""))
}

func (c BookingController) AddBookings(ctx *gin.Context) {
	id, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var bookingReq dto.AddBookingReq
	err = ctx.ShouldBindJSON(&bookingReq)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.ADD_BOOKINGS, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}

	bookingReq.UserId = id
	err = c.s.AddBookings(ctx, bookingReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, wrapper.Response(nil, nil, "add bookings success"))
}

func (c BookingController) GetBookingsById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idParam, err := uuid.Parse(idStr)
	if err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_BOOKING_BY_ID, apperror.ErrConvertingType, apperror.ErrConvertingType)
		ctx.Error(err)
		return
	}

	bookingDto, err := c.s.GetBookingsById(ctx, idParam)

	ctx.JSON(http.StatusOK, wrapper.Response(bookingDto, nil, ""))
}
