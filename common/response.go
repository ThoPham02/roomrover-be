package common

// Common response code
const (
	SUCCESS_CODE    = 0
	SUCCESS_MESSAGE = "Success"

	INVALID_REQUEST_CODE    = 1
	INVALID_REQUEST_MESSAGE = "Invalid request"

	DB_ERROR_CODE    = 2
	DB_ERROR_MESSAGE = "Database error"
)

// Account service response code from 10000-19999
const (
	USER_ALREADY_EXISTS_CODE    = 10000
	USER_ALREADY_EXISTS_MESSAGE = "User already exists"

	USER_NOT_FOUND_CODE    = 10001
	USER_NOT_FOUND_MESSAGE = "User not found"

	INVALID_PASSWORD_CODE    = 10002
	INVALID_PASSWORD_MESSAGE = "Invalid password"
)
