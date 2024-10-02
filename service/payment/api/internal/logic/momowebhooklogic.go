package logic

import (
	"context"

	"roomrover/service/payment/api/internal/svc"
	"roomrover/service/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MomoWebhookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMomoWebhookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MomoWebhookLogic {
	return &MomoWebhookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MomoWebhookLogic) MomoWebhook(req *types.MomoWebhookReq) (resp *types.MomoWebhookRes, err error) {
	l.Logger.Info("MomoWebhookLogic Start: ", req)

	

	return
}
