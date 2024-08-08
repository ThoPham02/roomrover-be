package function

import "roomrover/service/contract/model"

type ContractFunction interface {
	GetContractByRoomID(roomID int64) (contractID *model.ContractTbl, err error)
}
