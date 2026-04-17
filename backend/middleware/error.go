package middleware

import (
	"errors"
	"net/http"
	"flight/pkg/apperror"
	"flight/pkg/wrapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware(ctx *gin.Context) {
	ctx.Next()

	var errorDetails []apperror.ErrorDetail

	if len(ctx.Errors) > 0 {
		var ve validator.ValidationErrors
		if errors.As(ctx.Errors[0], &ve) {
			for _, fe := range ve {
				errorDetails = append(errorDetails, apperror.ErrorDetail{
					Field: fe.Field(),
					Message: apperror.ExtractValidationError(fe),
				})
			}
			ctx.JSON(http.StatusBadRequest, wrapper.Response(nil, errorDetails, ""))
		}

		var er *apperror.ErrorStruct
		if errors.As(ctx.Errors[0], &er) {
			errordetail := apperror.ErrorDetail{
				Field: er.Field,
				Message: er.Message,
			}
			ctx.AbortWithStatusJSON(er.Status, wrapper.Response(nil, []apperror.ErrorDetail{errordetail}, er.Message))
		}

	}
}