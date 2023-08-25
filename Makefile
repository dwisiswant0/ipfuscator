APP_NAME := ipfuscator
PKG_DIR := ./pkg/$(APP_NAME)/

vet:
	@go vet ./...

test:
	@go test -race -v $(PKG_DIR)

bench:
	@go test $(PKG_DIR) -bench "^BenchmarkTo"

cover: FILE := /tmp/$(APP_NAME).out
cover:
	@go test -race -coverprofile=$(FILE) -covermode=atomic $(PKG_DIR)
	@go tool cover -func=$(FILE)

build:
	@mkdir -p bin/
	@go build -ldflags "-s -w" -o ./bin/${APP_NAME} .

clean:
	@rm -rf bin/

ci: vet test build clean