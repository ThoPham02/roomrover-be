package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTblModel = (*customUserTblModel)(nil)

type (
	// UserTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTblModel.
	UserTblModel interface {
		userTblModel
		withSession(session sqlx.Session) UserTblModel
		FindOneByPhone(ctx context.Context, phone string) (*UserTbl, error)
		FindByIDs(ctx context.Context, userIDs []int64) ([]*UserTbl, error)
		FilterUser(ctx context.Context, phone string, limit, offset int64) ([]*UserTbl, error)
		CountUser(ctx context.Context, phone string) (int, error)
	}

	customUserTblModel struct {
		*defaultUserTblModel
	}
)

// NewUserTblModel returns a model for the database table.
func NewUserTblModel(conn sqlx.SqlConn) UserTblModel {
	return &customUserTblModel{
		defaultUserTblModel: newUserTblModel(conn),
	}
}

func (m *customUserTblModel) withSession(session sqlx.Session) UserTblModel {
	return NewUserTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserTblModel) FindOneByPhone(ctx context.Context, phone string) (*UserTbl, error) {
	var resp UserTbl
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userTblRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUserTblModel) FindByIDs(ctx context.Context, userIDs []int64) ([]*UserTbl, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	query := fmt.Sprintf("select %s from %s where `id` in (", userTblRows, m.table)
	var values []interface{}
	for _, id := range userIDs {
		query += "?,"
		values = append(values, id)
	}
	query = query[:len(query)-1] + ")"
	var resp []*UserTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUserTblModel) FilterUser(ctx context.Context, phone string, limit, offset int64) ([]*UserTbl, error) {
	query := fmt.Sprintf("select %s from %s where `phone` like ?", userTblRows, m.table)
	var resp []*UserTbl
	var vals []interface{}
	vals = append(vals, "%"+phone+"%")
	if limit > 0 {
        query += " limit ? offset ?"
        vals = append(vals, limit, offset)
    }

	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUserTblModel) CountUser(ctx context.Context, phone string) (int, error) {
	query := fmt.Sprintf("select count(*) from %s where `phone` like ?", m.table)
    var resp int
    var vals []interface{}
    vals = append(vals, "%"+phone+"%")

    err := m.conn.QueryRowCtx(ctx, &resp, query, vals...)
    switch err {
    case nil:
        return resp, nil
    case sqlc.ErrNotFound:
        return 0, nil
    default:
        return 0, err
    }
}