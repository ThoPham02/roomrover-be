package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	model "roomrover/service/contract/model"
	inventoryModel "roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractLogic {
	return &UpdateContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractLogic) UpdateContract(req *types.UpdateContractReq) (resp *types.UpdateContractRes, err error) {
	l.Logger.Info("UpdateContract", req)

	var userID int64
	var currentTime = common.GetCurrentTime()

	var paymentRenters []types.PaymentRenter

	var contractModel *model.ContractTbl
	var renterModel *accountModel.UserTbl
	var lessorModel *accountModel.UserTbl
	var roomModel *inventoryModel.RoomTbl
	var paymentModel *model.PaymentTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	err = json.Unmarshal([]byte(req.PaymentRenter), &paymentRenters)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	for _, renter := range paymentRenters {
		userCheck, err := l.svcCtx.AccountFunction.GetUserByID(renter.RenterID)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		if userCheck == nil {
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
	}

	contractModel, err = l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.CONTRACT_NOT_FOUND_CODE,
					Message: common.CONTRACT_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	paymentModel, err = l.svcCtx.PaymentModel.FindByContractID(l.ctx, contractModel.Id)
	if err != nil || paymentModel == nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	renterModel, err = l.svcCtx.AccountFunction.GetUserByID(req.RenterID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err = l.svcCtx.AccountFunction.GetUserByID(req.LessorID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	roomModel, err = l.svcCtx.InventFunction.GetRoomByID(req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if renterModel == nil || lessorModel == nil || roomModel == nil {
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	contractModel = &model.ContractTbl{
		Id:            contractModel.Id,
		Code:          contractModel.Code,
		Status:        sql.NullInt64{Valid: true, Int64: req.Status},
		RenterId:      sql.NullInt64{Valid: true, Int64: renterModel.Id},
		RenterNumber:  sql.NullString{Valid: true, String: renterModel.CCCDNumber.String},
		RenterDate:    sql.NullInt64{Valid: true, Int64: renterModel.CCCDDate.Int64},
		RenterAddress: sql.NullString{Valid: true, String: renterModel.Address.String},
		RenterName:    sql.NullString{Valid: true, String: renterModel.FullName.String},
		LessorId:      sql.NullInt64{Valid: true, Int64: lessorModel.Id},
		LessorNumber:  sql.NullString{Valid: true, String: lessorModel.CCCDNumber.String},
		LessorDate:    sql.NullInt64{Valid: true, Int64: lessorModel.CCCDDate.Int64},
		LessorAddress: sql.NullString{Valid: true, String: lessorModel.Address.String},
		LessorName:    sql.NullString{Valid: true, String: lessorModel.FullName.String},
		RoomId:        sql.NullInt64{Valid: true, Int64: roomModel.Id},
		CheckIn:       sql.NullInt64{Valid: true, Int64: req.CheckIn},
		Duration:      sql.NullInt64{Valid: true, Int64: req.Duration},
		Purpose:       sql.NullString{Valid: true, String: req.Purpose},
		CreatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		UpdatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		CreatedBy:     sql.NullInt64{Valid: true, Int64: userID},
		UpdatedBy:     sql.NullInt64{Valid: true, Int64: userID},
	}

	paymentModel = &model.PaymentTbl{
		Id:          paymentModel.Id,
		ContractId:  contractModel.Id,
		Amount:      req.Amount,
		Discount:    req.Discount,
		Deposit:     req.Deposit,
		DepositDate: req.DepositDate,
		NextBill:    common.GetNextMonthDate(req.CheckIn),
	}

	err = l.svcCtx.PaymentRenterModel.DeleteByPaymentID(l.ctx, paymentModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, renter := range paymentRenters {
		_, err = l.svcCtx.PaymentRenterModel.Insert(l.ctx, &model.PaymentRenterTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			PaymentId: sql.NullInt64{Valid: true, Int64: paymentModel.Id},
			UserId:    sql.NullInt64{Valid: true, Int64: renter.RenterID},
		})
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	roomModel.EIndex = sql.NullInt64{Int64: req.EIndex, Valid: true}
	roomModel.WIndex = sql.NullInt64{Int64: req.WIndex, Valid: true}
	roomModel.Status = common.ROOM_STATUS_RENTED
	err = l.svcCtx.InventFunction.UpdateRoom(roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.PaymentModel.Update(l.ctx, paymentModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.ContractModel.Update(l.ctx, contractModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UpdateContract Success: ", userID)
	return &types.UpdateContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
