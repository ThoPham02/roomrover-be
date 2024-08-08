package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create room
func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreateRoomReq) (resp *types.CreateRoomRes, err error) {
	l.Logger.Info("CreateRoom", req)

	var userID int64
	var currentTime = common.GetCurrentTime()

	var houseModel *model.HouseTbl
	var roomModel *model.RoomTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil {
		if err == model.ErrNotFound {
			l.Logger.Error(err)
			return &types.CreateRoomRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.CreateRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if houseModel.UserId != userID {
		l.Logger.Error("User is not owner of this house")
		return &types.CreateRoomRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	roomModel = &model.RoomTbl{
		Id:        l.svcCtx.ObjSync.GenServiceObjID(),
		HouseId:   req.HouseID,
		Name:      req.Name,
		Status:    common.ROOM_STATUS_DRAFT,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		CreatedBy: userID,
		UpdatedBy: userID,
	}
	_, err = l.svcCtx.RoomModel.Insert(l.ctx, roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("CreateRoom success: ", userID)
	return &types.CreateRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
