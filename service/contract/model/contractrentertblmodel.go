package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ContractRenterTblModel = (*customContractRenterTblModel)(nil)

type (
	// ContractRenterTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractRenterTblModel.
	ContractRenterTblModel interface {
		contractRenterTblModel
		withSession(session sqlx.Session) ContractRenterTblModel
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
