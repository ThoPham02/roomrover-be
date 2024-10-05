package model

import (
	"context"
	"fmt"

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
		CountRenterByContractID(ctx context.Context, contractID int64) (int64, error)
		GetRenterByContractID(ctx context.Context, contractID int64) ([]*ContractRenterTbl, error)
		DeleteByContractID(ctx context.Context, contractID int64) error
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
		values = append(values, row.Id, row.ContractId, row.UserId)
		query += "(?, ?, ?, ?),"
	}
	query = query[:len(query)-1]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}

func (m *customContractRenterTblModel) CountRenterByContractID(ctx context.Context, contractID int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `contract_id` = ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, contractID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customContractRenterTblModel) GetRenterByContractID(ctx context.Context, contractID int64) ([]*ContractRenterTbl, error) {
	query := fmt.Sprintf("select %s from %s where `contract_id` = ?", contractRenterTblRows, m.table)
	var resp []*ContractRenterTbl
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

func (m *customContractRenterTblModel) DeleteByContractID(ctx context.Context, contractID int64) error {
	query := fmt.Sprintf("delete from %s where `contract_id` =?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, contractID)
	return err
}
