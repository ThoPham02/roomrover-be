package function

import "roomrover/service/inventory/model"

type InventoryFunction interface {
	GetRoomByID(roomID int64) (room *model.RoomTbl, err error)
}
