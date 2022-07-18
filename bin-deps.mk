GOLANGCI_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLANGCI_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint

# db related tools

GOOSE_BIN=$(LOCAL_BIN)/goose
$(GOOSE_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

SQLBOILER_BIN=$(LOCAL_BIN)/sqlboiler
$(SQLBOILER_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/volatiletech/sqlboiler/v4

SQLBOILER_PSQL_BIN=$(LOCAL_BIN)/sqlboiler-psql
$(SQLBOILER_PSQL_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql

IFACEMAKER_BIN=$(LOCAL_BIN)/ifacemaker
$(IFACEMAKER_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/vburenin/ifacemaker