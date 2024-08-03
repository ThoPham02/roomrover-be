package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// FilterHouse
func NewFilterHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterHouseLogic {
	return &FilterHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterHouseLogic) FilterHouse(req *types.FilterHouseReq) (resp *types.FilterHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
