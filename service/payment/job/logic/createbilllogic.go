package logic

import (
	"context"
	"roomrover/service/payment/job/svc"

	contractModel "roomrover/service/contract/model"

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
	// var currentTime = common.GetCurrentTime()
	// var err error

	// var contractModels []*contractModel.ContractTbl

	// l.Logger.Info("Job create bill by time: ", currentTime)

	// contractModels, err = l.svcCtx.ContractFunc.GetContractByTime(currentTime)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return nil
	// }

	// l.Logger.Info(len(contractModels))

	// for _, contractModel := range contractModels {
	// 	var i int64 = common.GetBillIndexByTime(contractModel.Start, currentTime)

	// 	err = l.CreateBillFromContract(contractModel, i+1)
	// 	if err != nil {
	// 		l.Logger.Error(err)
	// 		continue
	// 	}

	// 	contractModel.NextBill = common.GetNextMonthDate(contractModel.Start, int(i+1))
	// 	err = l.svcCtx.ContractFunc.UpdateContract(contractModel)
	// 	if err != nil {
	// 		l.Logger.Error(err)
	// 		continue
	// 	}
	// }

	// l.Logger.Info("Job create bill by time Success: ", currentTime)
	return nil
}

func (l *CreateBillLogic) CreateBillFromContract(contractModel *contractModel.ContractTbl, month int64) error {
	// var err error
	// var billID = l.svcCtx.ObjSync.GenServiceObjID()
	// var currentTime = common.GetCurrentTime()
	// var total int64

	// var billModel *model.Bill
	// var billDetailModels []*model.BillDetail

	// l.Logger.Info("Create bill from contractID: ", contractModel.Id)

	// // get contract detail
	// contractDetailModel, err := l.svcCtx.ContractFunc.GetContractDetailByContractID(contractModel.Id)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return err
	// }

	// // create bill detail from contract detail
	// for _, contractDetail := range contractDetailModel {
	// 	billDetail := &model.BillDetail{
	// 		Id:                l.svcCtx.ObjSync.GenServiceObjID(),
	// 		BillId:            billID,
	// 		ContractServiceId: contractDetail.ServiceId,
	// 		Price:             contractDetail.Price,
	// 		Type:              contractDetail.Type,
	// 	}

	// 	if contractDetail.Type == common.CONTRACT_DETAIL_TYPE_FIXED {
	// 		billDetail.Quantity = 1
	// 	} else if contractDetail.Type == common.CONTRACT_DETAIL_TYPE_FIXED_USER {
	// 		count, err := l.svcCtx.ContractFunc.CountRenterByContractID(contractModel.Id)
	// 		if err != nil {
	// 			l.Logger.Error(err)
	// 			return err
	// 		}
	// 		billDetail.Quantity = count
	// 	} else {
	// 		billDetail.Quantity = common.NO_USE
	// 	}

	// 	billDetailModels = append(billDetailModels, billDetail)
	// 	if billDetail.Quantity == common.NO_USE {
	// 		continue
	// 	}
	// 	total += billDetail.Price * billDetail.Quantity
	// }

	// // create bill
	// billModel = &model.Bill{
	// 	Id:         billID,
	// 	ContractId: contractModel.Id,
	// 	Total:      total,
	// 	Paid:       0,
	// 	Status:     common.BILL_STATUS_UNPAID,
	// 	Month:      month,
	// 	CreatedAt:  currentTime,
	// 	UpdatedAt:  currentTime,
	// }

	// err = l.svcCtx.BillDetailModel.InsertMulti(l.ctx, billDetailModels)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return err
	// }
	// _, err = l.svcCtx.BillModel.Insert(l.ctx, billModel)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return err
	// }

	// l.Logger.Info("Create bill from contractID Success: ", contractModel.Id)
	return nil
}
