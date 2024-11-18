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
	ROOM_STATUS_INACTIVE = 1 // Cho xac nhan
	ROOM_STATUS_ACTIVE   = 2 // con cho thue
	ROOM_STATUS_RENTED   = 4 // da cho thue
	ROOM_STATUS_STOP     = 8 // tam dung
)

const (
	CONTACT_STATUS_TYPE_WATTING = 1 // Cho xac nhan
	CONTACT_STATUS_TYPE_CONFIRM = 2 // Dong y
	CONTACT_STATUS_TYPE_CANCEL  = 4 // Huy
)

// Contract Service
const (
	CONTRACT_STATUS_DRAF            = 1  // nhap
	CONTRACT_STATUS_WAIT_DEPOSIT    = 2  // cho dat coc
	CONTRACT_STATUS_ACTIVE          = 4  // hoat dong
	CONTRACT_STATUS_OUT_DATE        = 8  // het han
	CONTRACT_STATUS_INACTIVE        = 16 // huy
	CONTRACT_STATUS_NEARLY_OUT_DATE = 32 // sap het han
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
	PAYMENT_DETAIL_TYPE_FIXED      = 1 // co dinh
	PAYMENT_DETAIL_TYPE_USAGE      = 2 // sl su dung
	PAYMENT_DETAIL_TYPE_FIXED_USER = 4 // so luong nguoi
	PAYMENT_DETAIL_TYPE_ROOM       = 8 // gia phong
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

	NOTI_TYPE_CREATE_CONTACT           = 1    // tao lien he
	NOTI_TYPE_CONFIRM_CONTACT          = 2    // xac nhan lien he
	NOTI_TYPE_REJECT_CONTACT           = 4    // tu choi lien he
	NOTI_TYPE_CREATE_CONTRACT          = 8    // tao hop dong
	NOTI_TYPE_UPDATE_CONTRACT          = 16   // cap nhat hop dong
	NOTI_TYPE_CANCEL_CONTRACT          = 32   // huy hop dong
	NOTI_TYPE_CONFIRM_CONTRACT         = 64   // xac nhan hop dong
	NOTI_TYPE_OUT_DATE_CONTRACT        = 128  // hop dong het han
	NOTI_TYPE_NEARLY_OUT_DATE_CONTRACT = 256  // hop dong sap het han
	NOTI_TYPE_CREATE_BILL              = 512  // tao hoa don
	NOTI_TYPE_OUT_DATE_BILL            = 1024 // hoa don het han
	NOTI_TYPE_PAY_BILL                 = 2048 // thanh toan hoa don
)

const MIN_ID = 10000000000000

const (
	BILL_PAY_TYPE_MONEY    = 1
	BILL_PAY_TYPE_TRANSFER = 2
	BILL_PAY_TYPE_ZALO     = 4
)

const (
	BILL_PAY_STATUS_PROCESS = 1 // cho xu ly
	BILL_PAY_STATUS_DONE    = 2 // thanh cong
)
