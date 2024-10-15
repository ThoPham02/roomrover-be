package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentRenterTblModel = (*customPaymentRenterTblModel)(nil)

type (
	// PaymentRenterTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentRenterTblModel.
	PaymentRenterTblModel interface {
		paymentRenterTblModel
		GetRenterByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentRenterTbl, error)
		DeleteByPaymentID(ctx context.Context, paymentID int64) error
		CountRentersByPaymentID(ctx context.Context, paymentID int64) (int64, error)
	}

	customPaymentRenterTblModel struct {
		*defaultPaymentRenterTblModel
	}
)

// NewPaymentRenterTblModel returns a model for the database table.
func NewPaymentRenterTblModel(conn sqlx.SqlConn) PaymentRenterTblModel {
	return &customPaymentRenterTblModel{
		defaultPaymentRenterTblModel: newPaymentRenterTblModel(conn),
	}
}

func (m *customPaymentRenterTblModel) GetRenterByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentRenterTbl, error) {
	query := fmt.Sprintf("select %s from %s where `payment_id` = ? ", paymentRenterTblRows, m.table)
	var resp []*PaymentRenterTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, paymentID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPaymentRenterTblModel) DeleteByPaymentID(ctx context.Context, paymentID int64) error {
	query := fmt.Sprintf("delete from %s where `payment_id` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, paymentID)
	return err
}

func (m *customPaymentRenterTblModel) CountRentersByPaymentID(ctx context.Context, paymentID int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `payment_id` = ? ", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, paymentID)
	return count, err
}
