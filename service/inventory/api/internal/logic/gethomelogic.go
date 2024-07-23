package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetHome
func NewGetHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHomeLogic {
	return &GetHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHomeLogic) GetHome(req *types.GetHomeReq) (resp *types.GetHomeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
