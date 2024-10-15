package handler

import (
	"context"
	"roomrover/service/contract/job/logic"
	"roomrover/service/contract/job/svc"

	"github.com/robfig/cron/v3"
)

func CreateBillHandler(svcCtx *svc.ServiceContext) cron.FuncJob {
	return func() {
		ctx := context.Background()
		job := logic.NewCreateBillLogic(ctx, svcCtx)
		job.CreateBillByTime()
	}
}
