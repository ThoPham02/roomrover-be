package svc

import (
	"roomrover/service/account/api/internal/config"
	"roomrover/service/account/model"
	"roomrover/sync"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	UserModel    model.UsersTblModel
	ProfileModel model.ProfilesTblModel

	ObjSync *sync.ObjSync
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		UserModel:    model.NewUsersTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.DataSource)),
		ProfileModel: model.NewProfilesTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.DataSource)),
		ObjSync:      sync.NewObjSync(1),
	}
}
