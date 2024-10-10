package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ServiceTblModel = (*customServiceTblModel)(nil)

type (
	// ServiceTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customServiceTblModel.
	ServiceTblModel interface {
		serviceTblModel
		withSession(session sqlx.Session) ServiceTblModel
		FindByHouseID(ctx context.Context, houseID int64) ([]*ServiceTbl, error)
		DeleteByHouseID(ctx context.Context, houseID int64) error
		FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*ServiceTbl, error)
	}

	customServiceTblModel struct {
		*defaultServiceTblModel
	}
)

// NewServiceTblModel returns a model for the database table.
func NewServiceTblModel(conn sqlx.SqlConn) ServiceTblModel {
	return &customServiceTblModel{
		defaultServiceTblModel: newServiceTblModel(conn),
	}
}

func (m *customServiceTblModel) withSession(session sqlx.Session) ServiceTblModel {
	return NewServiceTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customServiceTblModel) FindByHouseID(ctx context.Context, houseID int64) ([]*ServiceTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` = ?", serviceTblRows, m.table)
	var resp []*ServiceTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, houseID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customServiceTblModel) DeleteByHouseID(ctx context.Context, houseID int64) error {
	query := fmt.Sprintf("delete from %s where `house_id` =?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, houseID)
	return err
}

func (m *customServiceTblModel) FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*ServiceTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` in (", serviceTblRows, m.table)
	var resp []*ServiceTbl
	var vals []interface{}
	for _, id := range houseIDs {
		query += "?,"
		vals = append(vals, id)
	}
	query = query[:len(query)-1] + ")"
	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
