package controller

import (
	"flight/modules/payment/publisher"
)

type PaymentController struct {
	publisher publisher.PaymentPublisher
}

func NewPaymentController() PaymentController {
	return PaymentController{}
}

// func (c PaymentController) DoPayment(ctx *gin.Context) {
// 	c.publisher.DoPayment(ctx)
// }
