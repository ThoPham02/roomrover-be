package api

import (
	"context"
	"database/sql"
	"roomrover/common"
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

func (af *AccountFunction) FindUserByPhone(phone string) (user *model.UserTbl, err error) {
	return af.AccountService.Ctx.UserModel.FindOneByPhone(context.Background(), phone)
}

func (af *AccountFunction) CreateInactivatedUser(userID int64, phone string, fullName, cccdNumber string, cccdDate int64, cccdAddress string) error {
	var current = common.GetCurrentTime()
	_, err := af.AccountService.Ctx.UserModel.Insert(context.Background(), &model.UserTbl{
		Id:          userID,
		Phone:       phone,
		Role:        sql.NullInt64{Valid: true, Int64: common.USER_ROLE_RENTER},
		Status:      common.USER_INACTIVE,
		FullName:    sql.NullString{Valid: true, String: fullName},
		CCCDNumber:  sql.NullString{Valid: true, String: cccdNumber},
		CCCDDate:    sql.NullInt64{Valid: true, Int64: cccdDate},
		CCCDAddress: sql.NullString{Valid: true, String: cccdAddress},
		CreatedAt:   sql.NullInt64{Valid: true, Int64: current},
		UpdatedAt:   sql.NullInt64{Valid: true, Int64: current},
	})
	return err
}
