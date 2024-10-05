package function

import "roomrover/service/inventory/model"

type InventoryFunction interface {
	GetRoomByID(roomID int64) (room *model.RoomTbl, err error)
	UpdateRoom(room *model.RoomTbl) error
	GetSericesByRoom(roomID int64) (services []*model.ServiceTbl, err error)
}
