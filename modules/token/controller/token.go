package controller

import (
	"flight/modules/token/service"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/wrapper"
	"net/http"

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
	refreshToken, err := getRefreshToken(ctx)
	if err != nil {
		err := apperror.NewErrStatusNotFound(constant.GENERATE_TOKEN, apperror.ErrRefreshTokenNotExists, apperror.ErrRefreshTokenNotExists)
		ctx.Error(err)
		return
	}

	newAccessToken, err := c.s.GenerateNewAccessToken(ctx, refreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(newAccessToken, nil, "get new access token success"))
}

func SetRefreshTokenCookie(c *gin.Context, token string) {
	c.SetCookie(
		"refresh_token",      // name
		token,               // value
		7*24*60*60,          // maxAge in seconds (7 days)
		"/api/v1/tokens", // path
		"",                  // domain ("" for localhost)
		true,                // secure (true = HTTPS only)
		true,                // httpOnly
	)
}

func getRefreshToken(c *gin.Context) (string, error) {
	return c.Cookie("refresh_token")
}
