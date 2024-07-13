package svc

import (
	"roomrover/service/account/api/internal/config"
	"roomrover/service/account/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.DataSource)),
	}
}
