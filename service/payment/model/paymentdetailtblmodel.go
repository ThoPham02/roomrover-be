package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaymentDetailTblModel = (*customPaymentDetailTblModel)(nil)

type (
	// PaymentDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentDetailTblModel.
	PaymentDetailTblModel interface {
		paymentDetailTblModel
		withSession(session sqlx.Session) PaymentDetailTblModel
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

func (m *customPaymentDetailTblModel) withSession(session sqlx.Session) PaymentDetailTblModel {
	return NewPaymentDetailTblModel(sqlx.NewSqlConnFromSession(session))
}
