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

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	l.Logger.Info("RegisterLogic: ", req)

	var userModel *model.UsersTbl
	var user types.User

	var token string
	var currentTime int64 = time.Now().Unix()
	var hashPW string

	// Check if the request is valid
	checkUserName := req.UserName == ""
	checkPassword := req.Password == ""
	checkRole := req.UserRole != common.USER_ROLE_RENTER && req.UserRole != common.USER_ROLE_LESSOR
	if checkUserName || checkPassword || checkRole {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	// Check if the user already exists
	userModel, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil && err != sql.ErrNoRows {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if userModel != nil {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.USER_ALREADY_EXISTS_CODE,
				Message: common.USER_ALREADY_EXISTS_MESS,
			},
		}, nil
	}

	// Hash the password
	hashPW, err = utils.HashPassword(req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	userModel = &model.UsersTbl{
		UserId:       l.svcCtx.ObjSync.GenServiceObjID(),
		ProfileId:    sql.NullInt64{},
		Username:     req.UserName,
		PasswordHash: hashPW,
		Email:        req.Email,
		Role:         sql.NullInt64{Valid: true, Int64: req.UserRole},
	}
	user = types.User{
		UserID:    userModel.UserId,
		ProfileID: 0,
		UserName:  userModel.Username,
		Email:     userModel.Email,
	}

	// Register the user
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

	token, err = utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, currentTime, l.svcCtx.Config.Auth.AccessExpire, userModel.UserId, user)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	resp = &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: token,
		User:  user,
	}

	l.Logger.Info("RegisterLogic success: ", resp)
	return resp, nil
}
