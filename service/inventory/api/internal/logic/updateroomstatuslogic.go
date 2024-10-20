package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoomStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoomStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoomStatusLogic {
	return &UpdateRoomStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoomStatusLogic) UpdateRoomStatus(req *types.UpdateRoomStatusReq) (resp *types.UpdateRoomStatusRes, err error) {
	l.Logger.Info("UpdateRoomStatus: ", req)

	var userID int64
	var countRenterHouse int64
	var countInActiveHouse int64
	var roomModel *model.RoomTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRoomStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, err
	}

	roomModel, err = l.svcCtx.RoomModel.FindOne(l.ctx, req.RoomID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UpdateRoomStatusRes{
				Result: types.Result{
					Code:    common.ROOM_NOT_FOUND_CODE,
					Message: common.ROOM_NOT_FOUND_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateRoomStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	if roomModel.Status == req.Status {
		return &types.UpdateRoomStatusRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	roomModel.Status = req.Status
	err = l.svcCtx.RoomModel.Update(l.ctx, roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRoomStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}
	houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, roomModel.HouseId.Int64)
	if err != nil || houseModel == nil {
		l.Logger.Error(err)
		return &types.UpdateRoomStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	if houseModel.Status != common.HOUSE_STATUS_INACTIVE {
		roomModels, count, err := l.svcCtx.RoomModel.FindByHouseID(l.ctx, roomModel.HouseId.Int64, 0, 0)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateRoomStatusRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, err
		}

		for _, room := range roomModels {
			if room.Status == common.ROOM_STATUS_INACTIVE {
				countInActiveHouse++
			} else if room.Status != common.ROOM_STATUS_ACTIVE {
				countRenterHouse++
			}
		}

		var countActiveHouse int64 = int64(count) - countRenterHouse - countInActiveHouse
		if countActiveHouse == 0 && countRenterHouse > 0 {
			houseModel.Status = common.HOUSE_STATUS_SOLD_OUT
		} else if countActiveHouse == 0 && countRenterHouse == 0 {
			houseModel.Status = common.HOUSE_STATUS_DRAFT
		} else {
			houseModel.Status = common.HOUSE_STATUS_ACTIVE
		}

		err = l.svcCtx.HouseModel.Update(l.ctx, houseModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateRoomStatusRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, err
		}
	}

	l.Logger.Info("UpdateRoomStatus success: ", userID)
	return &types.UpdateRoomStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
