package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHouseLogic {
	return &UpdateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHouseLogic) UpdateHouse(req *types.UpdateHouseReq) (resp *types.UpdateHouseRes, err error) {
	l.Logger.Info("UpdateHouse: ", req)

	var userID int64
	var imageUrls []string
	var currentTime int64 = time.Now().UnixMilli()
	var deleteRooms []int64
	var deleteServices []int64

	mapRoomExist := make(map[int64]bool)
	mapServiceExist := make(map[int64]bool)
	var services []types.Service
	var rooms []types.Room

	var houseModel *model.HouseTbl
	var albumModels []*model.AlbumTbl
	var updateRoomModels []*model.RoomTbl
	var createRoomModels []*model.RoomTbl

	var updateServiceModels []*model.ServiceTbl
	var createServiceModels []*model.ServiceTbl

	var houseStatus int64 = common.HOUSE_STATUS_DRAFT
	var roomStatus int64 = common.ROOM_STATUS_INACTIVE
	if req.Option == 1 {
		houseStatus = common.HOUSE_STATUS_ACTIVE
		roomStatus = common.ROOM_STATUS_ACTIVE
	}

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	if len(req.Albums) > 0 {
		err = json.Unmarshal([]byte(req.Albums), &imageUrls)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
	}

	if len(req.Rooms) > 0 {
		err = json.Unmarshal([]byte(req.Rooms), &rooms)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		for _, room := range rooms {
			if room.RoomID > common.MIN_ID {
				mapRoomExist[room.RoomID] = true

				roomModel, err := l.svcCtx.RoomModel.FindOne(l.ctx, room.RoomID)
				if err != nil || roomModel == nil {
					l.Logger.Error(err)
					return &types.UpdateHouseRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}

				roomModel.Capacity = sql.NullInt64{Int64: room.Capacity, Valid: true}
				roomModel.Name = sql.NullString{String: room.Name, Valid: true}
				roomModel.Status = roomStatus

				updateRoomModels = append(updateRoomModels, roomModel)
			} else {
				createRoomModels = append(createRoomModels, &model.RoomTbl{
					Id:       l.svcCtx.ObjSync.GenServiceObjID(),
					HouseId:  sql.NullInt64{Int64: req.HouseID, Valid: true},
					Name:     sql.NullString{Valid: true, String: room.Name},
					Status:   roomStatus,
					Capacity: sql.NullInt64{Valid: true, Int64: room.Capacity},
					EIndex:   sql.NullInt64{Valid: true, Int64: 0},
					WIndex:   sql.NullInt64{Valid: true, Int64: 0},
				})
			}
		}
	}
	oldRoomModels, _, err := l.svcCtx.RoomModel.FindByHouseID(l.ctx, req.HouseID, 0, 0)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, room := range oldRoomModels {
		if _, ok := mapRoomExist[room.Id]; !ok {
			deleteRooms = append(deleteRooms, room.Id)
		}
	}

	if len(req.Services) > 0 {
		err = json.Unmarshal([]byte(req.Services), &services)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}

		for _, service := range services {
			if service.ServiceID > common.MIN_ID {
				mapServiceExist[service.ServiceID] = true

				serviceModel, err := l.svcCtx.ServiceModel.FindOne(l.ctx, service.ServiceID)
				if err != nil || serviceModel == nil {
					l.Logger.Error(err)
					return &types.UpdateHouseRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}

				serviceModel.Name = sql.NullString{String: service.Name, Valid: true}
				serviceModel.Price = sql.NullInt64{Int64: service.Price, Valid: true}
				serviceModel.Unit = sql.NullInt64{Int64: service.Unit, Valid: true}

				updateServiceModels = append(updateServiceModels, serviceModel)
			} else {
				createServiceModels = append(createServiceModels, &model.ServiceTbl{
					Id:      l.svcCtx.ObjSync.GenServiceObjID(),
					HouseId: sql.NullInt64{Int64: req.HouseID, Valid: true},
					Name:    sql.NullString{Valid: true, String: service.Name},
					Price:   sql.NullInt64{Valid: true, Int64: service.Price},
					Unit:    sql.NullInt64{Valid: true, Int64: service.Unit},
				})
			}
		}
	}
	oldServiceModels, err := l.svcCtx.ServiceModel.FindByHouseID(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, service := range oldServiceModels {
		if _, ok := mapServiceExist[service.Id]; !ok {
			deleteServices = append(deleteServices, service.Id)
		}
	}

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if houseModel == nil {
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.HOUSE_NOT_FOUND_CODE,
				Message: common.HOUSE_NOT_FOUND_MESS,
			},
		}, nil
	}
	if houseModel.UserId != userID {
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.PERMISSION_DENIED_ERR_CODE,
				Message: common.PERMISSION_DENIED_ERR_MESS,
			},
		}, nil
	}

	for _, url := range imageUrls {
		albumModel := &model.AlbumTbl{
			Id:      l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId: sql.NullInt64{Int64: houseModel.Id, Valid: true},
			Url:     sql.NullString{String: url, Valid: true},
		}

		albumModels = append(albumModels, albumModel)
	}

	houseModel = &model.HouseTbl{
		Id:          houseModel.Id,
		UserId:      houseModel.UserId,
		Name:        sql.NullString{String: req.Name, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
		Type:        req.Type,
		Area:        req.Area,
		Price:       req.Price,
		Status:      houseStatus,
		BedNum:      sql.NullInt64{Valid: true, Int64: int64(req.BedNum)},
		LivingNum:   sql.NullInt64{Valid: true, Int64: int64(req.LivingNum)},
		Unit:        sql.NullInt64{Valid: true, Int64: int64(req.Unit)},
		Address:     sql.NullString{String: req.Address, Valid: true},
		WardId:      req.WardID,
		DistrictId:  req.DistrictID,
		ProvinceId:  req.ProvinceID,
		CreatedAt:   houseModel.CreatedAt,
		UpdatedAt:   sql.NullInt64{Int64: currentTime, Valid: true},
		CreatedBy:   houseModel.CreatedBy,
		UpdatedBy:   sql.NullInt64{Int64: userID, Valid: true},
	}

	err = l.svcCtx.AlbumModel.DeleteByHouseID(l.ctx, houseModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, roomID := range deleteRooms {
		err = l.svcCtx.RoomModel.Delete(l.ctx, roomID)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}
	for _, roomModel := range updateRoomModels {
		err = l.svcCtx.RoomModel.Update(l.ctx, roomModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}
	for _, roomModel := range createRoomModels {
		_, err = l.svcCtx.RoomModel.Insert(l.ctx, roomModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	for _, serviceID := range deleteServices {
		err = l.svcCtx.ServiceModel.Delete(l.ctx, serviceID)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}
	for _, serviceModel := range updateServiceModels {
		err = l.svcCtx.ServiceModel.Update(l.ctx, serviceModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}
	for _, serviceModel := range createServiceModels {
		_, err = l.svcCtx.ServiceModel.Insert(l.ctx, serviceModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	err = l.svcCtx.HouseModel.Update(l.ctx, houseModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, albumModel := range albumModels {
		_, err = l.svcCtx.AlbumModel.Insert(l.ctx, albumModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	l.Logger.Info("UpdateHouse Success:", userID)
	return &types.UpdateHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
