package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBillStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBillStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBillStatusLogic {
	return &UpdateBillStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBillStatusLogic) UpdateBillStatus(req *types.UpdateBillStatusReq) (resp *types.UpdateBillStatusRes, err error) {
	// todo: add your logic here and delete this line

	return
}
