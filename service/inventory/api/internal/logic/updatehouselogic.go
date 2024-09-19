package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHouseLogic {
	return &UpdateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHouseLogic) UpdateHouse(req *types.UpdateHouseReq) (resp *types.UpdateHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
