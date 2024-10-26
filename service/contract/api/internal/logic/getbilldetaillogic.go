package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBillDetailLogic {
	return &GetBillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBillDetailLogic) GetBillDetail(req *types.GetBillDetailReq) (resp *types.GetBillDetailRes, err error) {
	l.Logger.Info("GetBillDetail: ", req)

	var userID int64
	var details []types.BillDetail

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	billModel, err := l.svcCtx.BillModel.FindOne(l.ctx, req.ID)
	if err != nil || billModel == nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	paymentModel, err := l.svcCtx.PaymentModel.FindOne(l.ctx, billModel.PaymentId)
	if err != nil || paymentModel == nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contractModel, err := l.svcCtx.ContractModel.FindOne(l.ctx, paymentModel.ContractId)
	if err != nil || contractModel == nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
	if err != nil || lessorModel == nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	renterModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
	if err != nil || renterModel == nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	billDetails, err := l.svcCtx.BillDetailModel.GetDetailByBillID(l.ctx, billModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	for _, v := range billDetails {
		details = append(
			details,
			types.BillDetail{
				BillDetailID: v.Id,
				BillID:       v.BillId.Int64,
				Name:         v.Name.String,
				Price:        v.Price.Int64,
				Type:         v.Type.Int64,
				Quantity:     v.Quantity.Int64,
			},
		)
	}
	l.Logger.Info("GetBillDetail Success: ", userID)
	return &types.GetBillDetailRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Bill: types.Bill{
			BillID:       billModel.Id,
			Title:        billModel.Title.String,
			ContractCode: contractModel.Code.String,
			RenterID:     renterModel.Id,
			RenterName:   renterModel.FullName.String,
			RenterPhone:  renterModel.Phone,
			LessorID:     lessorModel.Id,
			LessorName:   lessorModel.FullName.String,
			LessorPhone:  lessorModel.Phone,
			PaymentID:    billModel.PaymentId,
			PaymentDate:  billModel.PaymentDate.Int64,
			Amount:       billModel.Amount,
			Remain:       billModel.Remain,
			Status:       billModel.Status,
			BillDetails:  details,
		},
	}, nil
}
