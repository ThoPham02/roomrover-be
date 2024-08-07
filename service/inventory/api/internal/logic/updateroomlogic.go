package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update room
func NewUpdateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoomLogic {
	return &UpdateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoomLogic) UpdateRoom(req *types.UpdateRoomReq) (resp *types.UpdateRoomRes, err error) {
	// todo: add your logic here and delete this line

	return
}
