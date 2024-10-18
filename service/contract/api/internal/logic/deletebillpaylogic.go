package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBillPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBillPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBillPayLogic {
	return &DeleteBillPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBillPayLogic) DeleteBillPay(req *types.DeleteBillPayReq) (resp *types.DeleteBillPayRes, err error) {
	l.Logger.Info("DeleteBillPay: ", req)

	var userID int64
	var billPay *model.BillPayTbl
	var bill *model.BillTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteBillPayRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, err
	}
	billPay, err = l.svcCtx.BillPayModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows || billPay == nil {
			return &types.DeleteBillPayRes{
				Result: types.Result{
					Code:    common.BILL_PAY_NOT_FOUND_CODE,
					Message: common.BILL_PAY_NOT_FOUND_MESS,
				},
			}, err
		}
		l.Logger.Error(err)
		return &types.DeleteBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	bill, err = l.svcCtx.BillModel.FindOne(l.ctx, billPay.BillId)
	if err != nil {
		if err == sql.ErrNoRows || bill == nil {
			return &types.DeleteBillPayRes{
				Result: types.Result{
					Code:    common.BILL_NOT_FOUND_CODE,
					Message: common.BILL_NOT_FOUND_MESS,
				},
			}, err
		}
		l.Logger.Error(err)
		return &types.DeleteBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	bill.Remain = bill.Remain + billPay.Amount
	if bill.Remain > 0 {
		bill.Status = common.BILL_STATUS_UNPAID
	}
	err = l.svcCtx.BillModel.Update(l.ctx, bill)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	err = l.svcCtx.BillPayModel.Delete(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	l.Logger.Info("DeleteBillPay Success: ", userID)
	return &types.DeleteBillPayRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
