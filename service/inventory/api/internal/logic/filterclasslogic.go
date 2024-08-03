package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// FilterClass
func NewFilterClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterClassLogic {
	return &FilterClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterClassLogic) FilterClass(req *types.FilterClassReq) (resp *types.FilterClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
