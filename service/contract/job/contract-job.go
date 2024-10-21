package job

import (
	"flag"
	"fmt"
	"roomrover/service/contract/job/config"
	"roomrover/service/contract/job/handler"
	"roomrover/service/contract/job/svc"

	"github.com/robfig/cron/v3"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("payment-job-conf", "etc/payment-job.yaml", "the config file")

type PaymentScheduler struct {
	C    config.Config
	Cron *cron.Cron
	Ctx  *svc.ServiceContext
}

func NewPaymentScheduler() *PaymentScheduler {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	cr := cron.New()
	handler.RegisterCronjob(cr, ctx)

	return &PaymentScheduler{
		C:    c,
		Cron: cr,
		Ctx:  ctx,
	}
}
func (cs *PaymentScheduler) Start() error {
	fmt.Println("Star Payment Scheduler")
	cs.Cron.Start()

	return nil
}
