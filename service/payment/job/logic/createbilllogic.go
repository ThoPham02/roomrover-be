package logic

import (
	"context"
	"roomrover/common"
	"roomrover/service/payment/job/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBillLogic {
	return &CreateBillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBillLogic) CreateBillByTime() error {
	var currentTime = common.GetCurrentTime()

	l.Logger.Info("Create bill by time: ", currentTime)

	return nil
}
