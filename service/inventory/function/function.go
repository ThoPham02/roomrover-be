package function

import "roomrover/service/inventory/model"

type InventoryFunction interface {
	GetRoomByID(roomID int64) (room *model.RoomTbl, err error)
	UpdateRoom(room *model.RoomTbl) error
	GetSericesByRoom(roomID int64) (services []*model.ServiceTbl, err error)
	GetRoomsByIDs(roomIDs []int64) (rooms []*model.RoomTbl, err error)
	GetHousesByIDs(houseIDs []int64) (houses []*model.HouseTbl, err error)
}
