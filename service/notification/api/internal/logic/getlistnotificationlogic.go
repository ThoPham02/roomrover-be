package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/notification/api/internal/svc"
	"roomrover/service/notification/api/internal/types"
	"roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListNotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListNotificationLogic {
	return &GetListNotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListNotificationLogic) GetListNotification(req *types.GetListNotificationReq) (resp *types.GetListNotificationRes, err error) {
	l.Logger.Info("GetListNotification: ", req)

	var userID int64
	var notis []types.Notification
	var total int = 0
	var notiModels []*model.NotificationTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListNotificationRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	total, err = l.svcCtx.NotificationModel.CountNotisByReceiver(l.ctx, userID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListNotificationRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if total == 0 {
		return &types.GetListNotificationRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	notiModels, err = l.svcCtx.NotificationModel.GetNotisByReceiver(l.ctx, userID, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListNotificationRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, notiModel := range notiModels {
		notis = append(notis, types.Notification{
			NotificationID: notiModel.Id,
			AssigneeID:     notiModel.Sender,
			AssignerID:     notiModel.Sender,
			RefID:          notiModel.RefId,
			RefType:        notiModel.RefType,
			Status:         notiModel.Status,
			Unread:         notiModel.Unread,
			CreatedAt:      notiModel.CreatedAt,
		})
	}

	l.Logger.Info("GetListNotification Success: ", userID)
	return &types.GetListNotificationRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Notifications: notis,
		Total:         total,
	}, nil
}
