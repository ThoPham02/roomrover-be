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
	userModel, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
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
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.USER_ALREADY_EXISTS_CODE,
				Message: common.USER_ALREADY_EXISTS_MESS,
			},
		}, nil
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

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		UserId:       l.svcCtx.ObjSync.GenServiceObjID(),
		Phone:        req.Phone,
		PasswordHash: hashpw,
		Role:         sql.NullInt64{Valid: true, Int64: req.UserRole},
		Status:       common.USER_ACTIVE,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("Register success")
	return &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
