package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

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
	l.Logger.Info("DeleteContract", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	// Check if the contract exists
	contractModel, err := l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.DeleteContractRes{
				Result: types.Result{
					Code:    common.CONTRACT_NOT_FOUND_CODE,
					Message: common.CONTRACT_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.DeleteContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if contractModel.Status.Int64 != common.CONTRACT_STATUS_DRAF {
		return &types.DeleteContractRes{
			Result: types.Result{
				Code:    common.CONTRACT_HAS_BEEN_APPROVED_CODE,
				Message: common.CONTRACT_HAS_BEEN_APPROVED_MESS,
			},
		}, nil
	}

	err = l.svcCtx.PaymentModel.DeleteByContractID(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.ContractModel.Delete(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("DeleteContract Success: ", userID)
	return &types.DeleteContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
