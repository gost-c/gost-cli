GO ?= go

build:
	@echo "====> Build gen"
	@goreleaser --skip-publish --rm-dist --snapshot
.PHONY: build

install.dev:
	@$(GO) get github.com/golang/dep/cmd/dep
	@$(GO) get github.com/goreleaser/goreleaser
	@dep ensure
.PHONY: install.dev

release:
	@echo "====> Build and release"
	@goreleaser
.PHONY: release
