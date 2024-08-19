package function

import "roomrover/service/contract/model"

type ContractFunction interface {
	GetContractByRoomID(roomID int64) (contractID *model.ContractTbl, err error)
	GetContractByTime(time int64) (contractIDs []*model.ContractTbl, err error)
	GetContractDetailByContractID(contractID int64) (contractDetail []*model.ContractDetailTbl, err error)
	CountRenterByContractID(contractID int64) (count int64, err error)
	UpdateContract(contract *model.ContractTbl) (err error)
}
