GO ?= go

build:
	@echo "====> Build gost"
	@$(GO) test ./...
	@$(GO) build -o bin/gost main.go
.PHONY: build

install.dev:
	@$(GO) get github.com/twitchtv/retool
	@retool add github.com/golang/dep/cmd/dep origin/master
	@retool do dep ensure -v
.PHONY: install.dev

release:
	@echo "====> Build and release"
	@curl -sL https://git.io/goreleaser | bash
.PHONY: release
