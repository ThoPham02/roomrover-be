package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBillDetailLogic {
	return &GetBillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBillDetailLogic) GetBillDetail(req *types.GetBillDetailReq) (resp *types.GetBillDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
