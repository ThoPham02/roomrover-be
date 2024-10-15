package svc

import (
	"roomrover/sync"

	"roomrover/service/contract/job/config"
	"roomrover/service/contract/model"
	inventFunc "roomrover/service/inventory/function"
	notificationFunc "roomrover/service/notification/function"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	ObjSync sync.ObjSync

	ContractModel      model.ContractTblModel
	PaymentModel       model.PaymentTblModel
	PaymentDetailModel model.PaymentDetailTblModel
	PaymentRenterModel model.PaymentRenterTblModel
	BillModel          model.BillTblModel
	BillDetailModel    model.BillDetailTblModel
	BillPayModel       model.BillPayTblModel

	InventFunction       inventFunc.InventoryFunction
	NotificationFunction notificationFunc.NotificationFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		ObjSync:            *sync.NewObjSync(0),
		ContractModel:      model.NewContractTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentModel:       model.NewPaymentTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentDetailModel: model.NewPaymentDetailTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentRenterModel: model.NewPaymentRenterTblModel(sqlx.NewMysql(c.DataSource)),
		BillModel:          model.NewBillTblModel(sqlx.NewMysql(c.DataSource)),
		BillDetailModel:    model.NewBillDetailTblModel(sqlx.NewMysql(c.DataSource)),
		BillPayModel:       model.NewBillPayTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetInventFunction(inventFunction inventFunc.InventoryFunction) {
	sc.InventFunction = inventFunction
}

func (ctx *ServiceContext) SetNotificationFunction(notificationFunction notificationFunc.NotificationFunction) {
	ctx.NotificationFunction = notificationFunction
}
