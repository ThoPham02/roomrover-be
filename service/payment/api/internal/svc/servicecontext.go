package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"roomrover/service/payment/api/internal/config"
	"roomrover/service/payment/api/internal/middleware"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
	}
}
