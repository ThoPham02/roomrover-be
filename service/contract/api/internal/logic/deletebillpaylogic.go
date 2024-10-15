package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBillPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBillPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBillPayLogic {
	return &DeleteBillPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBillPayLogic) DeleteBillPay(req *types.DeleteBillPayReq) (resp *types.DeleteBillPayRes, err error) {
	// todo: add your logic here and delete this line

	return
}
