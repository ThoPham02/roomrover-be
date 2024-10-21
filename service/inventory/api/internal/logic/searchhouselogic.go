package logic

import (
	"context"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHouseLogic {
	return &SearchHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchHouseLogic) SearchHouse(req *types.SearchHouseReq) (resp *types.SearchHouseRes, err error) {
	l.Logger.Info("SearchHouse: ", req)

	var total int64
	var houseIDs []int64
	var userIDs []int64

	var listHouses []types.House
	var mapUser = map[int64]types.User{}
	var mapService = map[int64][]types.Service{}
	var mapRoom = map[int64][]types.Room{}
	var mapAlbum = map[int64][]string{}

	var listHousesModel []*model.HouseTbl
	var roomModels []*model.RoomTbl
	var serviceModels []*model.ServiceTbl
	var albumModels []*model.AlbumTbl
	var userModels []*accountModel.UserTbl

	total, listHousesModel, err = l.svcCtx.HouseModel.SearchHouse(l.ctx, req.Search, req.DistrictID, req.ProvinceID, req.WardID, req.PriceFrom, req.PriceTo, req.AreaFrom, req.AreaTo, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, houseModel := range listHousesModel {
		houseIDs = append(houseIDs, houseModel.Id)
		userIDs = append(userIDs, houseModel.UserId)
	}

	roomModels, err = l.svcCtx.RoomModel.FindMultiByHouseIDs(l.ctx, houseIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, roomModel := range roomModels {
		mapRoom[roomModel.HouseId.Int64] = append(mapRoom[roomModel.HouseId.Int64], types.Room{
			RoomID:   roomModel.Id,
			Name:     roomModel.Name.String,
			Status:   roomModel.Status,
			Capacity: roomModel.Capacity.Int64,
			EIndex:   roomModel.EIndex.Int64,
			WIndex:   roomModel.WIndex.Int64,
		})
	}

	serviceModels, err = l.svcCtx.ServiceModel.FindMultiByHouseIDs(l.ctx, houseIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, serviceModel := range serviceModels {
		mapService[serviceModel.HouseId.Int64] = append(mapService[serviceModel.HouseId.Int64], types.Service{
			ServiceID: serviceModel.Id,
			HouseID:   serviceModel.HouseId.Int64,
			Name:      serviceModel.Name.String,
			Price:     serviceModel.Price.Int64,
			Unit:      serviceModel.Unit.Int64,
		})
	}

	albumModels, err = l.svcCtx.AlbumModel.FindMultiByHouseIDs(l.ctx, houseIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, albumModel := range albumModels {
		mapAlbum[albumModel.HouseId.Int64] = append(mapAlbum[albumModel.HouseId.Int64], albumModel.Url.String)
	}

	userModels, err = l.svcCtx.AccountFunction.GetUsersByIDs(userIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.SearchHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, userModel := range userModels {
		mapUser[userModel.Id] = types.User{
			UserID:    userModel.Id,
			Phone:     userModel.Phone,
			Role:      userModel.Role.Int64,
			Status:    userModel.Status,
			Address:   userModel.Address.String,
			FullName:  userModel.FullName.String,
			AvatarUrl: userModel.AvatarUrl.String,
			Birthday:  userModel.Birthday.Int64,
			Gender:    userModel.Gender.Int64,
		}
	}

	for _, houseModel := range listHousesModel {
		listHouses = append(listHouses, types.House{
			HouseID:     houseModel.Id,
			User:        mapUser[houseModel.UserId],
			Name:        houseModel.Name.String,
			Description: houseModel.Description.String,
			Type:        houseModel.Type,
			Status:      houseModel.Status,
			Area:        houseModel.Area,
			Price:       houseModel.Price,
			BedNum:      houseModel.BedNum.Int64,
			LivingNum:   houseModel.LivingNum.Int64,
			Albums:      mapAlbum[houseModel.Id],
			Rooms:       mapRoom[houseModel.Id],
			Services:    mapService[houseModel.Id],
			Address:     houseModel.Address.String,
			WardID:      houseModel.WardId,
			DistrictID:  houseModel.DistrictId,
			ProvinceID:  houseModel.ProvinceId,
			CreatedAt:   houseModel.CreatedAt.Int64,
			UpdatedAt:   houseModel.UpdatedAt.Int64,
			CreatedBy:   houseModel.CreatedBy.Int64,
			UpdatedBy:   houseModel.UpdatedBy.Int64,
		})
	}
	l.Logger.Info("SearchHouse Success: ", resp)
	return &types.SearchHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:  int(total),
		Houses: listHouses,
	}, nil
}
