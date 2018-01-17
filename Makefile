SHA=$(shell git rev-parse --short --verify HEAD)
TAG=$(shell git describe --tags HEAD 2> /dev/null || echo Undefined)
SRC=$(wildcard **/*.go)
BUILD=$(wildcard build/*)

LDFLAGS=""

default: clean deps build

clean:
	rm -rf build

deps:
	dep ensure

build: build/fortune-api

build/fortune-api: $(SRC)
	go build -ldflags ${LDFLAGS} -o $@ main.go

test:
	go test ./... -cover

.PHONY: test
