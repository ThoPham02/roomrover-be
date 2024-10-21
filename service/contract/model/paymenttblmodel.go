package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentTblModel = (*customPaymentTblModel)(nil)

type (
	// PaymentTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentTblModel.
	PaymentTblModel interface {
		paymentTblModel
		FindByContractID(ctx context.Context, contractID int64) (*PaymentTbl, error)
		DeleteByContractID(ctx context.Context, contractID int64) error
		FilterPaymentByTime(ctx context.Context, time int64) ([]*PaymentTbl, error)
	}

	customPaymentTblModel struct {
		*defaultPaymentTblModel
	}
)

// NewPaymentTblModel returns a model for the database table.
func NewPaymentTblModel(conn sqlx.SqlConn) PaymentTblModel {
	return &customPaymentTblModel{
		defaultPaymentTblModel: newPaymentTblModel(conn),
	}
}

func (m *customPaymentTblModel) FindByContractID(ctx context.Context, contractID int64) (*PaymentTbl, error) {
	query := fmt.Sprintf("select %s from %s where `contract_id` = ? limit 1", paymentTblRows, m.table)
	var resp PaymentTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, contractID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPaymentTblModel) DeleteByContractID(ctx context.Context, contractID int64) error {
	query := fmt.Sprintf("delete from %s where `contract_id` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, contractID)
	return err
}

func (m *customPaymentTblModel) FilterPaymentByTime(ctx context.Context, time int64) ([]*PaymentTbl, error) {
	var start = time - 86400000/2
	var end = time + 86400000/2

	query := fmt.Sprintf("select %s from %s where `next_bill` between ? and ? ", paymentTblRows, m.table)
	var resp []*PaymentTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, start, end)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
