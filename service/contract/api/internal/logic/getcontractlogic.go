package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get contract
func NewGetContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContractLogic {
	return &GetContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContractLogic) GetContract(req *types.GetContractReq) (resp *types.GetContractRes, err error) {
	// todo: add your logic here and delete this line

	return
}
