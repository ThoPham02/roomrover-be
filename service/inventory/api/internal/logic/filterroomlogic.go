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

	var userID int64
	var total int
	var mapHouse = map[int64]*model.HouseTbl{}

	var rooms []types.Room

	var roomModels []*model.RoomTbl

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

	total, err = l.svcCtx.RoomModel.CountRoom(l.ctx, req.Search, req.Type)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if total == 0 {
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
			Total: total,
			Rooms: []types.Room{},
		}, nil
	}
	roomModels, err = l.svcCtx.RoomModel.FilterRoom(l.ctx, req.Search, req.Type, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, roomModel := range roomModels {
		var houseModel *model.HouseTbl
		houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, roomModel.HouseId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.FilterRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		mapHouse[roomModel.HouseId.Int64] = houseModel
	}

	for _, roomModel := range roomModels {
		var houseModel = model.HouseTbl{}
		if mapHouse[roomModel.HouseId.Int64] != nil {
			houseModel = *mapHouse[roomModel.HouseId.Int64]
		}

		rooms = append(
			rooms,
			types.Room{
				RoomID:    roomModel.Id,
				HouseID:   roomModel.HouseId.Int64,
				HouseName: houseModel.Name.String,
				Area:      houseModel.Area,
				Price:     houseModel.Price,
				Type:      houseModel.Type,
				Name:      roomModel.Name.String,
				Status:    roomModel.Status,
				Capacity:  roomModel.Capacity.Int64,
				EIndex:    roomModel.EIndex.Int64,
				WIndex:    roomModel.WIndex.Int64,
			},
		)
	}

	l.Logger.Info("FilterRoom Success: ", userID)
	return &types.FilterRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total: total,
		Rooms: rooms,
	}, nil
}
