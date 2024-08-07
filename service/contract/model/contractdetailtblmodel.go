package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ContractDetailTblModel = (*customContractDetailTblModel)(nil)

type (
	// ContractDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractDetailTblModel.
	ContractDetailTblModel interface {
		contractDetailTblModel
		withSession(session sqlx.Session) ContractDetailTblModel
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
