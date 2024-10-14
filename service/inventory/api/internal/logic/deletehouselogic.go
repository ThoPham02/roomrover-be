package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHouseLogic {
	return &DeleteHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteHouseLogic) DeleteHouse(req *types.DeleteHouseReq) (resp *types.DeleteHouseRes, err error) {
	l.Logger.Info("DeleteHouse: ", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, err
	}

	house, err := l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}
	if house == nil {
		l.Logger.Error("House not found: ", req.HouseID)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, err
	}
	if house.UserId != userID {
		l.Logger.Error("User not authorized to delete house: ", userID)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, err
	}
	countContract, err := l.svcCtx.ContractFunction.CountContractByHouseID(req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}
	if countContract > 0 {
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.HOUSE_HAS_CONTRACT_ERR_CODE,
				Message: common.HOUSE_HAS_CONTRACT_ERR_MESS,
			},
		}, err
	}

	err = l.svcCtx.HouseModel.Delete(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	l.Logger.Info("DeleteHouse Success: ", userID)
	return &types.DeleteHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
