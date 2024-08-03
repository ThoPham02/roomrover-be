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

// GetHouse
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
	house = types.House{
		HouseID:    houseModel.Id,
		HouseName:  houseModel.Name,
		Type:       int(houseModel.Type),
		Services:   []types.HouseService{},
		ProvinceID: int(houseModel.ProvinceId.Int64),
		DistrictID: int(houseModel.DistrictId.Int64),
		WardID:     int(houseModel.WardId.Int64),
		Address:    houseModel.SpecificAddress,
		CreatedAt:  houseModel.CreatedAt,
		UpdatedAt:  houseModel.UpdatedAt,
	}

	l.Logger.Info("GetHouse success", userID)

	return &types.GetHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		House: house,
	}, nil
}
