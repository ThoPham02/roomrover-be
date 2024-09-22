package logic

import (
	"context"
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

	var albums []types.Album
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
		Name:        req.Name,
		Description: req.Description,
		// Util:        req.Util,

		Type:   req.Type,
		Area:   req.Area,
		Price:  req.Price,
		Status: common.HOUSE_STATUS_DRAFT,

		Address:    req.Address,
		WardId:     req.WardID,
		DistrictId: req.DistrictID,
		ProvinceId: req.ProvinceID,

		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	for _, url := range imageUrls {
		var typeAlbum = common.ALBUM_TYPE_OTHER
		albumModel := &model.AlbumTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			HouseId:   houseModel.Id,
			Url:       url,
			Type:      int64(typeAlbum),
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
			CreatedBy: userID,
			UpdatedBy: userID,
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
		UserID:      houseModel.UserId,
		Name:        houseModel.Name,
		Description: houseModel.Description,
		Type:        houseModel.Type,
		Area:        houseModel.Area,
		Price:       houseModel.Price,
		Status:      houseModel.Status,
		Address:     houseModel.Address,
		WardID:      houseModel.WardId,
		DistrictID:  houseModel.DistrictId,
		ProvinceID:  houseModel.ProvinceId,
		CreatedAt:   houseModel.CreatedAt,
		UpdatedAt:   houseModel.UpdatedAt,
		CreatedBy:   houseModel.CreatedBy,
		UpdatedBy:   houseModel.UpdatedBy,
		Albums:      albums,
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
