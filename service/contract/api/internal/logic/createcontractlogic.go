package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create contract
func NewCreateContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContractLogic {
	return &CreateContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContractLogic) CreateContract(req *types.CreateContractReq) (resp *types.CreateContractRes, err error) {
	// todo: add your logic here and delete this line

	return
}
