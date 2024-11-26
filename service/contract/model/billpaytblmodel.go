package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillPayTblModel = (*customBillPayTblModel)(nil)

type (
	// BillPayTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillPayTblModel.
	BillPayTblModel interface {
		billPayTblModel
		FindOneByTransID(ctx context.Context, appTransID string) (*BillPayTbl, error)
		GetPayByBillID(ctx context.Context, billID int64) ([]*BillPayTbl, error)
		FindByContractID(ctx context.Context, contractID int64) ([]*BillPayTbl, error)
	}

	customBillPayTblModel struct {
		*defaultBillPayTblModel
	}
)

// NewBillPayTblModel returns a model for the database table.
func NewBillPayTblModel(conn sqlx.SqlConn) BillPayTblModel {
	return &customBillPayTblModel{
		defaultBillPayTblModel: newBillPayTblModel(conn),
	}
}

func (m *customBillPayTblModel) FindOneByTransID(ctx context.Context, appTransID string) (*BillPayTbl, error) {
	query := fmt.Sprintf("select %s from %s where `trans_id` = ? limit 1", billPayTblRows, m.table)
	var resp BillPayTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, appTransID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customBillPayTblModel) GetPayByBillID(ctx context.Context, billID int64) ([]*BillPayTbl, error) {
	query := fmt.Sprintf("select %s from %s where `bill_id` = ?", billPayTblRows, m.table)
	var resp []*BillPayTbl
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

func (m *customBillPayTblModel) FindByContractID(ctx context.Context, contractID int64) ([]*BillPayTbl, error) {
	query := fmt.Sprintf(`
	SELECT 
		bill_pay_tbl.id,
		bill_pay_tbl.bill_id, 
		bill_pay_tbl.user_id,
		bill_pay_tbl.amount,
		bill_pay_tbl.pay_date,
		bill_pay_tbl.status,
		bill_pay_tbl.trans_id,
		bill_pay_tbl.type,
		bill_pay_tbl.url
	FROM bill_pay_tbl 
	JOIN bill_tbl ON bill_pay_tbl.bill_id = bill_tbl.id
	JOIN payment_tbl ON bill_tbl.payment_id = payment_tbl.id
	JOIN contract_tbl ON payment_tbl.contract_id = contract_tbl.id
	WHERE contract_tbl.id = ?`)
	var resp []*BillPayTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, contractID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
