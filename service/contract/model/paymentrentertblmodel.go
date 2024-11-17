package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentRenterTblModel = (*customPaymentRenterTblModel)(nil)

type RenterContactModel struct {
	Id          int64  `db:"id"`
	UserID      int64  `db:"user_id"`
	Name        string `db:"name"`
	RoomName    string `db:"room_name"`
	Phone       string `db:"phone"`
	CccdNumber  string `db:"cccd_number"`
	CccdDate    int64  `db:"cccd_date"`
	CccdAddress string `db:"cccd_address"`
	Status      int64  `db:"status"`
}

type (
	// PaymentRenterTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentRenterTblModel.
	PaymentRenterTblModel interface {
		paymentRenterTblModel
		GetRenterByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentRenterTbl, error)
		DeleteByPaymentID(ctx context.Context, paymentID int64) error
		CountRentersByPaymentID(ctx context.Context, paymentID int64) (int64, error)
		CountRenterContacts(ctx context.Context, userID int64, search string) (int, error)
		FilterRenterContacts(ctx context.Context, userID int64, search string, limit, offset int64) ([]*RenterContactModel, error)
	}

	customPaymentRenterTblModel struct {
		*defaultPaymentRenterTblModel
	}
)

// NewPaymentRenterTblModel returns a model for the database table.
func NewPaymentRenterTblModel(conn sqlx.SqlConn) PaymentRenterTblModel {
	return &customPaymentRenterTblModel{
		defaultPaymentRenterTblModel: newPaymentRenterTblModel(conn),
	}
}

func (m *customPaymentRenterTblModel) GetRenterByPaymentID(ctx context.Context, paymentID int64) ([]*PaymentRenterTbl, error) {
	query := fmt.Sprintf("select %s from %s where `payment_id` = ? ", paymentRenterTblRows, m.table)
	var resp []*PaymentRenterTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, paymentID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPaymentRenterTblModel) DeleteByPaymentID(ctx context.Context, paymentID int64) error {
	query := fmt.Sprintf("delete from %s where `payment_id` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, paymentID)
	return err
}

func (m *customPaymentRenterTblModel) CountRentersByPaymentID(ctx context.Context, paymentID int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `payment_id` = ? ", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, paymentID)
	return count, err
}

func (m *customPaymentRenterTblModel) CountRenterContacts(ctx context.Context, userID int64, search string) (int, error) {
	var count int
	query := `SELECT count(*)
	FROM payment_renter_tbl
	JOIN payment_tbl ON payment_renter_tbl.payment_id = payment_tbl.id
	JOIN contract_tbl ON contract_tbl.id = payment_tbl.contact_id
	JOIN room_tbl ON contract_tbl.room_id = room_tbl.id
	JOIN house_tbl ON room_tbl.house_id = house_tbl.id
	JOIN user_tbl ON payment_renter_tbl.user_id = user_tbl.id
	WHERE user_tbl.full_name like ? OR user_tbl.phone LIKE ?
	AND house_tbl.user_id = ?
	ORDER BY contract_tbl.id DESC `
	var vals []interface{}
	vals = append(vals, "%"+search+"%", "%"+search+"%", userID)
	err := m.conn.QueryRowCtx(ctx, &count, query, vals...)
	return count, err
}
func (m *customPaymentRenterTblModel) FilterRenterContacts(ctx context.Context, userID int64, search string, limit, offset int64) ([]*RenterContactModel, error) {
	query := `SELECT
    payment_renter_tbl.id AS id,
	payment_renter_tbl.user_id AS user_id,
    user_tbl.full_name AS name,
    CONCAT(house_tbl.name, ' - ', room_tbl.name) AS room_name,
    user_tbl.phone AS phone,
    user_tbl.CCCD_number AS cccd_number,
    user_tbl.CCCD_date AS cccd_date,
    user_tbl.CCCD_address AS cccd_address,
    payment_renter_tbl.status AS status
	FROM payment_renter_tbl
	JOIN payment_tbl ON payment_renter_tbl.payment_id = payment_tbl.id
	JOIN contract_tbl ON contract_tbl.id = payment_tbl.contact_id
	JOIN room_tbl ON contract_tbl.room_id = room_tbl.id
	JOIN house_tbl ON room_tbl.house_id = house_tbl.id
	JOIN user_tbl ON payment_renter_tbl.user_id = user_tbl.id
	WHERE user_tbl.full_name like ? OR user_tbl.phone LIKE ?
	AND house_tbl.user_id = ?
	ORDER BY contract_tbl.id DESC `
	// CONCAT(house_tbl.name, ' - ', room_tbl.name) LIKE ?
	var vals []interface{}
	vals = append(vals, "%"+search+"%", "%"+search+"%", userID)

	if limit > 0 {
		query += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}

	var resp []*RenterContactModel
	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
