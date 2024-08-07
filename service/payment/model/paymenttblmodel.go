package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaymentTblModel = (*customPaymentTblModel)(nil)

type (
	// PaymentTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentTblModel.
	PaymentTblModel interface {
		paymentTblModel
		withSession(session sqlx.Session) PaymentTblModel
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

func (m *customPaymentTblModel) withSession(session sqlx.Session) PaymentTblModel {
	return NewPaymentTblModel(sqlx.NewSqlConnFromSession(session))
}
