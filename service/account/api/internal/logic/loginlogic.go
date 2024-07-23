package logic

import (
	"context"
	"database/sql"
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
	l.Logger.Info("Login Success", req)

	var userModel *model.UsersTbl
	var profileModel *model.ProfilesTbl
	var token string
	var currentTime int64 = time.Now().Unix()
	var user types.User
	var profile types.Profile

	userModel, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel == nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.USER_NOT_FOUND_CODE,
				Message: common.USER_NOT_FOUND_MESS,
			},
		}, nil
	}

	var checkPassword bool = utils.ConfirmPassword(req.Password, userModel.PasswordHash)
	if !checkPassword {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.INVALID_PASSWORD_CODE,
				Message: common.INVALID_PASSWORD_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:    userModel.UserId,
		ProfileID: userModel.ProfileId.Int64,
		UserName:  userModel.Username,
		Email:     userModel.Email,
	}

	profileModel, err = l.svcCtx.ProfileModel.FindOne(l.ctx, userModel.ProfileId.Int64)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if profileModel != nil {
		profile = types.Profile{
			ProfileID: profileModel.ProfileId,
			FullName:  profileModel.Fullname.String,
			Dob:       profileModel.Dob.Int64,
			AvatarUrl: profileModel.AvatarUrl.String,
			Address:   profileModel.Address.String,
			Phone:     profileModel.Phone.String,
			CreatedAt: profileModel.CreatedAt.Int64,
			UpdatedAt: profileModel.UpdatedAt.Int64,
			CreatedBy: profileModel.CreatedBy.Int64,
			UpdatedBy: profileModel.UpdatedBy.Int64,
		}
	}

	token, err = utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, currentTime, l.svcCtx.Config.Auth.AccessExpire, userModel.UserId, user)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("Login Success", user)
	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token:   token,
		User:    user,
		Profile: profile,
	}, nil
}
