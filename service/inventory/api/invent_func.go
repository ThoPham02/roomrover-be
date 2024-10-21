package api

import (
	"context"
	"roomrover/service/inventory/function"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.InventoryFunction = (*InventoryFunction)(nil)

type InventoryFunction struct {
	function.InventoryFunction
	logx.Logger
	InventService *InventService
}

func NewInventoryFunction(svc *InventService) *InventoryFunction {
	ctx := context.Background()

	return &InventoryFunction{
		Logger:        logx.WithContext(ctx),
		InventService: svc,
	}
}

func (contractFunc *InventoryFunction) Start() error {
	return nil
}

func (sc *InventoryFunction) GetRoomByID(roomID int64) (room *model.RoomTbl, err error) {
	return sc.InventService.Ctx.RoomModel.FindOne(context.Background(), roomID)
}

func (sc *InventoryFunction) UpdateRoom(room *model.RoomTbl) error {
	return sc.InventService.Ctx.RoomModel.Update(context.Background(), room)
}

func (sc *InventoryFunction) GetSericesByRoom(roomID int64) (services []*model.ServiceTbl, err error) {
	roomModel, err := sc.InventService.Ctx.RoomModel.FindOne(context.Background(), roomID)
	if err != nil || roomModel == nil {
		return nil, err
	}

	return sc.InventService.Ctx.ServiceModel.FindByHouseID(context.Background(), roomModel.HouseId.Int64)
}

func (sc *InventoryFunction) GetRoomsByIDs(roomIDs []int64) (rooms []*model.RoomTbl, err error) {
    return sc.InventService.Ctx.RoomModel.FindByIDs(context.Background(), roomIDs)
}

func (sc *InventoryFunction) GetHousesByIDs(houseIDs []int64) (houses []*model.HouseTbl, err error) {
	return sc.InventService.Ctx.HouseModel.FindMultiByID(context.Background(), houseIDs)
}
