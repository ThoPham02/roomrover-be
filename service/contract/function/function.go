package function

import "roomrover/service/contract/model"

type ContractFunction interface {
	CountContractByHouseID(houseID int64) (count int64, err error)
	GetContractByID(contractID int64) (contract *model.ContractTbl, err error)
	GetActiveContractByRoomID(roomID int64) (contract *model.ContractTbl, err error)
	GetPaymentByContractID(contractID int64) (payments *model.PaymentTbl, err error)
	// GetPaymentByTime(time int64) (payments []*model.PaymentTbl, err error)
	GetBillByID(id int64) (*model.BillTbl, error)
	GetBillPayByContractID(contractID int64) (billPays []*model.BillPayTbl, err error)
	GetContractByRoom(roomID int64) ([]*model.ContractTbl, error)
	GetBillByContractID(contractID int64) (bills []*model.BillTbl, err error)
}
