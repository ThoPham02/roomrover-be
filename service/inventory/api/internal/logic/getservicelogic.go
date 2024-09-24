package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Service
func NewGetServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceLogic {
	return &GetServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceLogic) GetService(req *types.GetServiceReq) (resp *types.GetServiceRes, err error) {
	l.Logger.Info("GetService", req)

	var serviceModel *model.ServiceTbl
	var service types.Service
	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetServiceRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	serviceModel, err = l.svcCtx.ServiceModel.FindOne(l.ctx, req.ServiceID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetServiceRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	service = types.Service{
		ServiceID: serviceModel.Id,
		HouseID:   serviceModel.HouseId,
		Name:      serviceModel.Name,
		Price:     serviceModel.Price,
		Type:      serviceModel.Type,
		CreatedAt: serviceModel.CreatedAt,
		UpdatedAt: serviceModel.UpdatedAt,
		CreatedBy: serviceModel.CreatedBy,
		UpdatedBy: serviceModel.UpdatedBy,
	}

	l.Logger.Info("GetService success", userID)
	return &types.GetServiceRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Service: service,
	}, nil
}
