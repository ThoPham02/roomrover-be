package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRoomLogic {
	return &SearchRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchRoomLogic) SearchRoom(req *types.SearchRoomReq) (resp *types.SearchRoomRes, err error) {
	l.Logger.Info("SearchRoom: ", req)

	var userID int64
	var total int
	var rooms []types.Room
	var mapHouseService = make(map[int64][]types.Service)
	var houseIDs []int64
	var houseModels []*model.HouseTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	total, err = l.svcCtx.RoomModel.CountSearchRoom(l.ctx, userID, req.Type, req.Search, req.Status)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if total == 0 {
		return &types.SearchRoomRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
			Rooms: rooms,
			Total: total,
		}, nil
	}

	houseRoomModels, err := l.svcCtx.RoomModel.SearchRoom(l.ctx, userID, req.Type, req.Search, req.Status, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, room := range houseRoomModels {
		houseIDs = append(houseIDs, room.HouseId.Int64)
	}

	houseModels, err = l.svcCtx.HouseModel.FindMultiByID(l.ctx, houseIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, house := range houseModels {
		services, err := l.svcCtx.ServiceModel.FindByHouseID(l.ctx, house.Id)
		if err != nil {
			l.Logger.Error(err)
			return &types.SearchRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		for _, service := range services {
			mapHouseService[house.Id] = append(mapHouseService[house.Id], types.Service{
				ServiceID: service.Id,
				HouseID:   service.HouseId.Int64,
				Name:      service.Name.String,
				Price:     service.Price.Int64,
				Unit:      service.Unit.Int64,
			})
		}
	}

	for _, room := range houseRoomModels {
		rooms = append(rooms, types.Room{
			RoomID:     room.Id,
			HouseID:    room.HouseId.Int64,
			Name:       room.HouseRoomName.String,
			ProvinceID: room.ProvinceID.Int64,
			DistrictID: room.DistrictID.Int64,
			WardID:     room.WardID.Int64,
			Address:    room.Address.String,
			Area:       room.Area.Int64,
			Price:      room.Price.Int64,
			Status:     room.Status,
			Capacity:   room.Capacity.Int64,
			EIndex:     room.EIndex.Int64,
			WIndex:     room.WIndex.Int64,
			Services:   mapHouseService[room.HouseId.Int64],
		})
	}

	l.Logger.Info("SearchRoom Success: ", userID)
	return &types.SearchRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Rooms: rooms,
		Total: total,
	}, nil
}
