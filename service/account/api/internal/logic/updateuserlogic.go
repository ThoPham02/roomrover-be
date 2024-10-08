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

	if req.Password != "" {
		if len(req.Password) < 6 {
			return &types.UpdateUserRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		hashedPassword, _ := utils.HashPassword(req.Password)
		userModel.PasswordHash = hashedPassword
	}

	userModel.AvatarUrl = sql.NullString{String: req.AvatarUrl, Valid: true}
	userModel.FullName = sql.NullString{String: req.FullName, Valid: true}
	userModel.Address = sql.NullString{String: req.Address, Valid: true}
	userModel.Birthday = sql.NullInt64{Int64: req.Birthday, Valid: true}
	userModel.Gender = sql.NullInt64{Int64: req.Gender, Valid: true}
	userModel.CCCDNumber = sql.NullString{String: req.CccdNumber, Valid: true}
	userModel.CCCDDate = sql.NullInt64{Int64: req.CccdDate, Valid: true}
	userModel.CCCDAddress = sql.NullString{String: req.CccdAddress, Valid: true}
	userModel.AvatarUrl = sql.NullString{String: req.AvatarUrl, Valid: true}
	userModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}

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
		UserID:      userID,
		Phone:       userModel.Phone,
		Role:        userModel.Role.Int64,
		Status:      userModel.Status,
		Address:     userModel.Address.String,
		FullName:    userModel.FullName.String,
		AvatarUrl:   userModel.AvatarUrl.String,
		Birthday:    userModel.Birthday.Int64,
		Gender:      userModel.Gender.Int64,
		CccdNumber:  userModel.CCCDNumber.String,
		CccdDate:    userModel.CCCDDate.Int64,
		CccdAddress: userModel.CCCDAddress.String,
		CreatedAt:   userModel.CreatedAt.Int64,
		UpdatedAt:   userModel.UpdatedAt.Int64,
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
