package main

import (
	"flag"
	"fmt"
	"roomrover/cmd/config"

	contractApi "roomrover/service/contract/api"
	inventApi "roomrover/service/inventory/api"
	notificationApi "roomrover/service/notification/api"

	paymentScheduler "roomrover/service/contract/job"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("cronjob_config", "etc/cronjob-server.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	logx.DisableStat()
	defer server.Stop()

	inventService := inventApi.NewInventService(server)
	inventService.Start()
	inventFunc := inventApi.NewInventoryFunction(inventService)
	inventFunc.Start()

	contractService := contractApi.NewContractService(server)
	contractService.Start()
	contractFunc := contractApi.NewContractFunction(contractService)
	contractFunc.Start()

	notiService := notificationApi.NewNotificationService(server)
	notiService.Start()
	notiFunc := notificationApi.NewNotificationFunction(notiService)
	notiFunc.Start()

	paymentScheduler := paymentScheduler.NewPaymentScheduler()
	paymentScheduler.Start()
	paymentScheduler.Ctx.SetInventFunction(inventFunc)
	paymentScheduler.Ctx.SetNotificationFunction(notiFunc)

	fmt.Println("Starting Scheduler ....... ")
	server.Start()
}
