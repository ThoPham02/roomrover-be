package handler

import (
	"fmt"
	"roomrover/service/payment/job/svc"
	"time"

	"github.com/robfig/cron/v3"
)

func RegisterCronjob(cr *cron.Cron, serverCtx *svc.ServiceContext) {

}

func Run(cr *cron.Cron) {
	select {
	case <-time.After(time.Hour * 24 * 365 * 100):
		fmt.Println("stoped! ")
	}
}
