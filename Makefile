LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=notion-recurring-tasks

# Migration commands
MIGRATIONS_DIR := "db/migrations"

include bin-deps.mk


.PHONY: build
build:
	 CGO_ENABLED=0 go build -v -o $(LOCAL_BIN)/$(PROJECT_NAME) ./cmd

.PHONY: run
run:
	go run cmd/main.go

.PHONY: lint
lint: $(GOLANGCI_BIN)
	$(GOENV) $(GOLANGCI_BIN) run --fix ./...

.PHONY: test
test:
	$(GOENV) go test -race ./...

.PHONY: migrate-up
migrate-up: $(GOOSE_BIN)
	echo "$(DSN)"
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DSN)" up

.PHONY: migrate-down
migrate-down: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DSN)" down

.PHONY: migrate-reset
migrate-reset: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DSN)" reset

.PHONY: migrate-generate
migrate-generate: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) create $(name) sql

.PHONY: migrate-status
migrate-status: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "$(DSN)" status
