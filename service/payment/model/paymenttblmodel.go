package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentTblModel = (*customPaymentTblModel)(nil)

type (
	// PaymentTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentTblModel.
	PaymentTblModel interface {
		paymentTblModel
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

func (m *customPaymentTblModel) FilterPaymentByTime(ctx context.Context, time int64) ([]*PaymentTbl, error) {
	var startTime = time - 12*60*60*1000
	var endTime = time + 12*60*60*1000
	query := fmt.Sprintf("select %s from %s where `next_bill` between ? and ?", paymentTblRows, m.table)
	var resp []*PaymentTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, startTime, endTime)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
