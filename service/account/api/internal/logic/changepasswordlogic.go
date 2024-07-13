package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"
	localUtils "roomrover/service/account/utils"
	"roomrover/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.ChangePasswordRes, err error) {
	l.Logger.Info("ChangePasswordLogic: ", req)

	var userID int64

	var userModel *model.UsersTbl

	userID, err = utils.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	userModel, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if !localUtils.ConfirmPassword(req.OldPassword, userModel.PasswordHash) {
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESS,
			},
		}, nil
	}

	userModel.PasswordHash, err = localUtils.HashPassword(req.NewPassword)
	if err != nil {
		l.Logger.Error(err)
		return &types.ChangePasswordRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESS,
			},
		}, nil
	}

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

	l.Logger.Info("ChangePasswordLogic Success: ", userID)
	return &types.ChangePasswordRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
