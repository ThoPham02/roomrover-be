package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete Service
func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteServiceLogic) DeleteService(req *types.DeleteServiceReq) (resp *types.DeleteServiceRes, err error) {
	l.Logger.Info("DeleteRoom", req)

	var userID int64
	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteServiceRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.ServiceModel.Delete(l.ctx, req.ServiceID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteServiceRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("DeleteRoom success", userID)
	return &types.DeleteServiceRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
