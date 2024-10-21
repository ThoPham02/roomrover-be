package logic

import (
	"context"
	"database/sql"
	"fmt"
	"roomrover/common"
	"roomrover/service/contract/job/svc"
	"roomrover/service/contract/model"

	notiModel "roomrover/service/notification/model"

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

	paymentModels, err = l.svcCtx.PaymentModel.FilterPaymentByTime(l.ctx, currentTime+1*86400000) // thong bao 1 ngay truoc
	if err != nil {
		l.Logger.Error(err)
	}
	for _, paymentModel := range paymentModels {
		var paymentDetails []*model.PaymentDetailTbl
		var contractModel *model.ContractTbl

		paymentDetails, err = l.svcCtx.PaymentDetailModel.GetPaymentDetailByPaymentID(l.ctx, paymentModel.Id)
		if err != nil {
			l.Logger.Error(err)
			continue
		}
		contractModel, err = l.svcCtx.ContractModel.FindOne(l.ctx, paymentModel.ContractId)
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
			Id: l.svcCtx.ObjSync.GenServiceObjID(),
			// Title:       "Hóa đơn tháng " + common.GetMonthYear(paymentModel.NextBill),
			PaymentId:   paymentModel.Id,
			PaymentDate: sql.NullInt64{Valid: true, Int64: currentTime + 6*86400000}, // han thanh toan sau 5 ngay
			Amount:      paymentModel.Amount,
			Discount:    sql.NullInt64{Valid: true, Int64: paymentModel.Discount},
			Remain:      paymentModel.Amount - paymentModel.Discount,
			Status:      common.BILL_STATUS_UNPAID,
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

		noti := &notiModel.NotificationTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			Sender:    contractModel.LessorId.Int64,
			Receiver:  contractModel.RenterId.Int64,
			RefId:     billModel.Id,
			RefType:   common.NOTIFICATION_REF_TYPE_BILL,
			Title:     fmt.Sprintf("Hoàn thành hóa đơn thanh toán tháng %d", common.GetBillIndexByTime(contractModel.CheckIn.Int64, currentTime)),
			DueDate:   billModel.PaymentDate.Int64,
			Status:    common.NOTI_STATUS_PENDING,
			Unread:    common.NOTI_TYPE_UNREAD,
			CreatedAt: currentTime,
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
