package api

import (
	"flag"

	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/api/internal/handler"
	"roomrover/service/inventory/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("inventory-api", "etc/inventory-api.yaml", "the config file")

type InventoryService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewInventoryService(server *rest.Server) *InventoryService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &InventoryService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *InventoryService) Start() error {
	return nil
}
