.PHONY: all get test build

GO ?= go

all: build test

get:
	${GO} get

build:
	${GO} build -o ${BIN_NAME}

test: get
	${GO} test -v
