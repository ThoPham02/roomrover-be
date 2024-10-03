package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Room
func NewGetRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomLogic {
	return &GetRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomLogic) GetRoom(req *types.GetRoomReq) (resp *types.GetRoomRes, err error) {
	l.Logger.Info("GetRoom", req)

	var roomModel *model.RoomTbl
	var room types.Room
	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	roomModel, err = l.svcCtx.RoomModel.FindOne(l.ctx, req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	room = types.Room{
		RoomID:   roomModel.Id,
		HouseID:  roomModel.HouseId.Int64,
		Name:     roomModel.Name.String,
		Status:   roomModel.Status,
		Capacity: roomModel.Capacity.Int64,
		EIndex:   roomModel.EIndex.Int64,
		WIndex:   roomModel.WIndex.Int64,
	}

	l.Logger.Info("GetRoom success", userID)
	return &types.GetRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Room: room,
	}, nil
}
