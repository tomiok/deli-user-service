SHELL:=/bin/bash -O extglob
BINARY=user-service
VERSION=0.0.1

LDFLAGS=-ldflags "-X main.Version=${VERSION}"


build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/main/*.go

web:
	@clear
	@go run cmd/web/main/!(*_test).go -E dev

test:
	@clear
	@go test -v ./...

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...