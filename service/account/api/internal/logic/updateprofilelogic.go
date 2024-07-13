package logic

import (
	"context"
	"database/sql"
	"time"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"
	"roomrover/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProfileLogic) UpdateProfile(req *types.UpdateProfileReq) (resp *types.UpdateProfileRes, err error) {
	l.Logger.Info("UpdateProfileLogic: ", req)

	var userID int64
	var currentTime int64 = time.Now().UnixMilli()

	var userModel *model.UsersTbl
	var profileModel *model.ProfilesTbl

	var user types.User
	var profile types.Profile

	userID, err = utils.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateProfileRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	userModel, err = l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.UpdateProfileRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel == nil {
		return &types.UpdateProfileRes{
			Result: types.Result{
				Code:    common.USER_NOT_FOUND_CODE,
				Message: common.USER_NOT_FOUND_MESS,
			},
		}, nil
	}

	if userModel.ProfileId.Int64 == 0 {
		// create profile and update account
		var profileID int64 = l.svcCtx.ObjSync.GenServiceObjID()

		profileModel = &model.ProfilesTbl{
			ProfileId: profileID,
			Fullname:  sql.NullString{Valid: true, String: req.FullName},
			Dob:       sql.NullInt64{Valid: true, Int64: req.Dob},
			AvatarUrl: sql.NullString{Valid: true, String: req.AvatarUrl},
			Address:   sql.NullString{Valid: true, String: req.Address},
			Phone:     sql.NullString{Valid: true, String: req.Phone},
			CreatedAt: sql.NullInt64{Valid: true, Int64: currentTime},
			CreatedBy: sql.NullInt64{Valid: true, Int64: userID},
			UpdatedAt: sql.NullInt64{Valid: true, Int64: currentTime},
			UpdatedBy: sql.NullInt64{Valid: true, Int64: userID},
		}
		_, err = l.svcCtx.ProfileModel.Insert(l.ctx, profileModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateProfileRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		userModel.ProfileId = sql.NullInt64{Valid: true, Int64: profileID}
		err = l.svcCtx.UserModel.Update(l.ctx, userModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateProfileRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

	} else {
		// update profile
		profileModel, err = l.svcCtx.ProfileModel.FindOne(l.ctx, userModel.ProfileId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateProfileRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		profileModel.Fullname = sql.NullString{Valid: true, String: req.FullName}
		profileModel.Dob = sql.NullInt64{Valid: true, Int64: req.Dob}
		profileModel.AvatarUrl = sql.NullString{Valid: true, String: req.AvatarUrl}
		profileModel.Address = sql.NullString{Valid: true, String: req.Address}
		profileModel.Phone = sql.NullString{Valid: true, String: req.Phone}
		profileModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: currentTime}
		profileModel.UpdatedBy = sql.NullInt64{Valid: true, Int64: userID}

		err = l.svcCtx.ProfileModel.Update(l.ctx, profileModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateProfileRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	user = types.User{
		UserID:    userID,
		ProfileID: userModel.ProfileId.Int64,
		UserName:  userModel.Username,
		Email:     userModel.Email,
	}
	profile = types.Profile{
		UserID:    userID,
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

	l.Logger.Info("UpdateProfileLogic Success: ", userID)
	return &types.UpdateProfileRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Profile: profile,
		User:    user,
	}, nil
}
