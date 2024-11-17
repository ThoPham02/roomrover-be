package logic

import (
	"context"

	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRenterStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRenterStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRenterStatusLogic {
	return &UpdateRenterStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRenterStatusLogic) UpdateRenterStatus(req *types.UpdateRenterStatusReq) (resp *types.UpdateRenterStatusRes, err error) {
	// todo: add your logic here and delete this line

	return
}
