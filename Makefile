GO ?= go

build:
	@echo "====> Build gost"
	@$(GO) build -o bin/gost main.go
.PHONY: build

install.dev:
	@$(GO) get -u github.com/golang/dep/cmd/dep
	@$(GO) get -u github.com/goreleaser/goreleaser
	@dep ensure -vendor-only
.PHONY: install.dev

release:
	@echo "====> Build and release"
	@goreleaser
.PHONY: release
