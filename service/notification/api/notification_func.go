package api

import (
	"context"
	"roomrover/service/notification/function"
	"roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.NotificationFunction = (*NotificationFunction)(nil)

type NotificationFunction struct {
	function.NotificationFunction
	logx.Logger
	NotificationService *NotificationService
}

func NewNotificationFunction(svc *NotificationService) *NotificationFunction {
	ctx := context.Background()

	return &NotificationFunction{
		Logger:              logx.WithContext(ctx),
		NotificationService: svc,
	}
}

func (notificationFunc *NotificationFunction) Start() error {
	return nil
}

func (notificationFunc *NotificationFunction) CreateNotification(noti *model.NotificationTbl) error {
	return notificationFunc.NotificationService.Ctx.NotificationModel.Update(context.TODO(), noti)
}
