package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContactTblModel = (*customContactTblModel)(nil)

type (
	// ContactTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactTblModel.
	ContactTblModel interface {
		contactTblModel
		FindMultiByUser(ctx context.Context, renterID, lessorID int64) ([]*ContactTbl, error)
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

func (m *customContactTblModel) FindMultiByUser(ctx context.Context, renterID, lessorID int64) ([]*ContactTbl, error) {
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
	query += "order by `id` DESC"

	err := m.conn.QueryRowCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
