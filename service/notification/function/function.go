package function

import (
	"roomrover/service/notification/model"
)

type NotificationFunction interface {
	CreateNotification(noti *model.NotificationTbl) error
}
