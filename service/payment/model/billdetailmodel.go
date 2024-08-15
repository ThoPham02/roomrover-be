package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillDetailModel = (*customBillDetailModel)(nil)

type (
	// BillDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillDetailModel.
	BillDetailModel interface {
		billDetailModel
	}

	customBillDetailModel struct {
		*defaultBillDetailModel
	}
)

// NewBillDetailModel returns a model for the database table.
func NewBillDetailModel(conn sqlx.SqlConn) BillDetailModel {
	return &customBillDetailModel{
		defaultBillDetailModel: newBillDetailModel(conn),
	}
}
