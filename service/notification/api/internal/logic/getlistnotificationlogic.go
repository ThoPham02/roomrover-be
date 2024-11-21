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
		var noti = types.Notification{
			NotificationID: notiModel.Id,
			SenderID:       notiModel.Sender,
			ReceiverID:     notiModel.Receiver,
			RefID:          notiModel.RefId,
			RefType:        notiModel.RefType,
			Unread:         notiModel.Unread,
			CreatedAt:      notiModel.CreatedAt,
		}

		switch notiModel.RefType {
		case common.NOTI_TYPE_CREATE_CONTACT:
			contactModel, err := l.svcCtx.InventFunction.GetContactByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			accountModel, err := l.svcCtx.AccountFunction.GetUserByID(contactModel.RenterId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			houseModel, err := l.svcCtx.InventFunction.GetHouseByID(contactModel.HouseId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}

			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.RenterId.Int64,
				Name: accountModel.FullName.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.HouseId.Int64,
				Name: houseModel.Name.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.Id,
				Name: contactModel.Datetime.Int64,
			})

		case common.NOTI_TYPE_CONFIRM_CONTACT, common.NOTI_TYPE_REJECT_CONTACT:
			contactModel, err := l.svcCtx.InventFunction.GetContactByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			accountModel, err := l.svcCtx.AccountFunction.GetUserByID(contactModel.LessorId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			houseModel, err := l.svcCtx.InventFunction.GetHouseByID(contactModel.HouseId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}

			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.LessorId.Int64,
				Name: accountModel.FullName.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.HouseId.Int64,
				Name: houseModel.Name.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contactModel.Id,
				Name: contactModel.Datetime.Int64,
			})
		case common.NOTI_TYPE_CREATE_CONTRACT, common.NOTI_TYPE_UPDATE_CONTRACT:
			contractModel, err := l.svcCtx.ContractFunction.GetContractByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			accountModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contractModel.LessorId.Int64,
				Name: accountModel.FullName.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contractModel.Id,
				Name: contractModel.Code.String,
			})
		case common.NOTI_TYPE_CONFIRM_CONTRACT:
			contractModel, err := l.svcCtx.ContractFunction.GetContractByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			accountModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.RenterId.Int64)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contractModel.RenterId.Int64,
				Name: accountModel.FullName.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contractModel.Id,
				Name: contractModel.Code.String,
			})
		case common.NOTI_TYPE_CANCEL_CONTRACT, common.NOTI_TYPE_OUT_DATE_CONTRACT, common.NOTI_TYPE_NEARLY_OUT_DATE_CONTRACT:
			contractModel, err := l.svcCtx.ContractFunction.GetContractByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   contractModel.Id,
				Name: contractModel.Code.String,
			})
		case common.NOTI_TYPE_OUT_DATE_BILL, common.NOTI_TYPE_CREATE_BILL, common.NOTI_TYPE_PAY_BILL:
			billModel, err := l.svcCtx.ContractFunction.GetBillByID(noti.RefID)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   billModel.Id,
				Name: billModel.Title.String,
			})
			noti.NotiInfos = append(noti.NotiInfos, types.NotiInfo{
				ID:   billModel.Id,
				Name: billModel.PaymentDate,
			})

		default:
			continue
		}

		notis = append(notis, noti)
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
