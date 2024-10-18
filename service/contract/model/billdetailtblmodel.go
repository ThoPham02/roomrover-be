package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillDetailTblModel = (*customBillDetailTblModel)(nil)

type (
	// BillDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillDetailTblModel.
	BillDetailTblModel interface {
		billDetailTblModel
		GetDetailByBillID(ctx context.Context, billID int64) ([]*BillDetailTbl, error)
	}

	customBillDetailTblModel struct {
		*defaultBillDetailTblModel
	}
)

// NewBillDetailTblModel returns a model for the database table.
func NewBillDetailTblModel(conn sqlx.SqlConn) BillDetailTblModel {
	return &customBillDetailTblModel{
		defaultBillDetailTblModel: newBillDetailTblModel(conn),
	}
}

func (m *customBillDetailTblModel) GetDetailByBillID(ctx context.Context, billID int64) ([]*BillDetailTbl, error) {
	query := fmt.Sprintf("select %s from %s where `bill_id` = ?", billDetailTblRows, m.table)
	var resp []*BillDetailTbl
	err := m.conn.QueryRowsCtx(ctx, resp, query, billID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
