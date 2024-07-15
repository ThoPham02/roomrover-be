package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterRoomGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// FilterRoomGroup
func NewFilterRoomGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterRoomGroupLogic {
	return &FilterRoomGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterRoomGroupLogic) FilterRoomGroup(req *types.FilterRoomGroupReq) (resp *types.FilterRoomGroupRes, err error) {
	// todo: add your logic here and delete this line

	return
}
