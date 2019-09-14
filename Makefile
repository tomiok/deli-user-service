SHELL:=/bin/bash -O extglob
BINARY=user-service
VERSION=0.0.1

LDFLAGS=-ldflags "-X main.Version=${VERSION}"


build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/main/*.go

web:
	@clear
	@go run cmd/web/main/!(*_test).go -E dev