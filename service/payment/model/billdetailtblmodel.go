package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillDetailTblModel = (*customBillDetailTblModel)(nil)

type (
	// BillDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillDetailTblModel.
	BillDetailTblModel interface {
		billDetailTblModel
		withSession(session sqlx.Session) BillDetailTblModel
	}

	customBillDetailTblModel struct {
		*defaultBillDetailTblModel
	}
)

// NewBillDetailTblModel returns a model for the database table.
func NewBillDetailTblModel(conn sqlx.SqlConn) BillDetailTblModel {
	return &customBillDetailTblModel{
		defaultBillDetailTblModel: newBillDetailTblModel(conn),
	}
}

func (m *customBillDetailTblModel) withSession(session sqlx.Session) BillDetailTblModel {
	return NewBillDetailTblModel(sqlx.NewSqlConnFromSession(session))
}
