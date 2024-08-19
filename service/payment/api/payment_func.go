package api

import (
	"context"
	"roomrover/service/payment/function"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.PaymentFunction = (*PaymentFunction)(nil)

type PaymentFunction struct {
	function.PaymentFunction
	logx.Logger
	PaymentService *PaymentService
}

func NewPaymentFunction(svc *PaymentService) *PaymentFunction {
	ctx := context.Background()

	return &PaymentFunction{
		Logger:         logx.WithContext(ctx),
		PaymentService: svc,
	}
}

func (contractFunc *PaymentFunction) Start() error {
	return nil
}
