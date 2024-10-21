package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
	"roomrover/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterUserLogic {
	return &FilterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterUserLogic) FilterUser(req *types.FilterUserReq) (resp *types.FilterUserRes, err error) {
	l.Logger.Info("FilterUser: ", req)

	var users []types.User
	var count int

	var userModels []*model.UserTbl

	count, err = l.svcCtx.UserModel.CountUser(l.ctx, req.SearchPhone)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if count == 0 {
		return &types.FilterUserRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}
	userModels, err = l.svcCtx.UserModel.FilterUser(l.ctx, req.SearchPhone, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterUserRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, userModel := range userModels {
		users = append(users, types.User{
			UserID:      userModel.Id,
			Phone:       userModel.Phone,
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
		})
	}

	l.Logger.Info("FilterUser Success: ", resp)
	return &types.FilterUserRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Users: users,
	}, nil

}
