package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// FilterHome
func NewFilterHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterHomeLogic {
	return &FilterHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterHomeLogic) FilterHome(req *types.FilterHomeReq) (resp *types.FilterHomeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
