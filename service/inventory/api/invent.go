package api

import (
	"flag"

	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/api/internal/handler"
	"roomrover/service/inventory/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("invent-api", "etc/invent-api.yaml", "the config file")

type InventService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewInventService(server *rest.Server) *InventService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &InventService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *InventService) Start() error {
	return nil
}
