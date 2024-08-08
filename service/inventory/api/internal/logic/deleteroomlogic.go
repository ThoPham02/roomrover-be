package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete room
func NewDeleteRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoomLogic {
	return &DeleteRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoomLogic) DeleteRoom(req *types.DeleteRoomReq) (resp *types.DeleteRoomRes, err error) {
	l.Logger.Info("DeleteRoom", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contractModel, err := l.svcCtx.ContractFunction.GetContractByRoomID(req.RoomID)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if contractModel != nil {
		l.Logger.Error("DeleteRoom: Contract exists", req.RoomID)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.ROOM_HAS_CONTRACT_CODE,
				Message: common.ROOM_HAS_CONTRACT_MESS,
			},
		}, nil
	}

	err = l.svcCtx.RoomModel.Delete(l.ctx, req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("DeleteRoom Success", userID)

	return &types.DeleteRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
