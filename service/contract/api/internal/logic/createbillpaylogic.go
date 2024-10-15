package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBillPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBillPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBillPayLogic {
	return &CreateBillPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBillPayLogic) CreateBillPay(req *types.CreateBillPayReq) (resp *types.CreateBillPayRes, err error) {
	// add your logic here
	l.Logger.Infof("CreateBillPay request: %v", req)

	return &types.CreateBillPayRes{}, nil
}
