package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterRoomClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// FilterRoomClass
func NewFilterRoomClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterRoomClassLogic {
	return &FilterRoomClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterRoomClassLogic) FilterRoomClass(req *types.FilterRoomClassReq) (resp *types.FilterRoomClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
