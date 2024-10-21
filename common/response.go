package common

// Common response code
const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "Success"

	INVALID_REQUEST_CODE = 1
	INVALID_REQUEST_MESS = "Invalid request"

	DB_ERR_CODE = 2
	DB_ERR_MESS = "Database error"

	UNKNOWN_ERR_CODE = 3
	UNKNOWN_ERR_MESS = "Unknown error"

	PERMISSION_DENIED_ERR_CODE = 4
	PERMISSION_DENIED_ERR_MESS = "PERMISSION DENIED"
)

// Account service response code from 10000-19999
const (
	USER_ALREADY_EXISTS_CODE = 10000
	USER_ALREADY_EXISTS_MESS = "User already exists"

	USER_NOT_FOUND_CODE = 10001
	USER_NOT_FOUND_MESS = "User not found"

	INVALID_PASSWORD_CODE = 10002
	INVALID_PASSWORD_MESS = "Invalid password"
)

// Inventory service response code from 20000-29999
const (
	HOUSE_NOT_FOUND_CODE = 20000
	HOUSE_NOT_FOUND_MESS = "House not found"

	ROOM_NOT_FOUND_CODE = 20001
	ROOM_NOT_FOUND_MESS = "Room not found"

	ROOM_HAS_CONTRACT_CODE = 20002
	ROOM_HAS_CONTRACT_MESS = "Room has contract"

	HOUSE_HAS_CONTRACT_ERR_CODE = 20003
	HOUSE_HAS_CONTRACT_ERR_MESS = "House has contract"
)

// Contract service response code from 30000-39999
const (
	CONTRACT_NOT_FOUND_CODE = 30000
	CONTRACT_NOT_FOUND_MESS = "Contract not found"

	CONTRACT_ALREADY_EXISTS_CODE = 30001
	CONTRACT_ALREADY_EXISTS_MESS = "Contract already exists"

	CONTRACT_HAS_BEEN_APPROVED_CODE = 30002
	CONTRACT_HAS_BEEN_APPROVED_MESS = "Contract has been approved"

	CONTRACT_HAS_BEEN_REJECTED_CODE = 30003
	CONTRACT_HAS_BEEN_REJECTED_MESS = "Contract has been rejected"
)

const (
	BILL_NOT_FOUND_CODE = 40000
	BILL_NOT_FOUND_MESS = "Bill not found"

	BILL_PAY_NOT_FOUND_CODE = 40001
	BILL_PAY_NOT_FOUND_MESS = "Bill pay not found"
)
