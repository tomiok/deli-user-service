SHELL:=/bin/bash -O extglob
BINARY=exp-service
VERSION=0.0.1

LDFLAGS=-ldflags "-X main.Version=${VERSION}"


build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/*.go

web:
	@clear
	@go run cmd/web/!(*_test).go -E dev