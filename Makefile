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

# inventory service
INVENTORY_DIR=$(SERVICE_DIR)/inventory
INVENTORY_API_DIR=$(INVENTORY_DIR)/api
INVENTORY_MODEL_DIR=$(INVENTORY_DIR)/model

# list table name
USER_TBL=users_tbl
PROFILE_TBL=profiles_tbl

HOME_TBL=homes_tbl
ROOM_TBL=rooms_tbl
ROOM_GROUP_TBL=room_groups_tbl
ROOM_ALBUM_TBL=room_albums_tbl

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
	goctl model pg datasource -url="${DATASOURCE}" -table="${USER_TBL}" -dir="$(ACCOUNT_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${PROFILE_TBL}" -dir="$(ACCOUNT_MODEL_DIR)"

gen-inventory-model:
	goctl model pg datasource -url="${DATASOURCE}" -table="${HOME_TBL}" -dir="$(INVENTORY_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${ROOM_TBL}" -dir="$(INVENTORY_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${ROOM_GROUP_TBL}" -dir="$(INVENTORY_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${ROOM_ALBUM_TBL}" -dir="$(INVENTORY_MODEL_DIR)"

runs:
	go run main.go -f etc/server.yaml

