package function

import "roomrover/service/account/model"

type AccountFunction interface {
	GetUserByID(userID int64) (user *model.UserTbl, err error)
	GetUsersByIDs(userIDs []int64) (users []*model.UserTbl, err error)
	UpdateUser(user *model.UserTbl) error
	FindUserByPhone(phone string) (user *model.UserTbl, err error)
	CreateInactivatedUser(userID int64, phone string, fullName, cccdNumber string, cccdDate int64, cccdAddress string) error
}
