package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaymentTlModel = (*customPaymentTlModel)(nil)

type (
	// PaymentTlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentTlModel.
	PaymentTlModel interface {
		paymentTlModel
		withSession(session sqlx.Session) PaymentTlModel
	}

	customPaymentTlModel struct {
		*defaultPaymentTlModel
	}
)

// NewPaymentTlModel returns a model for the database table.
func NewPaymentTlModel(conn sqlx.SqlConn) PaymentTlModel {
	return &customPaymentTlModel{
		defaultPaymentTlModel: newPaymentTlModel(conn),
	}
}

func (m *customPaymentTlModel) withSession(session sqlx.Session) PaymentTlModel {
	return NewPaymentTlModel(sqlx.NewSqlConnFromSession(session))
}
