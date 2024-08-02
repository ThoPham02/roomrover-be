ROOT_DIR=$(shell pwd)
OUT_DIR=$(ROOT_DIR)/out
LOGS_DIR=$(ROOT_DIR)/logs

MODULE_NAME=roomrover
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/schema
MIGRATION_DIR=$(ROOT_DIR)/migration
DATASOURCE=mysql://thopb:524020@localhost:3306/humgroom

# account service
ACCOUNT_DIR=$(SERVICE_DIR)/account
ACCOUNT_API_DIR=$(ACCOUNT_DIR)/api
ACCOUNT_MODEL_DIR=$(ACCOUNT_DIR)/model

# inventory service
INVENTORY_DIR=$(SERVICE_DIR)/inventory
INVENTORY_API_DIR=$(INVENTORY_DIR)/api
INVENTORY_MODEL_DIR=$(INVENTORY_DIR)/model

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

gen-inventory-service:
	goctl api go -api $(API_DIR)/inventory.api -dir $(INVENTORY_API_DIR)

# gen db model
gen-account-model: 
	goctl model mysql ddl -src="${SCHEMA_DIR}/account.sql" -dir="${ACCOUNT_MODEL_DIR}" --ignore-columns=""

gen-inventory-model:
	goctl model mysql ddl -src="${SCHEMA_DIR}/inventory.sql" -dir="${INVENTORY_MODEL_DIR}" --ignore-columns=""

runs:
	go run main.go -f etc/server.yaml

