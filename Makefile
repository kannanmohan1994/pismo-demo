# VARIABLES
# =======================================================================================
# GO
include .env
GOCMD=go
GOLINT_IMAGE := golangci/golangci-lint:v1.61.0
GOTEST=$(GOCMD) test ./...
SWAG=$(GOCMD) run github.com/swaggo/swag/cmd/swag@v1.16.6
SWAGGER_DOCS_DIR=docs
SWAGGER_DOCS_SPEC=$(SWAGGER_DOCS_DIR)/swagger.json
SWAGGER_UI_SPEC=$(SWAGGER_DOCS_DIR)/swagger.json

# GOTOOLS

vet: # Vet examines Go source code and reports suspicious constructs
	${GOCMD} vet

fmt: # Gofmt is a tool that automatically formats Go source code
	gofmt

test: # GO Test
	$(GOTEST)

cover: # Go Test Coverage
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out

tidy: # Update Modules and Dependency Consistency
	${GOCMD} mod tidy

build: # Builds the project
	${GOCMD} build main.go

run: swagger # Builds the project and keeps swagger spec up to date
	${GOCMD} run main.go

# MIGRATIONS
migrate-up: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" up
migrate-down: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" down --all
migrate-create:
	docker run --rm \
		-u $(shell id -u):$(shell id -g) \
		-v $(PWD)/db/migrations:/migrations \
		migrate/migrate \
		create -ext sql -dir /migrations -seq $(name)
		
# MOCKS
generate-mocks:
	go get github.com/golang/mock
	mockgen -destination=internal/mocks/account_repo_mock.go -package=mocks pismo/internal/repo/account AccountRepository
	mockgen -destination=internal/mocks/user_repo_mock.go -package=mocks pismo/internal/repo/user UserRepository
	mockgen -destination=internal/mocks/transaction_repo_mock.go -package=mocks pismo/internal/repo/transaction TransactionRepository
	mockgen -destination=internal/mocks/operationtype_repo_mock.go -package=mocks pismo/internal/repo/operationtype OperationTypeRepository

# SWAGGER
swagger:
	mkdir -p ${SWAGGER_DOCS_DIR}
	${SWAG} init -g main.go -o ${SWAGGER_DOCS_DIR}
