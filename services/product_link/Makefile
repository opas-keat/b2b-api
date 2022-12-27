PROJECT_NAME := $(shell basename "$(PWD)")


## dev: run to start dev server
.PHONY: dev
dev:
	$(shell grep -v '^#' .env.dev) air main.go

## local: run to start local server
.PHONY: local
local:
	$(shell grep -v '^#' .env) air main.go

## sec-scan: run for basically security test
.PHONY: sec-scan
sec-scan:
	gosec -exclude=G404 ./...



.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
