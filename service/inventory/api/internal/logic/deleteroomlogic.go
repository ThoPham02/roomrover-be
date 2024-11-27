package logic

import (
	"context"

	"roomrover/common"
	contractModel "roomrover/service/contract/model"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoomLogic {
	return &DeleteRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoomLogic) DeleteRoom(req *types.DeleteRoomReq) (resp *types.DeleteRoomRes, err error) {
	l.Logger.Info("DeleteRoom: ", req)

	var userID int64
	var roomModel *model.RoomTbl
	var houseModel *model.HouseTbl
	var contractModels []*contractModel.ContractTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	roomModel, err = l.svcCtx.RoomModel.FindOne(l.ctx, req.ID)
	if err != nil || roomModel == nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	contractModels, err = l.svcCtx.ContractFunction.GetContractByRoom(req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if len(contractModels) > 0 {
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.ROOM_HAS_CONTRACT_CODE,
				Message: common.ROOM_HAS_CONTRACT_MESS,
			},
		}, nil
	}

	err = l.svcCtx.RoomModel.Delete(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, roomModel.HouseId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	totalRoom, err := l.svcCtx.RoomModel.CountRoomActiveByHouseID(l.ctx, roomModel.HouseId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if totalRoom == 0 && houseModel.Status == common.HOUSE_STATUS_ACTIVE {
		houseModel.Status = common.HOUSE_STATUS_SOLD_OUT
		
		err = l.svcCtx.HouseModel.Update(l.ctx, houseModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.DeleteRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	l.Logger.Info("DeleteRoom Success: ", userID)
	return &types.DeleteRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
