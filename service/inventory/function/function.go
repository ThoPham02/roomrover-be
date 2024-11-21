package function

import "roomrover/service/inventory/model"

type InventoryFunction interface {
	GetRoomByID(roomID int64) (room *model.RoomTbl, err error)
	UpdateRoom(room *model.RoomTbl) error
	GetSericesByRoom(roomID int64) (services []*model.ServiceTbl, err error)
	GetRoomsByIDs(roomIDs []int64) (rooms []*model.RoomTbl, err error)
	GetHousesByIDs(houseIDs []int64) (houses []*model.HouseTbl, err error)
	GetHouseRoomByRoomID(roomID int64) (houseRoom *model.HouseRoomTbl, err error)
	CountRoomActiveByHouseID(houseID int64) (count int64, err error)
	GetHouseByID(houseID int64) (house *model.HouseTbl, err error)
	UpdateHouse(house *model.HouseTbl) error
	GetContactByID(contactID int64) (contact *model.ContactTbl, err error)
}
