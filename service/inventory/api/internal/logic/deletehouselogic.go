package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete house
func NewDeleteHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHouseLogic {
	return &DeleteHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteHouseLogic) DeleteHouse(req *types.DeleteHouseReq) (resp *types.DeleteHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
