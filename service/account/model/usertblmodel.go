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
