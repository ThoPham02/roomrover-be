package common

// Constants
const (
	USER_ACTIVE   = 1
	USER_INACTIVE = 2
)

const NO_USE = -1

// User Service
const (
	USER_ROLE_SYS_ADMIN = 0
	USER_ROLE_ADMIN     = 1
	USER_ROLE_RENTER    = 2 // thue
	USER_ROLE_LESSOR    = 4 // cho thue
)

// Inventory Service
const (
	HOUSE_STATUS_DRAFT    = 1 // nhap
	HOUSE_STATUS_ACTIVE   = 2 // con phong
	HOUSE_STATUS_INACTIVE = 4 // khong hoat dong
	HOUSE_STATUS_SOLD_OUT = 8 // het phong
)

const (
	ROOM_STATUS_INACTIVE        = 1  // Cho xac nhan
	ROOM_STATUS_ACTIVE          = 2  // con cho thue
	ROOM_STATUS_RENTED          = 4  // da cho thue
	ROOM_STATUS_NEARLY_OUT_DATE = 8  // sap het han
	ROOM_STATUS_TRANSFER        = 16 // chuyen nhuong
)

const (
	CONTACT_STATUS_TYPE_WATTING = 1 // Cho xac nhan
	CONTACT_STATUS_TYPE_CONFIRM = 2 // Dong y
)

// Contract Service
const (
	CONTRACT_STATUS_DRAF            = 1  // nhap
	CONTRACT_STATUS_ACTIVE          = 2  // hoat dong
	CONTRACT_STATUS_INACTIVE        = 4  // huy
	CONTRACT_STATUS_NEARLY_OUT_DATE = 8  // sap het han
	CONTRACT_STATUS_TRANSFER        = 16 // chuyen nhuong
	CONTRACT_STATUS_OUT_DATE        = 32 // het han
)

const (
	CONTRACT_DETAIL_TYPE_FIXED      = 1
	CONTRACT_DETAIL_TYPE_USAGE      = 2
	CONTRACT_DETAIL_TYPE_FIXED_USER = 4
)

// Payment Service
const (
	PAYMENT_DETAIL_STATUS_DRAF    = 1
	PAYMENT_DETAIL_STATUS_PROCESS = 2
	PAYMENT_DETAIL_STATUS_DONE    = 4
)

const (
	PAYMENT_DETAIL_TYPE_FIXED      = 1
	PAYMENT_DETAIL_TYPE_USAGE      = 2
	PAYMENT_DETAIL_TYPE_FIXED_USER = 4
)

const (
	BILL_STATUS_DRAF     = 1
	BILL_STATUS_UNPAID   = 2
	BILL_STATUS_PAID     = 4
	BILL_STATUS_OUT_DATE = 8
)

// NOTIFICATION Service
const (
	NOTIFICATION_REF_TYPE_BILL = 1
)

const (
	NOTI_STATUS_PENDING = 1
	NOTI_STATUS_DONE    = 2

	NOTI_TYPE_UNREAD = 1
	NOTI_TYPE_READ   = 2
)

const MIN_ID = 1000000000

const (
	BILL_PAY_TYPE_MONEY    = 1
	BILL_PAY_TYPE_TRANSFER = 2
	BILL_PAY_TYPE_ZALO     = 4
)

const (
	BILL_PAY_STATUS_PROCESS = 1
	BILL_PAY_STATUS_DONE    = 2
)
