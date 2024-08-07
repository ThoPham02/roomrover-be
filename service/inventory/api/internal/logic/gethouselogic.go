package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get house
func NewGetHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHouseLogic {
	return &GetHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHouseLogic) GetHouse(req *types.GetHouseReq) (resp *types.GetHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
