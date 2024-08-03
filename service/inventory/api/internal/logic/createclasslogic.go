package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// CreateClass
func NewCreateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClassLogic {
	return &CreateClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClassLogic) CreateClass(req *types.CreateClassReq) (resp *types.CreateClassRes, err error) {
	// todo: add your logic here and delete this line

	return
}
