package function

import "roomrover/service/contract/model"

type ContractFunction interface {
	CountContractByHouseID(houseID int64) (count int64, err error)
	GetContractByID(contractID int64) (contract *model.ContractTbl, err error)
	GetActiveContractByRoomID(roomID int64) (contract *model.ContractTbl, err error)
	GetPaymentByContractID(contractID int64) (payments *model.PaymentTbl, err error)
	// GetPaymentByTime(time int64) (payments []*model.PaymentTbl, err error)
	GetBillByID(id int64) (*model.BillTbl, error)
}
