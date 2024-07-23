package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// CreateHome
func NewCreateHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomeLogic {
	return &CreateHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHomeLogic) CreateHome(req *types.CreateHomeReq) (resp *types.CreateHomeRes, err error) {
	


	return
}
