package svc

import (
	"roomrover/service/contract/api/internal/config"
	"roomrover/service/contract/model"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	ObjSync *sync.ObjSync

	ContractModel          model.ContractTblModel
	ContractRenterTblModel model.ContractRenterTblModel
	ContractDetailTblModel model.ContractDetailTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ObjSync: sync.NewObjSync(0),

		ContractModel:          model.NewContractTblModel(sqlx.NewMysql(c.DataSource)),
		ContractRenterTblModel: model.NewContractRenterTblModel(sqlx.NewMysql(c.DataSource)),
		ContractDetailTblModel: model.NewContractDetailTblModel(sqlx.NewMysql(c.DataSource)),
	}
}
