package svc

import (
	accountFunc "roomrover/service/account/function"
	contractFunc "roomrover/service/contract/function"
	inventFunc "roomrover/service/inventory/function"
	"roomrover/service/notification/api/internal/config"
	"roomrover/service/notification/api/internal/middleware"
	"roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
	NotificationModel   model.NotificationTblModel

	AccountFunction  accountFunc.AccountFunction
	InventFunction   inventFunc.InventoryFunction
	ContractFunction contractFunc.ContractFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		NotificationModel:   model.NewNotificationTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetAccountFunction(accountFunction accountFunc.AccountFunction) {
	sc.AccountFunction = accountFunction
}

func (sc *ServiceContext) SetInventFunction(inventFunction inventFunc.InventoryFunction) {
	sc.InventFunction = inventFunction
}

func (sc *ServiceContext) SetContractFunction(contractFunction contractFunc.ContractFunction) {
	sc.ContractFunction = contractFunction
}
