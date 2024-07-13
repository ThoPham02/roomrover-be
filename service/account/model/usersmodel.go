package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		CheckUserExists(userName string) bool
		Register(userName, password string) error
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) CheckUserExists(userName string) bool {
	query := `SELECT COUNT(*) FROM users WHERE user_name = ?`
	var count int

	err := m.conn.QueryRow(&count, query, userName)
	if err != nil || count > 0 {
		logx.Error(err)
		return false
	}

	return true
}

func (m *customUsersModel) Register(userName, password string) error {
	query := `INSERT INTO users (user_name, password) VALUES (?, ?)`
	_, err := m.conn.Exec(query, userName, password)
	return err
}
