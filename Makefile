.PHONY: all get test build

GO ?= go

all: build test

get:
	${GO} get

build:
	${GO} build

test: get
	${GO} test -v
