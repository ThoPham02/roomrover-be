package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// DeleteHome
func NewDeleteHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeLogic {
	return &DeleteHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteHomeLogic) DeleteHome(req *types.DeleteHomeReq) (resp *types.DeleteHomeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
