package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NotificationTblModel = (*customNotificationTblModel)(nil)

type (
	// NotificationTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationTblModel.
	NotificationTblModel interface {
		notificationTblModel
		DeleteNotiByRefID(ctx context.Context, refID int64) error
		GetNotisByReceiver(ctx context.Context, receiverID int64, limit, offset int64) ([]*NotificationTbl, error)
		CountNotisByReceiver(ctx context.Context, receiverID int64) (int, error)
		MarkRead(ctx context.Context, id int64) error
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

func (m *defaultNotificationTblModel) DeleteNotiByRefID(ctx context.Context, refID int64) error {
	query := fmt.Sprintf("delete from %s where `ref_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, refID)
	return err
}

func (m *defaultNotificationTblModel) GetNotisByReceiver(ctx context.Context, receiverID int64, limit, offset int64) ([]*NotificationTbl, error) {
	query := fmt.Sprintf("select %s from %s where `receiver` = ? ", notificationTblRows, m.table)
	var resp []*NotificationTbl
	query += " ORDER BY `created_at` DESC"
	if limit > 0 {
		query += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}

	err := m.conn.QueryRowsCtx(ctx, &resp, query, receiverID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNotificationTblModel) CountNotisByReceiver(ctx context.Context, receiverID int64) (int, error) {
	query := fmt.Sprintf("select count(*) from %s where `receiver` = ? ", m.table)
	var resp int
	err := m.conn.QueryRowCtx(ctx, &resp, query, receiverID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *defaultNotificationTblModel) MarkRead(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `unread` = 2 where `id` = ?", m.table)
    _, err := m.conn.ExecCtx(ctx, query, id)
    return err
}