SHA=$(shell git rev-parse --short --verify HEAD)
TAG=$(shell git describe --tags HEAD 2> /dev/null || echo Undefined)
PKG=$(shell go list ./...)
SRC=$(wildcard **/*.go)
BUILD=$(wildcard bin/*)

LDFLAGS=""

default: clean deps build

clean:
	rm -rf bin

deps:
	dep ensure

build: bin/fortune-api

bin/fortune-api: $(SRC)
	go build -ldflags ${LDFLAGS} -o $@ main.go

lint:
	(! gofmt -e -l . | read) || (gofmt -d . && false)
	golint -set_exit_status $(PKG)

test:
	go test ./... -cover

.PHONY: lint test
