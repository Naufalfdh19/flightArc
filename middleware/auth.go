package middleware

import (
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"flight/pkg/jwttoken"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth(ctx *gin.Context) {

	authHeader := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(authHeader) < 1 {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	token := authHeader[1]

	jwtToken := jwttoken.JwtTokenImpl{}
	jwtTokenClaims, err := jwtToken.ParseJwtTokenForAuth(ctx, token)
	if err != nil {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, err)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	if jwtTokenClaims.UserID == "" {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, err)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Set("user_id", jwtTokenClaims.UserID)
	ctx.Set("role", jwtTokenClaims.Role)

	ctx.Next()
}

func CheckUserAuth(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	if role != constant.USER {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func CheckAdminAuth(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	if role != constant.ADMIN {
		err := apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrUnauthorized, apperror.ErrUnauthorized)
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Next()
}
