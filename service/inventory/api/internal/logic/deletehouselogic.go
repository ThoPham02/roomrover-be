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

// Delete House
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
		}, nil
	}

	err = l.svcCtx.AlbumModel.DeleteByHouseID(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.ServiceModel.DeleteByHouseID(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.RoomModel.DeleteByHouseID(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.HouseModel.Delete(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("DeleteHouse Success: ", userID)
	return &types.DeleteHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
