package api

import (
	"flag"

	"roomrover/service/contract/api/internal/config"
	"roomrover/service/contract/api/internal/handler"
	"roomrover/service/contract/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("contract-api", "etc/contract-api.yaml", "the config file")

type ContractService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewContractService(server *rest.Server) *ContractService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &ContractService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *ContractService) Start() error {
	return nil
}
