package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"
	"roomrover/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req *types.GetProfileReq) (resp *types.GetProfileRes, err error) {
	l.Logger.Info("GetProfileLogic: ", req)

	var userID int64

	var userModel *model.UsersTbl
	var profileModel *model.ProfilesTbl

	var user types.User
	var profile types.Profile

	userID, err = utils.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetProfileRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	userModel, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.GetProfileRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel == nil {
		return &types.GetProfileRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	user = types.User{
		UserID:    userID,
		ProfileID: userModel.ProfileId.Int64,
		UserName:  userModel.Username,
		Email:     userModel.Email,
	}

	profileModel, err = l.svcCtx.ProfileModel.FindOne(l.ctx, userModel.ProfileId.Int64)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.GetProfileRes{
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
		}
	}

	l.Logger.Info("GetProfileLogic Success: ", userID)
	return &types.GetProfileRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Profile: profile,
		User:    user,
	}, nil
}
