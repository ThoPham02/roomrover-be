package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaymentUsedTblModel = (*customPaymentUsedTblModel)(nil)

type (
	// PaymentUsedTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentUsedTblModel.
	PaymentUsedTblModel interface {
		paymentUsedTblModel
		withSession(session sqlx.Session) PaymentUsedTblModel
	}

	customPaymentUsedTblModel struct {
		*defaultPaymentUsedTblModel
	}
)

// NewPaymentUsedTblModel returns a model for the database table.
func NewPaymentUsedTblModel(conn sqlx.SqlConn) PaymentUsedTblModel {
	return &customPaymentUsedTblModel{
		defaultPaymentUsedTblModel: newPaymentUsedTblModel(conn),
	}
}

func (m *customPaymentUsedTblModel) withSession(session sqlx.Session) PaymentUsedTblModel {
	return NewPaymentUsedTblModel(sqlx.NewSqlConnFromSession(session))
}
