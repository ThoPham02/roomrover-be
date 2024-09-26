package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomByHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get Room By House
func NewGetRoomByHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomByHouseLogic {
	return &GetRoomByHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomByHouseLogic) GetRoomByHouse(req *types.GetRoomByHouseReq) (resp *types.GetRoomByHouseRes, err error) {
	l.Logger.Info("GetRoomByHouse", req)

	var userID int64
	var roomModels []*model.RoomTbl
	var rooms []types.Room
	var total int

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomByHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	roomModels, total, err = l.svcCtx.RoomModel.FindByHouseID(l.ctx, req.HouseID, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomByHouseRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, room := range roomModels {
		rooms = append(rooms, types.Room{
			RoomID:    room.Id,
			HouseID:   room.HouseId,
			Name:      room.Name,
			Capacity:  room.Capacity,
			Remain:    room.Remain,
			Status:    room.Status,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
			CreatedBy: room.CreatedBy,
			UpdatedBy: room.UpdatedBy,
		})
	}

	l.Logger.Info("GetRoomByHouse success", userID)
	return &types.GetRoomByHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total: total,
		Rooms: rooms,
	}, nil
}
