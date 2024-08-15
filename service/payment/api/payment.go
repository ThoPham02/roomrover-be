package api

import (
	"flag"

	"roomrover/service/payment/api/internal/config"
	"roomrover/service/payment/api/internal/handler"
	"roomrover/service/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/payment-api.yaml", "the config file")

type PaymentService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewPaymentService(server *rest.Server) *PaymentService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &PaymentService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *PaymentService) Start() error {
	return nil
}
