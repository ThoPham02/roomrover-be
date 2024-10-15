package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentDetailTblModel = (*customPaymentDetailTblModel)(nil)

type (
	// PaymentDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentDetailTblModel.
	PaymentDetailTblModel interface {
		paymentDetailTblModel
		GetPaymentDetailByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentDetailTbl, error)
	}

	customPaymentDetailTblModel struct {
		*defaultPaymentDetailTblModel
	}
)

// NewPaymentDetailTblModel returns a model for the database table.
func NewPaymentDetailTblModel(conn sqlx.SqlConn) PaymentDetailTblModel {
	return &customPaymentDetailTblModel{
		defaultPaymentDetailTblModel: newPaymentDetailTblModel(conn),
	}
}

func (m *customPaymentDetailTblModel) GetPaymentDetailByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentDetailTbl, error) {
	query := fmt.Sprintf("select %s from %s where `payment_id` = ? ", paymentDetailTblRows, m.table)
	var resp []*PaymentDetailTbl
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
