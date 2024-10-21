package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterBillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterBillLogic {
	return &FilterBillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterBillLogic) FilterBill(req *types.FilterBillReq) (resp *types.FilterBillRes, err error) {
	l.Logger.Info("FilterBill: ", req)

	var userID int64
	var total int64
	var mapPaymentContractCode = make(map[int64]string)
	var mapBillDetail = make(map[int64][]types.BillDetail)
	var bills []types.Bill
	var billModels []*model.BillTbl
	var paymentModel *model.PaymentTbl
	var constractModel *model.ContractTbl
	var billDetails []*model.BillDetailTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterBillRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	total, err = l.svcCtx.BillModel.CountByCondition(l.ctx, model.FilterCondition{})
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterBillRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	billModels, err = l.svcCtx.BillModel.FilterBillByCondition(l.ctx, model.FilterCondition{})
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterBillRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, bill := range billModels {
		paymentModel, err = l.svcCtx.PaymentModel.FindOne(l.ctx, bill.PaymentId)
		if err != nil {
			if err == model.ErrNotFound || paymentModel == nil {
				return &types.FilterBillRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}
			l.Logger.Error(err)
			return &types.FilterBillRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		constractModel, err = l.svcCtx.ContractModel.FindOne(l.ctx, paymentModel.ContractId)
		if err != nil {
			if err == model.ErrNotFound || constractModel == nil {
				return &types.FilterBillRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}
			l.Logger.Error(err)
			return &types.FilterBillRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		billDetails, err = l.svcCtx.BillDetailModel.GetDetailByBillID(l.ctx, bill.Id)
		if err != nil {
			l.Logger.Error(err)
			return &types.FilterBillRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		for _, billDetail := range billDetails {
			mapBillDetail[bill.Id] = append(mapBillDetail[bill.Id], types.BillDetail{
				BillDetailID: billDetail.Id,
				BillID:       billDetail.BillId.Int64,
				Name:         billDetail.Name.String,
				Price:        billDetail.Price.Int64,
				Type:         billDetail.Type.Int64,
				Quantity:     billDetail.Quantity.Int64,
			})
		}

		mapPaymentContractCode[bill.PaymentId] = constractModel.Code.String
	}

	for _, bill := range billModels {
		bills = append(bills, types.Bill{
			BillID:       bill.Id,
			Title:        bill.Title.String,
			ContractCode: mapPaymentContractCode[bill.PaymentId],
			PaymentID:    bill.PaymentId,
			PaymentDate:  bill.PaymentDate.Int64,
			Amount:       bill.Amount,
			Remain:       bill.Remain,
			Status:       bill.Status,
			BillDetails:  mapBillDetail[bill.Id],
		})
	}

	l.Logger.Info("FilterBill Success:", userID)
	return &types.FilterBillRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Bills: bills,
		Total: total,
	}, nil
}
