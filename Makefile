ROOT_DIR=$(shell pwd)
OUT_DIR=$(ROOT_DIR)/out
LOGS_DIR=$(ROOT_DIR)/logs

MODULE_NAME=roomrover
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/schema
MIGRATION_DIR=$(ROOT_DIR)/migration
DATASOURCE=postgresql://thopb:hkIenQPTp61nQYeMAUVhTDlMo6dcOhSa@dpg-cq962piju9rs73b0k27g-a.singapore-postgres.render.com/humgroom

# account service
ACCOUNT_DIR=$(SERVICE_DIR)/account
ACCOUNT_API_DIR=$(ACCOUNT_DIR)/api
ACCOUNT_MODEL_DIR=$(ACCOUNT_DIR)/model

# list table name
USER_TBL=users_tbl
PROFILE_TBL=profiles_tbl

dep-init:
	go mod init $(MODULE_NAME)
	go mod tidy
dep:
	go mod tidy

# migrate 
migrate:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq $(DB_VERSION)
migrateup:
	migrate -path ${MIGRATION_DIR} -database "${DATASOURCE}" -verbose up
migratedown:
	migrate -path ${MIGRATION_DIR} -database "${DATASOURCE}" -verbose down

#gen api code
gen-account-service:
	goctl api go -api $(API_DIR)/account.api -dir $(ACCOUNT_API_DIR)

# gen db model
gen-account-model: 
	goctl model pg datasource -url="${DATASOURCE}" -table="${USER_TBL}" -dir="$(ACCOUNT_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${PROFILE_TBL}" -dir="$(ACCOUNT_MODEL_DIR)"

runs:
	go run main.go -f etc/server.yaml

