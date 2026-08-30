SHELL := /bin/bash

.PHONY: build dev format install install_modules install_tools postinstall test

all: install

###
# dependencies
###

install:
	$(MAKE) install_tools
	$(MAKE) install_modules
	$(MAKE) postinstall

install_modules:
	@echo ">>> installing modules"
	go mod tidy

install_tools:
	@echo ">>> installing tools"
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/githubnemo/CompileDaemon@latest

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
