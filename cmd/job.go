package main

import (
	"flag"
	"fmt"
	"roomrover/cmd/config"

	paymentScheduler "roomrover/service/payment/job"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("cronjob_config", "etc/cronjob-server.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	PaymentScheduler := paymentScheduler.NewPaymentScheduler()
	PaymentScheduler.Start()

	fmt.Println("Starting Scheduler ....... ")
	server.Start()
}
