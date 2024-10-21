package svc

import (
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		NotificationModel:   model.NewNotificationTblModel(sqlx.NewMysql(c.DataSource)),
	}
}
