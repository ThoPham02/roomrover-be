package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// DeleteClass
func NewDeleteClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassLogic {
	return &DeleteClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassLogic) DeleteClass(req *types.DeleteClassReq) (resp *types.DeleteClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
