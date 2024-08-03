package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UpdateClass
func NewUpdateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassLogic {
	return &UpdateClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClassLogic) UpdateClass(req *types.UpdateClassReq) (resp *types.UpdateClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
