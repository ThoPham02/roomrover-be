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

// Get house
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

	var houseModel *model.HouseTbl
	var albumModels []*model.AlbumTbl
	var roomModels []*model.RoomTbl
	var serviceModels []*model.ServiceTbl

	var house types.House
	var albums []types.Album
	var rooms []types.Room
	var services []types.Service

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

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.GetHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	albumModels, err = l.svcCtx.AlbumModel.FindByHouseID(l.ctx, req.HouseID)
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
		albums = append(albums, types.Album{
			AlbumID: albumModel.Id,
			HouseID: albumModel.HouseId,
			Url:     albumModel.Url,
			Type:    albumModel.Type,
		})
	}

	roomModels, err = l.svcCtx.RoomModel.FindByHouseID(l.ctx, req.HouseID)
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
		rooms = append(rooms, types.Room{
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

	serviceModels, err = l.svcCtx.ServiceModel.FindByHouseID(l.ctx, req.HouseID)
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
		services = append(services, types.Service{
			ServiceID: serviceModel.Id,
			HouseID:   serviceModel.HouseId,
			Name:      serviceModel.Name,
			Price:     serviceModel.Price,
			CreatedAt: serviceModel.CreatedAt,
			UpdatedAt: serviceModel.UpdatedAt,
			CreatedBy: serviceModel.CreatedBy,
			UpdatedBy: serviceModel.UpdatedBy,
		})
	}

	house = types.House{
		HouseID:     houseModel.Id,
		UserID:      userID,
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
		Rooms:       rooms,
		Services:    services,
	}

	l.Logger.Info("GetHouse Success: ", userID)
	return &types.GetHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
