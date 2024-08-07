package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update contract
func NewUpdateContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractLogic {
	return &UpdateContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractLogic) UpdateContract(req *types.UpdateContractReq) (resp *types.UpdateContractRes, err error) {
	// todo: add your logic here and delete this line

	return
}
