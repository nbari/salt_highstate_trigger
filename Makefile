.PHONY: all get test build

# To compile for freebsd
# env GOOS=freebsd GOARCH=amd64 go build

GO ?= go

all: build test

get:
	${GO} get -u

build:
	${GO} build

test: get
	${GO} test -v
