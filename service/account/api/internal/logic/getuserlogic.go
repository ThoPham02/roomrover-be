package logic

import (
	"context"
	"database/sql"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get User Info
func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRes, err error) {
	l.Logger.Info("GetUser", req)

	var userID int64
	var user types.User

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetUserRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		l.Logger.Error(err)
		if err == sql.ErrNoRows {
			return &types.GetUserRes{
				Result: types.Result{
					Code:    common.USER_NOT_FOUND_CODE,
					Message: common.USER_NOT_FOUND_MESS,
				},
			}, nil
		}
		return &types.GetUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user = types.User{
		UserID:    userID,
		Phone:     userModel.Phone,
		FullName:  userModel.FullName.String,
		Birthday:  userModel.Birthday.Int64,
		AvatarUrl: userModel.AvatarUrl.String,
		Address:   userModel.Address.String,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	l.Logger.Info("GetUser success")
	return &types.GetUserRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		User: user,
	}, nil
}
