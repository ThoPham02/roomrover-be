package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContractTblModel = (*customContractTblModel)(nil)

type (
	// ContractTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractTblModel.
	ContractTblModel interface {
		contractTblModel
		withSession(session sqlx.Session) ContractTblModel
		GetContractByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error)
		GetContractByTime(ctx context.Context, time int64) ([]*ContractTbl, error)
	}

	customContractTblModel struct {
		*defaultContractTblModel
	}
)

// NewContractTblModel returns a model for the database table.
func NewContractTblModel(conn sqlx.SqlConn) ContractTblModel {
	return &customContractTblModel{
		defaultContractTblModel: newContractTblModel(conn),
	}
}

func (m *customContractTblModel) withSession(session sqlx.Session) ContractTblModel {
	return NewContractTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customContractTblModel) GetContractByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error) {
	query := fmt.Sprintf("select %s from %s where `room_id` = ? limit 1", contractTblRows, m.table)
	var resp ContractTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, roomID)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customContractTblModel) GetContractByTime(ctx context.Context, time int64) ([]*ContractTbl, error) {
	var startTime = time - 2*86400000 // lay ra 2 ngay truoc
	query := fmt.Sprintf("select %s from %s where `next_bill` between ? and ?", contractTblRows, m.table)
	var resp []*ContractTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, startTime, time)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
