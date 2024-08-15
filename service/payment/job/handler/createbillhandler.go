package handler

import (
	"context"
	"roomrover/service/payment/job/logic"
	"roomrover/service/payment/job/svc"

	"github.com/robfig/cron/v3"
)

func CreateBillHandler(svcCtx *svc.ServiceContext) cron.FuncJob {
	return func() {
		ctx := context.Background()
		job := logic.NewCreateBillLogic(ctx, svcCtx)
		job.CreateBillByTime()
	}
}
