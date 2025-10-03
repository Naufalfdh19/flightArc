package controller

import (
	"flight/modules/user/converter"
	"flight/modules/user/dto"
	"flight/modules/user/entity"
	"flight/modules/user/queryparams"
	"flight/modules/user/service"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/pagination"
	"flight/pkg/wrapper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	s service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		s: s,
	}
}

func (c UserController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_USER_BY_ID, apperror.ErrConvertingType, apperror.ErrConvertingType)
		ctx.Error(err)
		return
	}
	user, err := c.s.GetUserById(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	userDto := converter.GetUserConverter{}.ToDto(*user)

	ctx.JSON(http.StatusOK, wrapper.Response(userDto, nil, ""))
}

func (c UserController) GetUsers(ctx *gin.Context) {
	var queryParamsDto queryparams.QueryParamsDto

	if err := ctx.ShouldBindQuery(&queryParamsDto); err != nil {
		err = apperror.NewErrInternalServerError(constant.GET_USERS, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}
	queryParams := queryparams.QueryParamsConverter{}.ConvertDtoToEntity(queryParamsDto)

	usersPagination, err := c.s.GetUsers(ctx, queryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	usersPaginationDto := pagination.Converter{}.ToDto(*usersPagination)
	users := usersPagination.Data.([]entity.User)
	var usersDto []dto.GetUserResponse
	for _, user := range users {
		userDto := converter.GetUserConverter{}.ToDto(user)
		usersDto = append(usersDto, userDto)
	}
	usersPaginationDto.Data = usersDto

	ctx.JSON(http.StatusOK, wrapper.Response(usersPaginationDto, nil, ""))
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	idRaw, exists := ctx.Get("user_id")
	if !exists {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
		ctx.Error(err)
		return
	}

	id, err := strconv.Atoi(idRaw.(string))
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.DELETE_USER_BY_ID, apperror.ErrConvertingType, apperror.ErrConvertingType)
		ctx.Error(err)
		return
	}

	var userDto dto.UpdateUserRequest
	err = ctx.ShouldBindJSON(&userDto)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.UPDATE_USER_BY_ID, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}

	user := converter.UpdateUserConverter{}.ToEntity(userDto)
	user.Id = id

	err = c.s.UpdateUserById(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(nil, nil, "update success"))
}

func (c UserController) DeleteUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.DELETE_USER_BY_ID, apperror.ErrConvertingType, apperror.ErrConvertingType)
		ctx.Error(err)
		return
	}

	err = c.s.DeleteUserById(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(nil, nil, "delete success"))
}

func (c UserController) Login(ctx *gin.Context) {
	var userDto dto.LoginRequest
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.LOGIN, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}

	user := converter.LoginRequestConverter{}.ToEntity(userDto)
	token, err := c.s.Login(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.LoginResponse{Token: token}

	ctx.JSON(http.StatusOK, wrapper.Response(response, nil, "login success"))
}

func (c UserController) Register(ctx *gin.Context) {
	var userDto dto.AddUserRequest
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.REGISTER, apperror.ErrBindingRequest, apperror.ErrBindingRequest)
		ctx.Error(err)
		return
	}

	user := converter.RegisterRequestConverter{}.ToEntity(userDto)

	err = c.s.Register(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, wrapper.Response(nil, nil, "register success"))
}

func (c UserController) UpdatePassword(ctx *gin.Context) {
	idRaw, exists := ctx.Get("user_id")
	if !exists {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
		ctx.Error(err)
		return
	}

	id, err := strconv.Atoi(idRaw.(string))
	if err != nil {
		err = apperror.NewErrStatusBadRequest(constant.UPDATE_PASSWORD, apperror.ErrConvertingType, apperror.ErrConvertingType)
		ctx.Error(err)
		return
	}

	var userDto dto.UpdatePasswordRequest
	err = ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	user := converter.UpdatePasswordConverter{}.ToEntity(userDto)
	user.Id = id

	err = c.s.UpdatePassword(ctx, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(nil, nil, "update password success"))
}
