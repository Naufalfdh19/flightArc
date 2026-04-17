package controller

import (
	"net/http"

	"flight/modules/payment/dto"
	"flight/modules/payment/service"
	"flight/pkg/apperror"
	"flight/pkg/common"
	paymentConstant "flight/pkg/constant"
	"flight/pkg/wrapper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentController struct {
	s service.PaymentService
}

func NewPaymentController(s service.PaymentService) PaymentController {
	return PaymentController{s: s}
}

func (c PaymentController) CreatePayment(ctx *gin.Context) {
	userID, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	var req dto.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(apperror.NewErrStatusBadRequest(paymentConstant.PAYMENT, apperror.ErrBindingRequest, apperror.ErrBindingRequest))
		return
	}

	payment, err := c.s.CreatePayment(ctx, userID, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, wrapper.Response(payment, nil, "create payment success"))
}

func (c PaymentController) GetPaymentByID(ctx *gin.Context) {
	userID, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	paymentID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(apperror.NewErrStatusBadRequest(paymentConstant.PAYMENT, apperror.ErrConvertingType, apperror.ErrConvertingType))
		return
	}

	payment, err := c.s.GetPaymentByID(ctx, userID, paymentID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(payment, nil, ""))
}

func (c PaymentController) HandleWebhook(ctx *gin.Context) {
	var req dto.PaymentWebhookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(apperror.NewErrStatusBadRequest(paymentConstant.PAYMENT, apperror.ErrBindingRequest, apperror.ErrBindingRequest))
		return
	}

	payment, err := c.s.HandleCallback(ctx, req, ctx.GetHeader("X-Payment-Signature"))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Response(payment, nil, "payment callback processed"))
}
