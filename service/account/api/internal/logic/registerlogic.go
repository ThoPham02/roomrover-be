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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Register New User
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	l.Logger.Info("Register request: ", req)

	var currentTime = common.GetCurrentTime()
	var token string
	var user types.User
	var iat = time.Now().Unix()
	var accessSecret = l.svcCtx.Config.Auth.AccessSecret
	var accessExpire = l.svcCtx.Config.Auth.AccessExpire

	// Check role
	if req.UserRole != common.USER_ROLE_RENTER && req.UserRole != common.USER_ROLE_LESSOR {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	// Check if the user exists
	userModel, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.PhoneNumber)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if userModel != nil {
		if userModel.Status == common.USER_ACTIVE {
			return &types.RegisterRes{
				Result: types.Result{
					Code:    common.USER_ALREADY_EXISTS_CODE,
					Message: common.USER_ALREADY_EXISTS_MESS,
				},
			}, nil
		}

		// Update user info
		userModel.FullName = sql.NullString{Valid: true, String: req.FullName}
		userModel.CCCDNumber = sql.NullString{Valid: true, String: req.Cccd}
		userModel.CCCDDate = sql.NullInt64{Valid: true, Int64: req.IssueDate}
		userModel.CCCDAddress = sql.NullString{Valid: true, String: req.IssuePlace}
		userModel.Role = sql.NullInt64{Valid: true, Int64: req.UserRole}
		userModel.Status = common.USER_ACTIVE
		userModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: currentTime}
	} else {
		userModel = &model.UserTbl{
			Id:           l.svcCtx.ObjSync.GenServiceObjID(),
			Phone:        req.PhoneNumber,
			PasswordHash: "",
			FullName:     sql.NullString{Valid: true, String: req.FullName},
			CCCDNumber:   sql.NullString{Valid: true, String: req.Cccd},
			CCCDDate:     sql.NullInt64{Valid: true, Int64: req.IssueDate},
			CCCDAddress:  sql.NullString{Valid: true, String: req.IssuePlace},
			Role:         sql.NullInt64{Valid: true, Int64: req.UserRole},
			Status:       common.USER_ACTIVE,
			CreatedAt:    sql.NullInt64{Valid: true, Int64: currentTime},
			UpdatedAt:    sql.NullInt64{Valid: true, Int64: currentTime},
		}
	}

	// Register new user
	hashpw, err := utils.HashPassword(req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	userModel.PasswordHash = hashpw
	err = l.svcCtx.UserModel.Delete(l.ctx, userModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, userModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:      userModel.Id,
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

	// Generate token
	token, err = utils.GetJwtToken(accessSecret, iat, accessExpire, userModel.Id, user)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("Register success")
	return &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: token,
		User:  user,
	}, nil
}
