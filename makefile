.PHONY: default dev run clear install test docs build db_generate migrations_up migrations_down migrations_create

include .env

default: dev

dev:
	@air
run:
	@go run ./cmd/server
clear:
	@rm ./tmp/main
install:
	@go mod download && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && go install github.com/swaggo/swag/cmd/swag@latest && go install github.com/pressly/goose/v3/cmd/goose@latest && go install github.com/air-verse/air@latest
test:
	@go test ./...
docs:
	@swag init -g ./cmd/server/main.go -o ./docs
build:
	@GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./tmp/server ./cmd/server
db_generate:
	@sqlc generate
migrations_up:
	@goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up
migrations_down:
	@goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" down
migrations_create:
	@goose create $(NAME) sql
