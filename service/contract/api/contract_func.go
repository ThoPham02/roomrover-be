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

func (contractFunc *ContractFunction) GetContractByTime(time int64) ([]*model.ContractTbl, error) {
	return contractFunc.ContractService.Ctx.ContractModel.GetContractByTime(context.Background(), time)
}

func (contractFunc *ContractFunction) GetContractDetailByContractID(contractID int64) ([]*model.ContractDetailTbl, error) {
	return contractFunc.ContractService.Ctx.ContractDetailModel.GetContractDetailByContractID(context.Background(), contractID)
}

func (contractFunc *ContractFunction) CountRenterByContractID(contractID int64) (int64, error) {
	return contractFunc.ContractService.Ctx.ContractRenterModel.CountRenterByContractID(context.Background(), contractID)
}

func (contractFunc *ContractFunction) UpdateContract(contract *model.ContractTbl) error {
	return contractFunc.ContractService.Ctx.ContractModel.Update(context.Background(), contract)
}

