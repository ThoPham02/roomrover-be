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
