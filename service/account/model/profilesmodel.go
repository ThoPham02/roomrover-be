package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProfilesModel = (*customProfilesModel)(nil)

type (
	// ProfilesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProfilesModel.
	ProfilesModel interface {
		profilesModel
	}

	customProfilesModel struct {
		*defaultProfilesModel
	}
)

// NewProfilesModel returns a model for the database table.
func NewProfilesModel(conn sqlx.SqlConn) ProfilesModel {
	return &customProfilesModel{
		defaultProfilesModel: newProfilesModel(conn),
	}
}
