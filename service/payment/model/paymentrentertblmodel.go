package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentRenterTblModel = (*customPaymentRenterTblModel)(nil)

type (
	// PaymentRenterTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentRenterTblModel.
	PaymentRenterTblModel interface {
		paymentRenterTblModel
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

func (m *customPaymentRenterTblModel) CountRentersByPaymentID(ctx context.Context, paymentID int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `payment_id` = ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, paymentID)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
