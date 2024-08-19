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
	HOUSE_STATUS_DRAFT    = 1
	HOUSE_STATUS_ACTIVE   = 2
	HOUSE_STATUS_INACTIVE = 4
)

const (
	ROOM_STATUS_DRAFT    = 1
	ROOM_STATUS_ACTIVE   = 2
	ROOM_STATUS_INACTIVE = 4
)

// Contract Service
const (
	CONTRACT_STATUS_PENDING  = 1
	CONTRACT_STATUS_ACTIVE   = 2
	CONTRACT_STATUS_INACTIVE = 4
)

const (
	CONTRACT_DETAIL_TYPE_FIXED      = 1
	CONTRACT_DETAIL_TYPE_USAGE      = 2
	CONTRACT_DETAIL_TYPE_FIXED_USER = 4
)

// Payment Service
const (
	BILL_STATUS_UNPAID = 1
	BILL_STATUS_PAID   = 2
)
