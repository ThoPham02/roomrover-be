package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UpdateHome
func NewUpdateHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeLogic {
	return &UpdateHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHomeLogic) UpdateHome(req *types.UpdateHomeReq) (resp *types.UpdateHomeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
