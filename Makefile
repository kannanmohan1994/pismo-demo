# VARIABLES
# =======================================================================================
# GO
include .env
GOCMD=go
GOLINT_IMAGE := golangci/golangci-lint:v1.54.2
GOTEST=$(GOCMD) test ./...

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

run: # Builds the project
	${GOCMD} run main.go

lint:
	docker run -t --rm -v $(PWD):/app -w /app --network host $(GOLINT_IMAGE) golangci-lint run

# MIGRATIONS
migrate-up: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" up
migrate-down: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" down --all

# MOCKS
generate-mocks:
	go get github.com/golang/mock
	mockgen -destination=internal/usecase/company/repo_mock.go -package=company xm/internal/repo/company CompanyRepository
	mockgen -destination=internal/usecase/user/repo_mock.go -package=user xm/internal/repo/user UserRepository


# SWAGGER
swagger:
	docker run --rm -v $(PWD):/go/src/xm -w /go/src/xm quay.io/goswagger/swagger generate spec -o /go/src/xm/swagger-ui/swagger.json


