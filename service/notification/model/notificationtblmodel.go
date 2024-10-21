package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ NotificationTblModel = (*customNotificationTblModel)(nil)

type (
	// NotificationTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationTblModel.
	NotificationTblModel interface {
		notificationTblModel
	}

	customNotificationTblModel struct {
		*defaultNotificationTblModel
	}
)

// NewNotificationTblModel returns a model for the database table.
func NewNotificationTblModel(conn sqlx.SqlConn) NotificationTblModel {
	return &customNotificationTblModel{
		defaultNotificationTblModel: newNotificationTblModel(conn),
	}
}
