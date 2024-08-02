package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User Login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	l.Logger.Infof("Login request: %v", req)

	var currentTime = common.GetCurrentTime()
	var user types.User

	// Check if the user exists
	userModel, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
	if err != nil {
		l.Logger.Error(err)
		if err == sql.ErrNoRows {
			return &types.LoginRes{
				Result: types.Result{
					Code:    common.USER_NOT_FOUND_CODE,
					Message: common.USER_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:    userModel.UserId,
		Phone:     userModel.Phone,
		FullName:  userModel.FullName.String,
		Birthday:  userModel.Birthday.Int64,
		AvatarUrl: userModel.AvatarUrl.String,
		Address:   userModel.Address.String,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	// Check if the password is correct
	if !utils.ConfirmPassword(req.Password, userModel.PasswordHash) {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESS,
			},
		}, nil
	}

	// Generate token
	token, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, currentTime, l.svcCtx.Config.Auth.AccessExpire, userModel.UserId, user)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("Login success")
	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: token,
		User:  user,
	}, nil
}
