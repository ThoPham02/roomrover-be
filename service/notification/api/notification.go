package api

import (
	"flag"
	"roomrover/service/notification/api/internal/config"
	"roomrover/service/notification/api/internal/handler"
	"roomrover/service/notification/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/notification-api.yaml", "the config file")

type NotificationService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewNotificationService(server *rest.Server) *NotificationService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &NotificationService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *NotificationService) Start() error {
	return nil
}
