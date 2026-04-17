package controller

import (
	"flight/modules/admin/service"
	"flight/modules/user/converter"
	"flight/modules/user/dto"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	s service.AdminService
}

func NewAdminController(s service.AdminService) AdminController {
	return AdminController{
		s: s,
	}
}

func (c AdminController) Login(ctx *gin.Context) {
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

func (c AdminController) Register(ctx *gin.Context) {
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
