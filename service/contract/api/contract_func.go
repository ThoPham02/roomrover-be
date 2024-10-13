package api

import (
	"context"
	"roomrover/service/contract/function"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.ContractFunction = (*ContractFunction)(nil)

type ContractFunction struct {
	function.ContractFunction
	logx.Logger
	ContractService *ContractService
}

func NewContractFunction(svc *ContractService) *ContractFunction {
	ctx := context.Background()

	return &ContractFunction{
		Logger:          logx.WithContext(ctx),
		ContractService: svc,
	}
}

func (contractFunc *ContractFunction) Start() error {
	return nil
}

func (contractFunc *ContractFunction) GetPaymentByTime(time int64) (payments []*model.PaymentTbl, err error) {
	return contractFunc.ContractService.Ctx.PaymentModel.FindMultiByTime(context.TODO(), time)
}

func (contractFunc *ContractFunction) GetContractByID(contractID int64) (contract *model.ContractTbl, err error) {
    return contractFunc.ContractService.Ctx.ContractModel.FindOne(context.TODO(), contractID)
}
