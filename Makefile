# Common vars
IMPORT_PATH ?= github.com/ew0s/trade-bot
BUILD_DIR ?= bin
PKG_DIR = .pkg
GOROOT ?= /usr/local/go

.PHONY: test
test:
	go test -count=1 ./... -covermode=atomic -v -race

.PHONY: generate
generate:
	rm -rf $$(find $(BINARIES_DIR) -maxdepth 2 -name env -type d)
	go install $(IMPORT_PATH)/internal/configer
	go generate ./...

.PHONY: swagger-api
swagger-api:
	@swag init --dir=./cmd/api --output=./cmd/api/swagger --parseVendor --parseDependency --parseInternal
	@swag fmt --dir=./cmd/api
