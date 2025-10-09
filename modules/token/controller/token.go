package controller

import (
	"flight/modules/token/service"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/wrapper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	s service.TokenService
}

func NewTokenController(s service.TokenService) TokenController {
	return TokenController{
		s: s,
	}
}

func (c TokenController) GenerateNewAccessToken(ctx *gin.Context) {
	authHeader := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(authHeader) < 1 {
		err := apperror.NewErrStatusUnauthorized(constant.GENERATE_TOKEN, apperror.ErrAccessTokenNotExists, apperror.ErrAccessTokenNotExists)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	accessToken := authHeader[1]

	newAccessToken, err := c.s.GenerateNewAccessToken(ctx, accessToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(newAccessToken, nil, "get new access token success"))
}