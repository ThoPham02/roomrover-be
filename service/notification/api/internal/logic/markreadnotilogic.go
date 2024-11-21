package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/notification/api/internal/svc"
	"roomrover/service/notification/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkReadNotiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarkReadNotiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkReadNotiLogic {
	return &MarkReadNotiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkReadNotiLogic) MarkReadNoti(req *types.MarkReadReq) (resp *types.MarkReadRes, err error) {
	l.Logger.Info("MarkReadNoti ", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.MarkReadRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.NotificationModel.MarkRead(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.MarkReadRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("MarkReadNoti Success: ", userID)
	return &types.MarkReadRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
