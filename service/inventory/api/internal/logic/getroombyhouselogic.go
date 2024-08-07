package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomByHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get room by house
func NewGetRoomByHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomByHouseLogic {
	return &GetRoomByHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomByHouseLogic) GetRoomByHouse(req *types.GetRoomByHouseReq) (resp *types.GetRoomByHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
