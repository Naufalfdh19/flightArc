package common

import (
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(ctx *gin.Context) (int, error) {
	idRaw, exists := ctx.Get("user_id")
	if !exists {
		return 0, apperror.NewErrStatusUnauthorized(constant.CHECK_AUTH, apperror.ErrTokenInvalid, apperror.ErrTokenInvalid)
	}

	id, err := strconv.Atoi(idRaw.(string))
	if err != nil {
		return 0, apperror.NewErrStatusBadRequest(constant.UPDATE_PASSWORD, apperror.ErrConvertingType, apperror.ErrConvertingType)
	}

	return id, nil
}