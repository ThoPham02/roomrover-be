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

	var houseModel *model.HouseTbl
	var albumModels []*model.AlbumTbl

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
		Status:      houseModel.Status,
		BedNum:      sql.NullInt64{},
		LivingNum:   sql.NullInt64{},
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
