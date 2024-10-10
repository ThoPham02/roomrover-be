package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Filter house
func NewFilterHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterHouseLogic {
	return &FilterHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterHouseLogic) FilterHouse(req *types.FilterHouseReq) (resp *types.FilterHouseRes, err error) {
	l.Logger.Info("FilterHouse: ", req)

	var userID int64
	var total int64
	var imageUrls []string
	var listHouses []types.House
	var houses []*model.HouseTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	total, houses, err = l.svcCtx.HouseModel.FilterHouse(l.ctx, userID, req.Search, req.Limit, req.Offset)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.FilterHouseRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.FilterHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if total == 0 {
		return &types.FilterHouseRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	for _, house := range houses {
		var albumModels []*model.AlbumTbl
		albumModels, err = l.svcCtx.AlbumModel.FindByHouseID(l.ctx, house.Id)
		if err != nil {
			l.Logger.Error(err)
			return &types.FilterHouseRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		for _, album := range albumModels {
			imageUrls = append(imageUrls, album.Url.String)
		}

		listHouses = append(listHouses, types.House{
			HouseID:     house.Id,
			Name:        house.Name.String,
			Description: house.Description.String,
			Type:        house.Type,
			Status:      house.Status,
			Area:        house.Area,
			Price:       house.Price,
			BedNum:      house.BedNum.Int64,
			LivingNum:   house.LivingNum.Int64,
			Albums:      imageUrls,
			Rooms:       []types.Room{},
			Services:    []types.Service{},
			Address:     house.Address.String,
			WardID:      house.WardId,
			DistrictID:  house.DistrictId,
			ProvinceID:  house.ProvinceId,
			CreatedAt:   house.CreatedAt.Int64,
			UpdatedAt:   house.UpdatedAt.Int64,
			CreatedBy:   house.CreatedBy.Int64,
			UpdatedBy:   house.UpdatedBy.Int64,
		})
	}

	l.Logger.Info("FilterHouse Success: ", userID)
	return &types.FilterHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:      total,
		ListHouses: listHouses,
	}, nil
}
