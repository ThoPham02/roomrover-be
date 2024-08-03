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

type CreateHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// CreateHouse
func NewCreateHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHouseLogic {
	return &CreateHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHouseLogic) CreateHouse(req *types.CreateHouseReq) (resp *types.CreateHouseRes, err error) {
	l.Logger.Info("CreateHouse", req)

	var userID int64
	var currentTime = common.GetCurrentTime()
	var house types.House

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

	_, err = l.svcCtx.ProvincesModel.FindOne(l.ctx, int64(req.ProvinceID))
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.CreateHouseRes{
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
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.CreateHouseRes{
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
			return &types.CreateHouseRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.CreateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.HouseModel.Insert(l.ctx, &model.Houses{
		Id:              l.svcCtx.ObjSync.GenServiceObjID(),
		Name:            req.HouseName,
		Type:            int64(req.Type),
		SpecificAddress: req.Address,
		WardId:          sql.NullInt64{Valid: true, Int64: int64(req.WardID)},
		DistrictId:      sql.NullInt64{Valid: true, Int64: int64(req.DistrictID)},
		ProvinceId:      sql.NullInt64{Valid: true, Int64: int64(req.ProvinceID)},
		CreatedAt:       currentTime,
		UpdatedAt:       currentTime,
		CreatedBy:       userID,
		UpdatedBy:       userID,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("CreateHouse success")
	return &types.CreateHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
