package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update User Info
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserRes, err error) {
	l.Logger.Info("UpdateUser", req)

	var userID int64
	var user types.User

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateUserRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	// Check if the user exists
	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.UpdateUserRes{
				Result: types.Result{
					Code:    common.USER_NOT_FOUND_CODE,
					Message: common.USER_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.UpdateUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	userModel.FullName = sql.NullString{String: req.FullName, Valid: true}
	userModel.Birthday = sql.NullInt64{Int64: req.Dob, Valid: true}
	userModel.AvatarUrl = sql.NullString{String: req.AvatarUrl, Valid: true}
	userModel.UpdatedAt = common.GetCurrentTime()

	err = l.svcCtx.UserModel.Update(l.ctx, userModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:    userID,
		Phone:     userModel.Phone,
		FullName:  userModel.FullName.String,
		Birthday:  userModel.Birthday.Int64,
		AvatarUrl: userModel.AvatarUrl.String,
		Address:   userModel.Address.String,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	l.Logger.Info("UpdateUser success")
	return &types.UpdateUserRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		User: user,
	}, nil
}
