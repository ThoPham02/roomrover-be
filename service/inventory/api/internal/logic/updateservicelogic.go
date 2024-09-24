package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update Service
func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateServiceLogic) UpdateService(req *types.UpdateServiceReq) (resp *types.UpdateServiceRes, err error) {
	l.Logger.Info("UpdateService", req)

	var serviceModel *model.ServiceTbl
	var userID int64
	var currentTime = common.GetCurrentTime()

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateServiceRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	serviceModel, err = l.svcCtx.ServiceModel.FindOne(l.ctx, req.ServiceID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateServiceRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if serviceModel == nil {
		l.Logger.Error(err)
		return &types.UpdateServiceRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	serviceModel.Name = req.Name
	serviceModel.Price = req.Price
	serviceModel.Type = req.Type
	serviceModel.UpdatedAt = currentTime
	serviceModel.UpdatedBy = userID
	err = l.svcCtx.ServiceModel.Update(l.ctx, serviceModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateServiceRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UpdateService success", userID)
	return &types.UpdateServiceRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
