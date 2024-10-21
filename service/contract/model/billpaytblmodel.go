package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillPayTblModel = (*customBillPayTblModel)(nil)

type (
	// BillPayTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillPayTblModel.
	BillPayTblModel interface {
		billPayTblModel
	}

	customBillPayTblModel struct {
		*defaultBillPayTblModel
	}
)

// NewBillPayTblModel returns a model for the database table.
func NewBillPayTblModel(conn sqlx.SqlConn) BillPayTblModel {
	return &customBillPayTblModel{
		defaultBillPayTblModel: newBillPayTblModel(conn),
	}
}
