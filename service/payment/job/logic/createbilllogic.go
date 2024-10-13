package logic

import (
	"context"
	"database/sql"
	"roomrover/common"
	contractModel "roomrover/service/contract/model"
	notificationModel "roomrover/service/notification/model"
	"roomrover/service/payment/job/svc"
	"roomrover/service/payment/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBillLogic {
	return &CreateBillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBillLogic) CreateBillByTime() error {
	var currentTime = common.GetCurrentTime()
	var err error

	var paymentModels []*model.PaymentTbl

	l.Logger.Info("CreateBillByTime Start Time: ", currentTime)

	paymentModels, err = l.svcCtx.PaymentModel.FilterPaymentByTime(l.ctx, currentTime+5*86400000)
	if err != nil {
		l.Logger.Error(err)
	}
	for _, paymentModel := range paymentModels {
		var paymentDetails []*model.PaymentDetailTbl
		var contractModel *contractModel.ContractTbl

		paymentDetails, err = l.svcCtx.PaymentDetailModel.GetPaymentDetailByPaymentID(l.ctx, paymentModel.Id)
		if err != nil {
			l.Logger.Error(err)
			continue
		}
		contractModel, err = l.svcCtx.ContractFunction.GetContractByID(paymentModel.ContractId)
		if err != nil || contractModel == nil {
			l.Logger.Error(err)
			continue
		}
		paymentModel.NextBill = common.GetNextMonthDate(contractModel.CheckIn.Int64)
		err = l.svcCtx.PaymentModel.Update(l.ctx, paymentModel)
		if err != nil {
			l.Logger.Error(err)
			continue
		}

		var billModel = &model.BillTbl{
			Id:          l.svcCtx.ObjSync.GenServiceObjID(),
			PaymentId:   sql.NullInt64{Valid: true, Int64: paymentModel.Id},
			PaymentDate: sql.NullInt64{Valid: true, Int64: currentTime + 5*86400000},
			Amount:      sql.NullInt64{Valid: true, Int64: paymentModel.Amount},
			Discount:    sql.NullInt64{Valid: true, Int64: paymentModel.Discount},
			Status:      sql.NullInt64{Valid: true, Int64: common.BILL_STATUS_UNPAID},
		}
		_, err = l.svcCtx.BillModel.Insert(l.ctx, billModel)
		if err != nil {
			l.Logger.Error(err)
			continue
		}
		for _, paymentDetail := range paymentDetails {
			var billDetailModel = &model.BillDetailTbl{
				Id:       l.svcCtx.ObjSync.GenServiceObjID(),
				BillId:   sql.NullInt64{Valid: true, Int64: billModel.Id},
				Name:     paymentDetail.Name,
				Price:    paymentDetail.Price,
				Type:     paymentDetail.Type,
				Quantity: sql.NullInt64{Valid: true, Int64: 0},
				Status:   sql.NullInt64{Valid: true, Int64: common.PAYMENT_DETAIL_STATUS_DONE},
			}

			switch paymentDetail.Type.Int64 {
			case common.PAYMENT_DETAIL_TYPE_FIXED:
				billDetailModel.Quantity.Int64 = 1
			case common.PAYMENT_DETAIL_TYPE_USAGE:
				billDetailModel.Quantity.Int64 = 0
				billDetailModel.Status.Int64 = common.PAYMENT_DETAIL_STATUS_DRAF
			case common.PAYMENT_DETAIL_TYPE_FIXED_USER:
				var count int64
				count, err = l.svcCtx.PaymentRenterModel.CountRentersByPaymentID(l.ctx, paymentModel.Id)
				if err != nil {
					l.Logger.Error(err)
					continue
				}
				billDetailModel.Quantity.Int64 = count
			}

			_, err = l.svcCtx.BillDetailModel.Insert(l.ctx, billDetailModel)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
		}

		noti := &notificationModel.NotificationTbl{
			Id:          l.svcCtx.ObjSync.GenServiceObjID(),
			Sender:      contractModel.LessorId.Int64,
			Receiver:    contractModel.RenterId.Int64,
			RefId:       billModel.Id,
			RefType:     common.NOTIFICATION_REF_TYPE_BILL,
			Title:       "",
			Description: "",
			Priority:    0,
			DueDate:     0,
			Status:      0,
			Unread:      0,
			CreatedAt:   0,
		}
		err = l.svcCtx.NotificationFunction.CreateNotification(noti)
		if err != nil {
			l.Logger.Error(err)
			continue
		}
	}

	l.Logger.Info("CreateBillByTime Start Time Success: ", common.GetCurrentTime())

	return nil
}
