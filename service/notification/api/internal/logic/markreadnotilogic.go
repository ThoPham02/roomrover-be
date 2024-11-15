package logic

import (
	"context"

	"roomrover/service/notification/api/internal/svc"
	"roomrover/service/notification/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkReadNotiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarkReadNotiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkReadNotiLogic {
	return &MarkReadNotiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkReadNotiLogic) MarkReadNoti(req *types.MarkReadReq) (resp *types.MarkReadRes, err error) {
	// todo: add your logic here and delete this line

	return
}
