SHELL := /bin/bash

.PHONY: build format install install_modules install_tools postinstall test

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

postinstall:
	@echo ">>> setting up git hooks"
	git config core.hooksPath "${PWD}/.hooks"

###
# build
###

build:
	${PWD}/scripts/build.sh

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
