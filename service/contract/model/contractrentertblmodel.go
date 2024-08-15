package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContractRenterTblModel = (*customContractRenterTblModel)(nil)

type (
	// ContractRenterTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractRenterTblModel.
	ContractRenterTblModel interface {
		contractRenterTblModel
		withSession(session sqlx.Session) ContractRenterTblModel
		InsertMulti(ctx context.Context, data []*ContractRenterTbl) error
	}

	customContractRenterTblModel struct {
		*defaultContractRenterTblModel
	}
)

// NewContractRenterTblModel returns a model for the database table.
func NewContractRenterTblModel(conn sqlx.SqlConn) ContractRenterTblModel {
	return &customContractRenterTblModel{
		defaultContractRenterTblModel: newContractRenterTblModel(conn),
	}
}

func (m *customContractRenterTblModel) withSession(session sqlx.Session) ContractRenterTblModel {
	return NewContractRenterTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customContractRenterTblModel) InsertMulti(ctx context.Context, data []*ContractRenterTbl) error {
	if len(data) == 0 {
		return nil
	}

	var values []interface{}
	query := "insert into %s (%s) values "
	for _, row := range data {
		values = append(values, row.Id, row.ContractId, row.RenterId, row.Type)
		query += "(?, ?, ?, ?),"
	}
	query = query[:len(query)-1]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}	