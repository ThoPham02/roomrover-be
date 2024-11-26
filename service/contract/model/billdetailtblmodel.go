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
		CountQuantityByBillAndDetailID(ctx context.Context, billID, detailID int64) (int64, error)
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
	err := m.conn.QueryRowsCtx(ctx, &resp, query, billID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customBillDetailTblModel) CountQuantityByBillAndDetailID(ctx context.Context, billID, detailID int64) (int64, error) {
	query := fmt.Sprintf("select IFNULL(SUM(quantity), 0) from %s where `bill_id` = ? and `payment_detail_id` = ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, billID, detailID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}
