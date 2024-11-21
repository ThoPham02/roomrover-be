package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContactLogic {
	return &DeleteContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContactLogic) DeleteContact(req *types.DeleteContactReq) (resp *types.DeleteContactRes, err error) {
	l.Logger.Info("DeleteContact:", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContactRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.NotiFunction.DeleteNotiByRefID(req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	
	err = l.svcCtx.ContactModel.Delete(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("DeleteContact Success:", userID)
	return &types.DeleteContactRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
