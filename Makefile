ROOT_DIR=$(shell pwd)
OUT_DIR=$(ROOT_DIR)/out
LOGS_DIR=$(ROOT_DIR)/logs

MODULE_NAME=roomrover
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/schema

# account service
ACCOUNT_DIR=$(SERVICE_DIR)/account
ACCOUNT_API_DIR=$(ACCOUNT_DIR)/api
ACCOUNT_MODEL_DIR=$(ACCOUNT_DIR)/model

dep-init:
	go mod init $(MODULE_NAME)
	go mod tidy

dep:
	go mod tidy

#gen api code
gen-account-service:
	goctl api go -api $(API_DIR)/account.api -dir $(ACCOUNT_API_DIR)

# gen db model
gen-account-model: 
	goctl model mysql ddl -src="$(SCHEMA_DIR)/account.sql" -dir="$(ACCOUNT_MODEL_DIR)"

runs:
	go run main.go -f etc/server.yaml

dev:
	@docker-compose down
	@docker-compose up --build

gen-key:
	openssl genrsa -out etc/key.pem 4096
	openssl rsa -in etc/key.pem -pubout > etc/key.pub