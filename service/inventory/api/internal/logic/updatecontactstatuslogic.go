package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"
	notiModel "roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContactStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContactStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContactStatusLogic {
	return &UpdateContactStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContactStatusLogic) UpdateContactStatus(req *types.UpdateContactStatusReq) (resp *types.UpdateContactStatusRes, err error) {
	l.Logger.Info("UpdateContactStatus: ", req)

	var userID int64
	var contactModel *model.ContactTbl
	var currentTime = common.GetCurrentTime()
	var refType int64 =  common.NOTI_TYPE_CONFIRM_CONTACT

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContactStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contactModel, err = l.svcCtx.ContactModel.FindOne(l.ctx, req.ID)
	if err != nil || contactModel == nil {
		l.Logger.Error(err)
		return &types.UpdateContactStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contactModel.Status = req.Status
	err = l.svcCtx.ContactModel.Update(l.ctx, contactModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContactStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if contactModel.Status == common.CONTACT_STATUS_TYPE_CANCEL {
		refType = common.NOTI_TYPE_REJECT_CONTACT
	}
	err = l.svcCtx.NotiFunction.CreateNotification(&notiModel.NotificationTbl{
		Id:        l.svcCtx.ObjSync.GenServiceObjID(),
		Sender:    userID,
		Receiver:  contactModel.RenterId.Int64,
		RefId:     contactModel.Id,
		RefType:   refType,
		Unread:    common.NOTI_TYPE_UNREAD,
		CreatedAt: currentTime,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContactStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UpdateContactStatus Success: ", userID)
	return &types.UpdateContactStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
