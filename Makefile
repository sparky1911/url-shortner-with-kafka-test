MAIN_PACKAGE_PATH := ./cmd/server/main.go
BINARY_NAME := url-shortener
DB_DRIVER := mysql
# Local DSN for Goose (running outside Docker)
DB_STRING := "root:pass@tcp(localhost:3306)/url_shortener?parseTime=true"
GOOSE := $(shell which goose)
MIGRATIONS_DIR := ./migrations

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the application
.PHONY: build
build: tidy
	go build -o=./bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the application locally
.PHONY: run
run: build
	./bin/${BINARY_NAME}

## test: run all tests
.PHONY: test
test:
	go test -v -race ./...

## db/up: Runs all pending migrations
.PHONY: db/up
db/up:
	@if [ -z "$(GOOSE)" ]; then echo "Goose not found. Install with: go install github.com/pressly/goose/v3/cmd/goose@latest"; exit 1; fi
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) up

## db/status: Display the status of the migrations
.PHONY: db/status
db/status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) status

## db/create: Create new migration file (usage: make db/create name=add_hash_column)
.PHONY: db/create
db/create:
ifdef name
	$(GOOSE) -dir $(MIGRATIONS_DIR) create $(name) sql
else
	@echo "Usage: make db/create name=<migration_name>"
endif

## docker/up: Start all containers in background
.PHONY: docker/up
docker/up:
	docker-compose up -d

## docker/down: Stop all containers
.PHONY: docker/down
docker/down:
	docker-compose down