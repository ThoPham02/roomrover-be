package svc

import (
	"roomrover/service/contract/api/internal/config"
	"roomrover/service/contract/api/internal/middleware"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware

	ContractModel          model.ContractTblModel
	ContractRenterTblModel model.ContractRenterTblModel
	ContractDetailTblModel model.ContractDetailTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,

		ContractModel:          model.NewContractTblModel(sqlx.NewMysql(c.DataSource)),
		ContractRenterTblModel: model.NewContractRenterTblModel(sqlx.NewMysql(c.DataSource)),
		ContractDetailTblModel: model.NewContractDetailTblModel(sqlx.NewMysql(c.DataSource)),
	}
}
