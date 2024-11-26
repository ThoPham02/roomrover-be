package model

import (
	"context"
	"fmt"
	"roomrover/common"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContactTblModel = (*customContactTblModel)(nil)

type (
	// ContactTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactTblModel.
	ContactTblModel interface {
		contactTblModel
		FindMultiByUser(ctx context.Context, renterID, lessorID, from, to, status, limit, offset int64) ([]*ContactTbl, error)
		CountByUser(ctx context.Context, renterID, lessorID, from, to, status int64) (int, error)
		GetCurrentContact(ctx context.Context, userID int64) ([]*ContactTbl, error)
	}

	customContactTblModel struct {
		*defaultContactTblModel
	}
)

// NewContactTblModel returns a model for the database table.
func NewContactTblModel(conn sqlx.SqlConn) ContactTblModel {
	return &customContactTblModel{
		defaultContactTblModel: newContactTblModel(conn),
	}
}

func (m *customContactTblModel) FindMultiByUser(ctx context.Context, renterID, lessorID, from, to, status, limit, offset int64) ([]*ContactTbl, error) {
	query := fmt.Sprintf("select %s from %s where 1 = 1", contactTblRows, m.table)
	var resp []*ContactTbl
	var vals []interface{}

	if renterID != 0 {
		query += " and `renter_id` = ?"
		vals = append(vals, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		vals = append(vals, lessorID)
	}
	if from != 0 {
		query += " and `datetime` >= ?"
		vals = append(vals, from)
	}
	if to != 0 {
		query += " and `datetime` <= ?"
		vals = append(vals, to)
	}
	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
	}
	query += " order by `id` DESC"
	if limit != 0 {
		query += " limit ? offset ? "
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

func (m *customContactTblModel) CountByUser(ctx context.Context, renterID, lessorID, from, to, status int64) (int, error) {
	query := fmt.Sprintf("select count(*) from %s where 1 = 1", m.table)
	var resp int
	var vals []interface{}

	if renterID != 0 {
		query += " and `renter_id` = ?"
		vals = append(vals, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		vals = append(vals, lessorID)
	}
	if from != 0 {
		query += " and `datetime` >= ?"
		vals = append(vals, from)
	}
	if to != 0 {
		query += " and `datetime` <= ?"
		vals = append(vals, to)
	}
	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
	}

	err := m.conn.QueryRowCtx(ctx, &resp, query, vals...)
	return resp, err
}

func (m *customContactTblModel) GetCurrentContact(ctx context.Context, userID int64) ([]*ContactTbl, error) {
	query := fmt.Sprintf("select %s from %s where `lessor_id` = ? and `datetime` between ? and ?", contactTblRows, m.table)
	var resp []*ContactTbl
	var currentDay int64 = common.GetCurrentTime()/86400000*86400000 - 7*3600*1000

	err := m.conn.QueryRowsCtx(ctx, &resp, query, userID, currentDay, currentDay+86400000)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
