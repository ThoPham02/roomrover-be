package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ContractTblModel = (*customContractTblModel)(nil)

type (
	// ContractTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractTblModel.
	ContractTblModel interface {
		contractTblModel
		withSession(session sqlx.Session) ContractTblModel
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
