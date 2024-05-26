package api

import (
	"flag"

	"roomrover/service/account/api/internal/config"
	"roomrover/service/account/api/internal/handler"
	"roomrover/service/account/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("account-api", "etc/account-api.yaml", "the config file")

type AccountService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewAccountService(server *rest.Server) *AccountService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &AccountService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *AccountService) Start() error {
	return nil
}
