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