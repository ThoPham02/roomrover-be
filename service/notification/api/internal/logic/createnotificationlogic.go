package logic

import (
	"context"

	"roomrover/service/notification/api/internal/svc"
	"roomrover/service/notification/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNotificationLogic {
	return &CreateNotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNotificationLogic) CreateNotification(req *types.CreateNotificationReq) (resp *types.CreateNotificationRes, err error) {
	// todo: add your logic here and delete this line

	return
}
