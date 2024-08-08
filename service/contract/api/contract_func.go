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

func (contractFunc *ContractFunction) GetContractByRoomID(roomID int64) (*model.ContractTbl, error) {
	return contractFunc.ContractService.Ctx.ContractModel.GetContractByRoomID(context.Background(), roomID)
}
