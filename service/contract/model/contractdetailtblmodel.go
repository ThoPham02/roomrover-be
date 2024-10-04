package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContractDetailTblModel = (*customContractDetailTblModel)(nil)

type (
	// ContractDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractDetailTblModel.
	ContractDetailTblModel interface {
		contractDetailTblModel
		withSession(session sqlx.Session) ContractDetailTblModel
		InsertMulti(ctx context.Context, data []*ContractDetailTbl) error
		GetContractDetailByContractID(ctx context.Context, contractID int64) ([]*ContractDetailTbl, error)
	}

	customContractDetailTblModel struct {
		*defaultContractDetailTblModel
	}
)

// NewContractDetailTblModel returns a model for the database table.
func NewContractDetailTblModel(conn sqlx.SqlConn) ContractDetailTblModel {
	return &customContractDetailTblModel{
		defaultContractDetailTblModel: newContractDetailTblModel(conn),
	}
}

func (m *customContractDetailTblModel) withSession(session sqlx.Session) ContractDetailTblModel {
	return NewContractDetailTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customContractDetailTblModel) InsertMulti(ctx context.Context, data []*ContractDetailTbl) error {
	if len(data) == 0 {
		return nil
	}

	var values []interface{}
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, contractDetailTblRowsExpectAutoSet)
	for _, row := range data {
		values = append(values, row.Id, row.ContractId, row.Name, row.Type, row.Price)
		query += "(?, ?, ?, ?, ?),"
	}
	query = query[:len(query)-1]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}

func (m *customContractDetailTblModel) GetContractDetailByContractID(ctx context.Context, contractID int64) ([]*ContractDetailTbl, error) {
	query := fmt.Sprintf("select %s from %s where `contract_id` = ?", contractDetailTblRows, m.table)
	var resp []*ContractDetailTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, contractID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
