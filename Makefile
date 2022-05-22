# Common vars
IMPORT_PATH ?= github.com/ew0s/trade-bot
BUILD_DIR ?= bin
PKG_DIR = .pkg
GOROOT ?= /usr/local/go

ENV := $(if $(ENV),$(ENV),local)

.PHONY: test
test:
	go test -count=1 ./... -covermode=atomic -v -race

.PHONY: generate
generate:
	rm -rf $$(find $(BINARIES_DIR) -maxdepth 2 -name env -type d)
	go install $(IMPORT_PATH)/internal/configer
	go generate ./...

.PHONY: install-swag
install-swag:
	go get -u github.com/swaggo/swag/cmd/swag
	go install -mod=readonly github.com/swaggo/swag/cmd/swag

.PHONY: swagger-api
swagger-api:
	@swag init --dir=./cmd/api --output=./cmd/api/swagger --parseVendor --parseDependency --parseInternal
	@swag fmt --dir=./cmd/api

.PHONY: migrate-status
migrate-status:
	ENV=$(ENV) ./migrations/trade-bot/migrator.sh status

.PHONY: migrate-up
migrate-up:
	ENV=$(ENV) ./migrations/trade-bot/migrator.sh up

.PHONY: migrate-down
migrate-down:
	ENV=$(ENV) ./migrations/trade-bot/migrator.sh down

.PHONY: migrate-new
migrate-new:
	./scripts/migrate-new.sh

.PHONY: migrate-add-timestamp
migrate-add-timestamp:
	./scripts/migrate-add-timestamp.sh

# See https://github.com/rubenv/sql-migrate
.PHONY: install-migrator
	go install -mod=readonly github.com/rubenv/sql-migrate/...
