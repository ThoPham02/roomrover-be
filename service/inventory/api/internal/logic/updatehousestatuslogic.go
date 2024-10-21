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

type UpdateHouseStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update House Status
func NewUpdateHouseStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHouseStatusLogic {
	return &UpdateHouseStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHouseStatusLogic) UpdateHouseStatus(req *types.UpdateHouseStatusReq) (resp *types.UpdateHouseStatusRes, err error) {
	l.Logger.Info("UpdateHouseStatus", req)

	var userID int64
	var houseModel *model.HouseTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, req.HouseID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UpdateHouseStatusRes{
				Result: types.Result{
					Code:    common.HOUSE_NOT_FOUND_CODE,
					Message: common.HOUSE_NOT_FOUND_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateHouseStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	houseModel.Status = req.Status
	houseModel.UpdatedBy = sql.NullInt64{Valid: true, Int64: userID}
	houseModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}

	err = l.svcCtx.HouseModel.Update(l.ctx, houseModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateHouseStatusRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UpdateHouseStatus Success", userID)
	return &types.UpdateHouseStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
