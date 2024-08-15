package logic

import (
	"context"

	"roomrover/service/payment/api/internal/svc"
	"roomrover/service/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFilePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFilePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFilePaymentLogic {
	return &UploadFilePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFilePaymentLogic) UploadFilePayment(req *types.UploadFilePaymentReq) (resp *types.UploadFilePaymentRes, err error) {
	// todo: add your logic here and delete this line

	return
}
