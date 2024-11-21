package logic

import (
	"context"

	"roomrover/common"
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
	l.Logger.Info("UpdateRenterStatus ", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRenterStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	renterModel, err := l.svcCtx.PaymentRenterModel.FindOne(l.ctx, req.ID)
	if err!= nil || renterModel == nil {
		l.Logger.Error(err)
        return &types.UpdateRenterStatusRes{
            Result: types.Result{
                Code:    common.DB_ERR_CODE,
                Message: common.DB_ERR_MESS,
            },
        }, nil
    }

	renterModel.Status = req.Status
	err = l.svcCtx.PaymentRenterModel.Update(l.ctx, renterModel)
	if err!= nil {
		l.Logger.Error(err)
        return &types.UpdateRenterStatusRes{
            Result: types.Result{
                Code:    common.DB_ERR_CODE,
                Message: common.DB_ERR_MESS,
            },
        }, nil
    }

	l.Logger.Info("UpdateRenterStatus Success", userID)
	return &types.UpdateRenterStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
