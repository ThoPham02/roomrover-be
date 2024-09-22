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
		listHouses = append(listHouses, types.House{
			HouseID:     house.Id,
			UserID:      userID,
			Name:        house.Name,
			Description: house.Description,
			Type:        house.Type,
			Area:        house.Area,
			Price:       house.Price,
			Status:      house.Status,
			Address:     house.Address,
			WardID:      house.WardId,
			DistrictID:  house.DistrictId,
			ProvinceID:  house.ProvinceId,
			CreatedAt:   house.CreatedAt,
			UpdatedAt:   house.UpdatedAt,
			CreatedBy:   house.CreatedBy,
			UpdatedBy:   house.UpdatedBy,
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
