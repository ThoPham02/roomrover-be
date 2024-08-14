package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create Service
func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateServiceLogic) CreateService(req *types.CreateServiceReq) (resp *types.CreateServiceRes, err error) {
	l.Logger.Info("CreateService", req)

	var userID int64
	var currentTime = common.GetCurrentTime()
	var service types.Service

	var serviceModel *model.ServiceTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateServiceRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	serviceModel = &model.ServiceTbl{
		Id:        l.svcCtx.ObjSync.GenServiceObjID(),
		HouseId:   req.HouseID,
		Name:      req.Name,
		Price:     req.Price,
		Type:      req.Type,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	_, err = l.svcCtx.ServiceModel.Insert(l.ctx, serviceModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateServiceRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
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

	l.Logger.Info("CreateService", userID)
	return &types.CreateServiceRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Service: service,
	}, nil
}
