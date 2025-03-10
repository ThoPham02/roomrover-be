package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"
	notiModel "roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContractStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContractStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractStatusLogic {
	return &UpdateContractStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractStatusLogic) UpdateContractStatus(req *types.UpdateContractStatusReq) (resp *types.UpdateContractStatusRes, err error) {
	l.Logger.Info("UpdateContractStatus: ", req)
	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contract, err := l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UpdateContractStatusRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateContractStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contract.Status = sql.NullInt64{Valid: true, Int64: req.Status}
	contract.UpdatedBy = sql.NullInt64{Valid: true, Int64: userID}
	contract.UpdatedAt = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}
	err = l.svcCtx.ContractModel.Update(l.ctx, contract)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	if req.Status == common.CONTRACT_STATUS_INACTIVE {
		noti := &notiModel.NotificationTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			Sender:    contract.RenterId.Int64,
			Receiver:  contract.LessorId.Int64,
			RefId:     contract.Id,
			RefType:   common.NOTI_TYPE_CONFIRM_CONTRACT,
			Unread:    common.NOTI_TYPE_UNREAD,
			CreatedAt: common.GetCurrentTime(),
		}

		err = l.svcCtx.NotiFunction.CreateNotification(noti)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractStatusRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		noti.Id = l.svcCtx.ObjSync.GenServiceObjID()
		noti.Sender = contract.LessorId.Int64
		noti.Receiver = contract.RenterId.Int64
		err = l.svcCtx.NotiFunction.CreateNotification(noti)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractStatusRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	l.Logger.Info("UpdateContractStatus success: ", contract)
	return &types.UpdateContractStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
