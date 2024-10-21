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

type CreateBillPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBillPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBillPayLogic {
	return &CreateBillPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBillPayLogic) CreateBillPay(req *types.CreateBillPayReq) (resp *types.CreateBillPayRes, err error) {
	l.Logger.Info("CreateBillPay: ", req)

	var userID int64

	var bill *model.BillTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateBillPayRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, err
	}

	bill, err = l.svcCtx.BillModel.FindOne(l.ctx, req.BillID)
	if err != nil {
		if err == sql.ErrNoRows || bill == nil {
			return &types.CreateBillPayRes{
				Result: types.Result{
					Code:    common.BILL_NOT_FOUND_CODE,
					Message: common.BILL_NOT_FOUND_MESS,
				},
			}, err
		}
		l.Logger.Error(err)
		return &types.CreateBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	bill.Remain = bill.Remain - req.Amount
	if bill.Remain <= 0 {
		bill.Status = common.BILL_STATUS_PAID
	}

	_, err = l.svcCtx.BillPayModel.Insert(l.ctx, &model.BillPayTbl{
		Id:      l.svcCtx.ObjSync.GenServiceObjID(),
		BillId:  req.BillID,
		UserId:  userID,
		Amount:  req.Amount,
		PayDate: req.PayDate,
		Type:    req.PayType,
		Url:     sql.NullString{Valid: true, String: req.Url},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	err = l.svcCtx.BillModel.Update(l.ctx, bill)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateBillPayRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, err
	}

	l.Logger.Info("CreateBillPay Success: ", userID)
	return &types.CreateBillPayRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
