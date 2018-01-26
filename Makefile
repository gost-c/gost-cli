GO ?= go

build:
	@echo "====> Build gost"
	@retool do goreleaser --skip-publish --rm-dist --snapshot
.PHONY: build

install.dev:
	@$(GO) get github.com/twitchtv/retool
	@retool add github.com/golang/dep/cmd/dep origin/master
	@retool add github.com/goreleaser/goreleaser v0.45.1
	@retool do dep ensure -v
.PHONY: install.dev

release:
	@echo "====> Build and release"
	@retool do goreleaser
.PHONY: release
