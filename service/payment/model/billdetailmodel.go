package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillDetailModel = (*customBillDetailModel)(nil)

type (
	// BillDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillDetailModel.
	BillDetailModel interface {
		billDetailModel
		InsertMulti(ctx context.Context, data []*BillDetail) error
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

func (m *customBillDetailModel) InsertMulti(ctx context.Context, data []*BillDetail) error {
	if len(data) == 0 {
		return nil
	}

	var values []interface{}
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, billDetailRowsExpectAutoSet)
	for _, row := range data {
		values = append(values, row.Id, row.BillId, row.ContractServiceId, row.Price, row.Type, row.Quantity)
		query += "(?, ?, ?, ?, ?, ?),"
	}
	query = query[:len(query)-1]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}
