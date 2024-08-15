package api

import (
	"context"
	"roomrover/service/account/function"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.AccountFunction = (*AccountFunction)(nil)

type AccountFunction struct {
	function.AccountFunction
	logx.Logger
	AccountService *AccountService
}

func NewAccountFunction(svc *AccountService) *AccountFunction {
	ctx := context.Background()

	return &AccountFunction{
		Logger:         logx.WithContext(ctx),
		AccountService: svc,
	}
}

func (contractFunc *AccountFunction) Start() error {
	return nil
}
