GO ?= go

build:
	@echo "====> Build gost"
	@goreleaser --skip-publish --rm-dist --snapshot
.PHONY: build

install.dev:
	@$(GO) get -u github.com/golang/dep/cmd/dep
	@$(GO) get -u github.com/goreleaser/goreleaser
	@dep ensure -v
.PHONY: install.dev

release:
	@echo "====> Build and release"
	@goreleaser
.PHONY: release
