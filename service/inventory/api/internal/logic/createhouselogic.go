package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHouseLogic {
	return &CreateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHouseLogic) CreateHouse(req *types.CreateHouseReq) (resp *types.CreateHouseRes, err error) {
	l.Logger.Info("CreateHouse: ", req)

	var userID int64
	var houseID int64 = l.svcCtx.ObjSync.GenServiceObjID()
	var currentTime = common.GetCurrentTime()

	var albums []string
	var services []types.Service
	var rooms []types.Room

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, err
	}

	if len(req.Albums) > 0 {
		err = json.Unmarshal([]byte(req.Albums), &albums)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, err
		}
	}

	if len(req.Rooms) > 0 {
		err = json.Unmarshal([]byte(req.Rooms), &rooms)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, err
		}
	}

	if len(req.Services) > 0 {
		err = json.Unmarshal([]byte(req.Services), &services)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, err
		}
	}

	_, err = l.svcCtx.HouseModel.Insert(l.ctx, &model.HouseTbl{
		Id:          houseID,
		UserId:      userID,
		Name:        sql.NullString{Valid: true, String: req.Name},
		Description: sql.NullString{Valid: true, String: req.Description},
		Type:        req.Type,
		Area:        req.Area,
		Price:       req.Price,
		Status:      common.HOUSE_STATUS_INACTIVE,
		BedNum:      sql.NullInt64{Valid: true, Int64: int64(req.BedNum)},
		LivingNum:   sql.NullInt64{Valid: true, Int64: int64(req.LivingNum)},
		Address:     sql.NullString{Valid: true, String: req.Address},
		WardId:      req.WardID,
		DistrictId:  req.DistrictID,
		ProvinceId:  req.ProvinceID,
		CreatedAt:   sql.NullInt64{Valid: true, Int64: currentTime},
		UpdatedAt:   sql.NullInt64{Valid: true, Int64: currentTime},
		CreatedBy:   sql.NullInt64{Valid: true, Int64: userID},
		UpdatedBy:   sql.NullInt64{Valid: true, Int64: userID},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}
	for _, album := range albums {
		_, err = l.svcCtx.AlbumModel.Insert(l.ctx, &model.AlbumTbl{
			Id:      l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId: sql.NullInt64{Valid: true, Int64: houseID},
			Url:     sql.NullString{Valid: true, String: album},
		})
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, err
		}
	}
	for _, room := range rooms {
		_, err = l.svcCtx.RoomModel.Insert(l.ctx, &model.RoomTbl{
			Id:       l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId:  sql.NullInt64{Valid: true, Int64: houseID},
			Name:     sql.NullString{Valid: true, String: room.Name},
			Status:   common.ROOM_STATUS_INACTIVE,
			Capacity: sql.NullInt64{Valid: true, Int64: room.Capacity},
			EIndex:   sql.NullInt64{Valid: true, Int64: 0},
			WIndex:   sql.NullInt64{Valid: true, Int64: 0},
		})
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, err
		}
	}
	for _, service := range services {
		_, err = l.svcCtx.ServiceModel.Insert(l.ctx, &model.ServiceTbl{
			Id:      l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId: sql.NullInt64{Valid: true, Int64: houseID},
			Name:    sql.NullString{Valid: true, String: service.Name},
			Price:   sql.NullInt64{Valid: true, Int64: service.Price},
			Unit:    sql.NullInt64{Valid: true, Int64: service.Unit},
		})
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, err
		}
	}

	l.Logger.Info("CreateHouse Success: ", userID)
	return &types.CreateHouseRes{
		Result: types.Result{
			Code: common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
