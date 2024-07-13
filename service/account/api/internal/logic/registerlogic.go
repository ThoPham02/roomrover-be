package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"

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

	var token string

	var userType types.User

	// Check if the request is valid
	if req.UserName == "" || req.Password == "" {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESSAGE,
			},
		}, nil
	}

	// Check if the user already exists
	if l.svcCtx.UserModel.CheckUserExists(req.UserName) {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.USER_ALREADY_EXISTS_CODE,
				Message: common.USER_ALREADY_EXISTS_MESSAGE,
			},
		}, nil
	}

	// Register the user
	err = l.svcCtx.UserModel.Register(req.UserName, req.Password)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESSAGE,
			},
		}, nil
	}

	resp = &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESSAGE,
		},
		Token: token,
		User:  userType,
	}

	l.Logger.Info("RegisterLogic success: ", resp)
	return
}
