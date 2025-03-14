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

	var house types.House
	var imageUrls []string
	var room []types.Room
	var service []types.Service
	var user types.User

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
		imageUrls = append(imageUrls, albumModel.Url.String)
	}

	roomModels, _, err := l.svcCtx.RoomModel.FindByHouseID(l.ctx, req.ID, 0, 0)
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
			RoomID:   roomModel.Id,
			HouseID:  roomModel.HouseId.Int64,
			Name:     roomModel.Name.String,
			Status:   roomModel.Status,
			Capacity: roomModel.Capacity.Int64,
			EIndex:   roomModel.EIndex.Int64,
			WIndex:   roomModel.WIndex.Int64,
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
			HouseID:   serviceModel.HouseId.Int64,
			Name:      serviceModel.Name.String,
			Price:     serviceModel.Price.Int64,
			Unit:      serviceModel.Unit.Int64,
		})
	}

	userModel, err := l.svcCtx.AccountFunction.GetUserByID(houseModel.UserId)
	if err != nil || userModel == nil {
		l.Logger.Error(err)
		return &types.GetHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:      userModel.Id,
		Phone:       userModel.Phone,
		Role:        userModel.Role.Int64,
		Status:      userModel.Status,
		Address:     userModel.Address.String,
		FullName:    userModel.FullName.String,
		AvatarUrl:   userModel.AvatarUrl.String,
		Birthday:    userModel.Birthday.Int64,
		Gender:      userModel.Gender.Int64,
		CccdNumber:  userModel.CCCDNumber.String,
		CccdDate:    userModel.CCCDDate.Int64,
		CccdAddress: userModel.CCCDAddress.String,
		CreatedAt:   userModel.CreatedAt.Int64,
		UpdatedAt:   userModel.UpdatedAt.Int64,
	}

	house = types.House{
		HouseID:     houseModel.Id,
		User:        user,
		Name:        houseModel.Name.String,
		Description: houseModel.Description.String,
		Type:        houseModel.Type,
		Status:      houseModel.Status,
		Area:        houseModel.Area,
		Price:       houseModel.Price,
		BedNum:      houseModel.BedNum.Int64,
		LivingNum:   houseModel.LivingNum.Int64,
		Unit:        houseModel.Unit.Int64,
		Albums:      imageUrls,
		Rooms:       room,
		Services:    service,
		Address:     houseModel.Address.String,
		WardID:      houseModel.WardId,
		DistrictID:  houseModel.DistrictId,
		ProvinceID:  houseModel.ProvinceId,
		CreatedAt:   houseModel.CreatedAt.Int64,
		UpdatedAt:   houseModel.UpdatedAt.Int64,
		CreatedBy:   houseModel.CreatedBy.Int64,
		UpdatedBy:   houseModel.UpdatedBy.Int64,
	}

	l.Logger.Info("GetHouse Success")
	return &types.GetHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
