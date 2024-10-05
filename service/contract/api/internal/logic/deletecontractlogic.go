package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete contract
func NewDeleteContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContractLogic {
	return &DeleteContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContractLogic) DeleteContract(req *types.DeleteContractReq) (resp *types.DeleteContractRes, err error) {
	// todo: add your logic here and delete this line

	return &types.DeleteContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
