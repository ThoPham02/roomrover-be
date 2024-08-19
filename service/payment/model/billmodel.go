package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillModel = (*customBillModel)(nil)

type (
	// BillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillModel.
	BillModel interface {
		billModel
	}

	customBillModel struct {
		*defaultBillModel
	}
)

// NewBillModel returns a model for the database table.
func NewBillModel(conn sqlx.SqlConn) BillModel {
	return &customBillModel{
		defaultBillModel: newBillModel(conn),
	}
}
