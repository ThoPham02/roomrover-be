package logic

import (
	"context"
	"database/sql"

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

// UpdateHouse
func NewUpdateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHouseLogic {
	return &UpdateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHouseLogic) UpdateHouse(req *types.UpdateHouseReq) (resp *types.UpdateHouseRes, err error) {
	l.Logger.Info("UpdateHouse", req)

	var userID int64
	var house types.House
	var currentTime = common.GetCurrentTime()

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

	_, err = l.svcCtx.ProvincesModel.FindOne(l.ctx, int64(req.ProvinceID))
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.DistrictsModel.FindOne(l.ctx, int64(req.DistrictID))
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.WardsModel.FindOne(l.ctx, int64(req.WardID))
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.UpdateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	houseModel = &model.Houses{
		Id:              houseModel.Id,
		Name:            req.HouseName,
		Type:            int64(req.Type),
		SpecificAddress: req.Address,
		WardId:          sql.NullInt64{},
		DistrictId:      sql.NullInt64{},
		ProvinceId:      sql.NullInt64{},
		CreatedAt:       houseModel.CreatedAt,
		UpdatedAt:       currentTime,
		CreatedBy:       houseModel.CreatedBy,
		UpdatedBy:       userID,
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
	house = types.House{
		HouseID:    houseModel.Id,
		HouseName:  houseModel.Name,
		Type:       int(houseModel.Type),
		ProvinceID: int(houseModel.ProvinceId.Int64),
		DistrictID: int(houseModel.DistrictId.Int64),
		WardID:     int(houseModel.WardId.Int64),
		Address:    houseModel.SpecificAddress,
		CreatedAt:  houseModel.CreatedAt,
		UpdatedAt:  houseModel.UpdatedAt,
	}

	l.Logger.Info("UpdateHouse success")
	return &types.UpdateHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
