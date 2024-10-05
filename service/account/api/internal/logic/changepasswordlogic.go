package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"
	"roomrover/service/account/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Change User Password
func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.ChangePasswordRes, err error) {
	l.Logger.Info("ChangePassword request: ", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
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
			return &types.ChangePasswordRes{
				Result: types.Result{
					Code:    common.USER_NOT_FOUND_CODE,
					Message: common.USER_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	// Check if the old password is correct
	if !utils.ConfirmPassword(req.OldPassword, userModel.PasswordHash) {
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESS,
			},
		}, nil
	}

	// Change password
	hashpw, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	userModel.PasswordHash = hashpw
	userModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}

	err = l.svcCtx.UserModel.Update(l.ctx, userModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.ChangePasswordRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
