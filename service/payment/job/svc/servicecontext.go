package svc

import (
	"roomrover/service/payment/job/config"
	"roomrover/service/payment/model"
	"roomrover/sync"

	contractFunc "roomrover/service/contract/function"
	inventFunc "roomrover/service/inventory/function"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	ObjSync sync.ObjSync

	BillModel       model.BillModel
	BillDetailModel model.BillDetailModel

	InventFunction inventFunc.InventoryFunction
	ContractFunc   contractFunc.ContractFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ObjSync: *sync.NewObjSync(0),

		BillModel:       model.NewBillModel(sqlx.NewMysql(c.DataSource)),
		BillDetailModel: model.NewBillDetailModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetInventFunction(inventFunction inventFunc.InventoryFunction) {
	sc.InventFunction = inventFunction
}

func (ctx *ServiceContext) SetContractFunction(contractFunction contractFunc.ContractFunction) {
	ctx.ContractFunc = contractFunction
}
