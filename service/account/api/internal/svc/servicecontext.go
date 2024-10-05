package svc

import (
	"roomrover/service/account/api/internal/config"
	"roomrover/service/account/model"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserTblModel
	ObjSync   *sync.ObjSync
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserTblModel(sqlx.NewMysql(c.DataSource)),
		ObjSync:   sync.NewObjSync(1),
	}
}
