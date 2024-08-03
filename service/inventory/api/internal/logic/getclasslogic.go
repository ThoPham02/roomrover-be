package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetClass
func NewGetClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassLogic {
	return &GetClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassLogic) GetClass(req *types.GetClassReq) (resp *types.GetClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
