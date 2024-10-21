package handler

import (
	"fmt"
	"roomrover/service/contract/job/svc"
	"time"

	"github.com/robfig/cron/v3"
)

func RegisterCronjob(cr *cron.Cron, serverCtx *svc.ServiceContext) {
	createBillByTime := CreateBillHandler(serverCtx)
	cr.AddFunc(serverCtx.Config.CreateBillByTimeJob.Time, createBillByTime)
}

func Run(cr *cron.Cron) {
	time.Sleep(time.Hour * 24 * 365 * 100)
	fmt.Println("stopped!")
}
