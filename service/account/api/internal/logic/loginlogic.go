package logic

import (
	"context"
	"time"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"
	"roomrover/service/account/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	var userModel *model.Users
	var token string
	var currentTime int64 = time.Now().UnixMilli()
	var user types.User

	userModel, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESSAGE,
			},
		}, nil
	}
	if userModel == nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.USER_NOT_FOUND_CODE,
				Message: common.USER_NOT_FOUND_MESSAGE,
			},
		}, nil
	}

	var checkPassword bool = utils.ConfirmPassword(req.Password, userModel.PasswordHash)
	if !checkPassword {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESSAGE,
			},
		}, nil
	}

	user = types.User{
		UserID:    userModel.UserId,
		ProfileID: userModel.ProfileId.Int64,
		UserName:  userModel.Username,
		Email:     userModel.Email,
	}

	token, err = utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, currentTime, l.svcCtx.Config.Auth.AccessExpire, userModel.UserId, user)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESSAGE,
			},
		}, nil
	}

	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESSAGE,
		},
		Token: token,
		User:  user,
	}, nil
}
