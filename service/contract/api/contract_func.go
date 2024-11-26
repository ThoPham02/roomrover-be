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

func (cfFunc *ContractFunction) CountContractByHouseID(houseID int64) (count int64, err error) {
	return cfFunc.ContractService.Ctx.ContractModel.CountByHouseID(context.TODO(), houseID)
}
func (cfFunc *ContractFunction) GetContractByID(contractID int64) (contract *model.ContractTbl, err error) {
	return cfFunc.ContractService.Ctx.ContractModel.FindOne(context.TODO(), contractID)
}

func (cfFunc *ContractFunction) GetActiveContractByRoomID(roomID int64) (contract *model.ContractTbl, err error) {
	return cfFunc.ContractService.Ctx.ContractModel.FindActiveByRoomID(context.TODO(), roomID)
}

func (cfFunc *ContractFunction) GetPaymentByContractID(contractID int64) (payments *model.PaymentTbl, err error) {
	return cfFunc.ContractService.Ctx.PaymentModel.FindByContractID(context.TODO(), contractID)
}

func (cfFunc *ContractFunction) GetBillPayByContractID(userID int64) (bills []*model.BillPayTbl, err error) {
	return cfFunc.ContractService.Ctx.BillPayModel.FindByContractID(context.TODO(), userID)
}

func (cfFunc *ContractFunction) GetContractByRoom(roomID int64) ([]*model.ContractTbl, error) {
	return cfFunc.ContractService.Ctx.ContractModel.FindByRoomID(context.TODO(), roomID)
}

// func (cfFunc *ContractFunction) GetPaymentByTime(time int64) (payments []*model.PaymentTbl, err error) {
// 	return cfFunc.ContractService.Ctx.PaymentModel.FindMultiByTime(context.TODO(), time)
// }
