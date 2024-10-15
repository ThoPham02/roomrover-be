package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterBillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterBillLogic {
	return &FilterBillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterBillLogic) FilterBill(req *types.FilterBillReq) (resp *types.FilterBillRes, err error) {
	// todo: add your logic here and delete this line

	return
}
