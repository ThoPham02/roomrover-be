package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHouseLogic {
	return &GetHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHouseLogic) GetHouse(req *types.GetHouseReq) (resp *types.GetHouseRes, err error) {
	l.Logger.Info("GetHouse", req)

	var userID int64

	var house types.House
	var album []types.Album
	var room []types.Room
	var service []types.Service

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == model.ErrNotFound {
			l.Logger.Error(err)
			return &types.GetHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	albumModels, err := l.svcCtx.AlbumModel.FindByHouseID(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, albumModel := range albumModels {
		album = append(album, types.Album{
			AlbumID:   albumModel.Id,
			HouseID:   albumModel.HouseId,
			Url:       albumModel.Url,
			Type:      albumModel.Type,
			CreatedAt: albumModel.CreatedAt,
			UpdatedAt: albumModel.UpdatedAt,
			CreatedBy: albumModel.CreatedBy,
			UpdatedBy: albumModel.UpdatedBy,
		})
	}

	roomModels, err := l.svcCtx.RoomModel.FindByHouseID(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, roomModel := range roomModels {
		room = append(room, types.Room{
			RoomID:    roomModel.Id,
			HouseID:   roomModel.HouseId,
			Name:      roomModel.Name,
			Status:    roomModel.Status,
			CreatedAt: roomModel.CreatedAt,
			UpdatedAt: roomModel.UpdatedAt,
			CreatedBy: roomModel.CreatedBy,
			UpdatedBy: roomModel.UpdatedBy,
		})
	}

	serviceModels, err := l.svcCtx.ServiceModel.FindByHouseID(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, serviceModel := range serviceModels {
		service = append(service, types.Service{
			ServiceID: serviceModel.Id,
			HouseID:   serviceModel.HouseId,
			Name:      serviceModel.Name,
			Price:     serviceModel.Price,
			Type:      serviceModel.Type,
			CreatedAt: serviceModel.CreatedAt,
			UpdatedAt: serviceModel.UpdatedAt,
			CreatedBy: serviceModel.CreatedBy,
			UpdatedBy: serviceModel.UpdatedBy,
		})
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
		Albums:      album,
		Rooms:       room,
		Services:    service,
	}

	l.Logger.Info("GetHouse Success", userID)
	return &types.GetHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
