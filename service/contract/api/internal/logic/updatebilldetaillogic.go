package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBillDetailLogic {
	return &UpdateBillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBillDetailLogic) UpdateBillDetail(req *types.UpdateBillDetailReq) (resp *types.UpdateBillDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
