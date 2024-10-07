package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"

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
		if err == model.ErrNotFound {
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

	l.Logger.Info("GetUser success")
	return &types.GetUserRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		User: user,
	}, nil
}
