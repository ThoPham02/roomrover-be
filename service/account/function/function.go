package function

import "roomrover/service/account/model"

type AccountFunction interface {
	GetUserByID(userID int64) (user *model.UserTbl, err error)
	GetUsersByIDs(userIDs []int64) (users []*model.UserTbl, err error)
	UpdateUser(user *model.UserTbl) error
}
