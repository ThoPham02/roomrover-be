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
		Id:      l.svcCtx.ObjSync.GenServiceObjID(),
		HouseId: sql.NullInt64{Valid: true, Int64: req.HouseID},
		Name:    sql.NullString{Valid: true, String: req.Name},
		Price:   sql.NullInt64{Valid: true, Int64: req.Price},
		Unit:    sql.NullInt64{Valid: true, Int64: req.Unit},
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
		HouseID:   serviceModel.HouseId.Int64,
		Name:      serviceModel.Name.String,
		Price:     serviceModel.Price.Int64,
		Unit:      serviceModel.Unit.Int64,
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
