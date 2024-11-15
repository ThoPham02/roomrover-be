package function

import (
	"roomrover/service/notification/model"
)

type NotificationFunction interface {
	CreateNotification(noti *model.NotificationTbl) error
	DeleteNotiByRefID(refID int64) error
}
