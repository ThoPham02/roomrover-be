package main

import (
	"flag"
	"fmt"
	"roomrover/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	accountApi "roomrover/service/account/api"
	contractApi "roomrover/service/contract/api"
	inventApi "roomrover/service/inventory/api"
	paymentApi "roomrover/service/payment/api"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

// @BasePath /
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))

	logx.DisableStat()
	defer server.Stop()

	accountService := accountApi.NewAccountService(server)
	accountService.Start()
	inventService := inventApi.NewInventService(server)
	inventService.Start()
	contractService := contractApi.NewContractService(server)
	contractService.Start()
	paymentService := paymentApi.NewPaymentService(server)
	paymentService.Start()

	contractFunc := contractApi.NewContractFunction(contractService)
	contractFunc.Start()

	inventService.Ctx.SetContractFunction(contractFunc)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
