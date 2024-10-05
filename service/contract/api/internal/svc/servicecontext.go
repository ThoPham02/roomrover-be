package svc

import (
	"roomrover/service/contract/api/internal/config"
	"roomrover/service/contract/api/internal/middleware"
	"roomrover/service/contract/model"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	accountFunc "roomrover/service/account/function"
	inventFunc "roomrover/service/inventory/function"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
	ObjSync             sync.ObjSync

	ContractModel       model.ContractTblModel
	ContractRenterModel model.ContractRenterTblModel
	ContractDetailModel model.ContractDetailTblModel
	PaymentModel        model.PaymentTblModel

	AccountFunction accountFunc.AccountFunction
	InventFunction  inventFunc.InventoryFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		ObjSync:             *sync.NewObjSync(0),

		ContractModel:       model.NewContractTblModel(sqlx.NewMysql(c.DataSource)),
		ContractRenterModel: model.NewContractRenterTblModel(sqlx.NewMysql(c.DataSource)),
		ContractDetailModel: model.NewContractDetailTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentModel:        model.NewPaymentTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetAccountFunction(accountFunction accountFunc.AccountFunction) {
	sc.AccountFunction = accountFunction
}

func (sc *ServiceContext) SetInventFunction(inventFunction inventFunc.InventoryFunction) {
	sc.InventFunction = inventFunction
}
