package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomsByHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetRoomsByHouse
func NewGetRoomsByHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomsByHouseLogic {
	return &GetRoomsByHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomsByHouseLogic) GetRoomsByHouse(req *types.GetRoomsByHouseReq) (resp *types.FilterRoomRes, err error) {
	// todo: add your logic here and delete this line

	return
}
