package api

import (
	"context"
	"roomrover/service/inventory/function"

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
