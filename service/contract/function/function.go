package function

import "roomrover/service/contract/model"

type ContractFunction interface {
	GetPaymentByTime(time int64) (payments []*model.PaymentTbl, err error)
	GetContractByID(contractID int64) (contract *model.ContractTbl, err error)
}
