package logic

import (
	"context"

	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Upload file house
func NewUploadFileHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileHouseLogic {
	return &UploadFileHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileHouseLogic) UploadFileHouse(req *types.UploadFileHouseReq) (resp *types.UploadFileHouseRes, err error) {
	// todo: add your logic here and delete this line

	return
}
