package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterContractLogic {
	return &FilterContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterContractLogic) FilterContract(req *types.FilterContractReq) (resp *types.FilterContractRes, err error) {
	// todo: add your logic here and delete this line

	return
}
