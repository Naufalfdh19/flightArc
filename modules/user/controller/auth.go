package controller

import (
	"flight/modules/user/converter"
	"flight/modules/user/dto"
	"flight/pkg/apperror"
	"flight/pkg/common"
	"flight/pkg/constant"
	"flight/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (c UserController) GenerateNewAccessToken(ctx *gin.Context) {
	id, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var userDto dto.UpdatePasswordRequest
	err = ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	token, err := c.s.GenerateNewAccessToken(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.LoginResponse{Token: token}

	ctx.JSON(http.StatusOK, wrapper.Response(response, nil, "generate access token success"))
}
