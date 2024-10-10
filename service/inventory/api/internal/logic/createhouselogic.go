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

// Create house
func NewCreateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHouseLogic {
	return &CreateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHouseLogic) CreateHouse(req *types.CreateHouseReq) (resp *types.CreateHouseRes, err error) {
	l.Logger.Infof("CreateHouse: %v", req)

	var userID int64
	var currentTime = common.GetCurrentTime()
	var imageUrls []string

	var house types.House

	var houseModel *model.HouseTbl
	var albumModels []*model.AlbumTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	if len(req.Albums) > 0 {
		l.Logger.Info(req.Albums)

		err = json.Unmarshal([]byte(req.Albums), &imageUrls)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
	}

	houseModel = &model.HouseTbl{
		Id:          l.svcCtx.ObjSync.GenServiceObjID(),
		UserId:      userID,
		Name:        sql.NullString{String: req.Name, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},

		Type:   req.Type,
		Area:   req.Area,
		Price:  req.Price,
		Status: common.HOUSE_STATUS_DRAFT,

		Address:    sql.NullString{String: req.Address, Valid: true},
		WardId:     req.WardID,
		DistrictId: req.DistrictID,
		ProvinceId: req.ProvinceID,

		CreatedAt: sql.NullInt64{Int64: currentTime, Valid: true},
		UpdatedAt: sql.NullInt64{Int64: currentTime, Valid: true},
		CreatedBy: sql.NullInt64{Int64: userID, Valid: true},
		UpdatedBy: sql.NullInt64{Int64: userID, Valid: true},
	}

	for _, url := range imageUrls {
		albumModel := &model.AlbumTbl{
			Id:      l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId: sql.NullInt64{Int64: houseModel.Id, Valid: true},
			Url:     sql.NullString{String: url, Valid: true},
		}

		albumModels = append(albumModels, albumModel)
	}

	_, err = l.svcCtx.HouseModel.Insert(l.ctx, houseModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateHouseRes{
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
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	house = types.House{
		HouseID:     houseModel.Id,
		Name:        houseModel.Name.String,
		Description: houseModel.Description.String,
		Type:        houseModel.Type,
		Status:      houseModel.Status,
		Area:        houseModel.Area,
		Price:       houseModel.Price,
		BedNum:      houseModel.BedNum.Int64,
		LivingNum:   houseModel.LivingNum.Int64,
		Albums:      imageUrls,
		Rooms:       []types.Room{},
		Services:    []types.Service{},
		Address:     houseModel.Address.String,
		WardID:      houseModel.WardId,
		DistrictID:  houseModel.DistrictId,
		ProvinceID:  houseModel.ProvinceId,
		CreatedAt:   houseModel.CreatedAt.Int64,
		UpdatedAt:   houseModel.UpdatedAt.Int64,
		CreatedBy:   houseModel.CreatedBy.Int64,
		UpdatedBy:   houseModel.UpdatedBy.Int64,
	}

	l.Logger.Info("CreateHouse success: ", userID)
	return &types.CreateHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
