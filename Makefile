SHELL := /bin/bash

.PHONY: build dev format install test

all: install

###
# dependencies
###

install:
	@echo ">>> installing dependencies"
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/githubnemo/CompileDaemon@latest
	go mod tidy
	${MAKE} postinstall

postinstall:
	@echo ">>> setting up git hooks"
	git config core.hooksPath "${PWD}/.hooks"

###
# app
###

build:
	${PWD}/scripts/build.sh

dev:
	${PWD}/scripts/dev.sh

###
# formatting
###

format:
	@echo ">>> formatting go files"
	gofmt -w .

###
# testing
###

test:
	@echo ">>> running tests"
	go test -v ./...
