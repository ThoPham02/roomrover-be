package main

import (
	"flag"
	"fmt"
	"roomrover/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	accountApi "roomrover/service/account/api"
	inventoryApi "roomrover/service/inventory/api"
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

	inventoryService := inventoryApi.NewInventoryService(server)
	inventoryService.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
