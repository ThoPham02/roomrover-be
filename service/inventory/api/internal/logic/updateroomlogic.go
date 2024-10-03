package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update Room
func NewUpdateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoomLogic {
	return &UpdateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoomLogic) UpdateRoom(req *types.UpdateRoomReq) (resp *types.UpdateRoomRes, err error) {
	l.Logger.Info("UpdateRoom", req)

	var roomModel *model.RoomTbl
	var userID int64
	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	roomModel, err = l.svcCtx.RoomModel.FindOne(l.ctx, req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if roomModel == nil {
		l.Logger.Error(err)
		return &types.UpdateRoomRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	roomModel.Name = sql.NullString{String: req.Name, Valid: true}
	roomModel.Capacity = sql.NullInt64{Int64: req.Capacity, Valid: true}

	err = l.svcCtx.RoomModel.Update(l.ctx, roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UpdateRoom success", userID)
	return &types.UpdateRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
