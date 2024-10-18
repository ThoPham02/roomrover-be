package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillTblModel = (*customBillTblModel)(nil)

type FilterCondition struct {
}

type (
	// BillTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillTblModel.
	BillTblModel interface {
		billTblModel
		CountByCondition(ctx context.Context, condition FilterCondition) (int64, error)
		FilterBillByCondition(ctx context.Context, condition FilterCondition) ([]*BillTbl, error)
	}

	customBillTblModel struct {
		*defaultBillTblModel
	}
)

// NewBillTblModel returns a model for the database table.
func NewBillTblModel(conn sqlx.SqlConn) BillTblModel {
	return &customBillTblModel{
		defaultBillTblModel: newBillTblModel(conn),
	}
}

func (m *customBillTblModel) CountByCondition(ctx context.Context, condition FilterCondition) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where 1=1 ", m.table)
	var args []interface{}
	var count int64

	err := m.conn.QueryRowCtx(ctx, &count, query, args...)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customBillTblModel) FilterBillByCondition(ctx context.Context, condition FilterCondition) ([]*BillTbl, error) {
	query := fmt.Sprintf("select %s from %s where 1=1 ", billTblRows, m.table)
	var args []interface{}
	var resp []*BillTbl



	err := m.conn.QueryRowsCtx(ctx, resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
