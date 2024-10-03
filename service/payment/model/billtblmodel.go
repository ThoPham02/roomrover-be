package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillTblModel = (*customBillTblModel)(nil)

type (
	// BillTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillTblModel.
	BillTblModel interface {
		billTblModel
		withSession(session sqlx.Session) BillTblModel
	}

	customBillTblModel struct {
		*defaultBillTblModel
	}
)

// NewBillTblModel returns a model for the database table.
func NewBillTblModel(conn sqlx.SqlConn) BillTblModel {
	return &customBillTblModel{
		defaultBillTblModel: newBillTblModel(conn),
	}
}

func (m *customBillTblModel) withSession(session sqlx.Session) BillTblModel {
	return NewBillTblModel(sqlx.NewSqlConnFromSession(session))
}
