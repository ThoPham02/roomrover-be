package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"
	notiModel "roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContactLogic {
	return &CreateContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContactLogic) CreateContact(req *types.CreateContactReq) (resp *types.CreateContactRes, err error) {
	l.Logger.Info("CreateContact: ", req)

	var userID int64
	var currentTime = common.GetCurrentTime()

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContactRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil || houseModel == nil {
		l.Logger.Info(err)
		return &types.CreateContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	var contact = &model.ContactTbl{
		Id:       l.svcCtx.ObjSync.GenServiceObjID(),
		HouseId:  sql.NullInt64{Valid: true, Int64: req.HouseID},
		RenterId: sql.NullInt64{Valid: true, Int64: userID},
		LessorId: sql.NullInt64{Valid: true, Int64: req.LessorID},
		Datetime: sql.NullInt64{Valid: true, Int64: req.Datetime},
		Status:   common.CONTACT_STATUS_TYPE_WATTING,
	}
	_, err = l.svcCtx.ContactModel.Insert(l.ctx, contact)
	if err != nil {
		l.Logger.Info(err)
		return &types.CreateContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.NotiFunction.CreateNotification(&notiModel.NotificationTbl{
		Id:          l.svcCtx.ObjSync.GenServiceObjID(),
		Sender:      userID,
		Receiver:    req.LessorID,
		RefId:       contact.Id,
		RefType:     common.NOTI_TYPE_CREATE_CONTACT,
		Unread:      common.NOTI_TYPE_UNREAD,
		CreatedAt:   currentTime,
	})
	if err != nil {
		l.Logger.Info(err)
		return &types.CreateContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("CreateContact Success: ", userID)
	return &types.CreateContactRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
