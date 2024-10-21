package api

import (
	"context"
	"roomrover/service/account/function"
	"roomrover/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.AccountFunction = (*AccountFunction)(nil)

type AccountFunction struct {
	function.AccountFunction
	logx.Logger
	AccountService *AccountService
}

func NewAccountFunction(svc *AccountService) *AccountFunction {
	ctx := context.Background()

	return &AccountFunction{
		Logger:         logx.WithContext(ctx),
		AccountService: svc,
	}
}

func (contractFunc *AccountFunction) Start() error {
	return nil
}

func (af *AccountFunction) GetUserByID(userID int64) (user *model.UserTbl, err error) {
	return af.AccountService.Ctx.UserModel.FindOne(context.Background(), userID)
}

func (af *AccountFunction) GetUsersByIDs(userIDs []int64) (users []*model.UserTbl, err error) {
	return af.AccountService.Ctx.UserModel.FindByIDs(context.Background(), userIDs)
}

func (af *AccountFunction) UpdateUser(user *model.UserTbl) error {
	return af.AccountService.Ctx.UserModel.Update(context.Background(), user)
}
