package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterRoomLogic {
	return &FilterRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterRoomLogic) FilterRoom(req *types.FilterRoomReq) (resp *types.FilterRoomRes, err error) {
	l.Logger.Info("FilterRoom: ", req)
	var count int
	var rooms []types.Room
	var roomModels []*model.RoomTbl

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	userModel, err := l.svcCtx.AccountFunction.GetUserByID(userID)
	if err != nil || userModel == nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	if userModel.Role.Int64 != common.USER_ROLE_LESSOR {
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.PERMISSION_DENIED_ERR_CODE,
				Message: common.PERMISSION_DENIED_ERR_MESS,
			},
		}, nil
	}

	count, err = l.svcCtx.RoomModel.CountRoom(l.ctx, userID, req.Search, req.Type, req.Status)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if count == 0 {
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	roomModels, err = l.svcCtx.RoomModel.FilterRoom(l.ctx, userID, req.Search, req.Type, req.Status, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, room := range roomModels {
		houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, room.HouseId.Int64)
		if err != nil || houseModel == nil {
			l.Logger.Error(err)
			return &types.FilterRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		rooms = append(rooms, types.Room{
			RoomID:    room.Id,
			HouseID:   room.HouseId.Int64,
			Name:      room.Name.String,
			HouseName: houseModel.Name.String,
			Area:      houseModel.Area,
			Price:     houseModel.Price,
			Type:      houseModel.Type,
			Status:    room.Status,
			Capacity:  room.Capacity.Int64,
			EIndex:    room.EIndex.Int64,
			WIndex:    room.WIndex.Int64,
		})
	}

	l.Logger.Info("FilterRoom Success: ", req)
	return &types.FilterRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total: count,
		Rooms: rooms,
	}, nil
}
