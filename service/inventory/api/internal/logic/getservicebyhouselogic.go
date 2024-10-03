package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceByHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Service By House
func NewGetServiceByHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceByHouseLogic {
	return &GetServiceByHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceByHouseLogic) GetServiceByHouse(req *types.GetServiceByHouseReq) (resp *types.GetServiceByHouseRes, err error) {
	l.Logger.Info("GetServiceByHouse", req)

	var serviceModels []*model.ServiceTbl
	var services []types.Service
	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetServiceByHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	serviceModels, err = l.svcCtx.ServiceModel.FindByHouseID(l.ctx, req.HouseID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetServiceByHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, service := range serviceModels {
		services = append(services, types.Service{
			ServiceID: service.Id,
			HouseID:   service.HouseId.Int64,
			Name:      service.Name.String,
			Price:     service.Price.Int64,
			Unit:      service.Unit.Int64,
		})
	}
	l.Logger.Info("GetServiceByHouse success", userID)
	return &types.GetServiceByHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Services: services,
	}, nil
}
