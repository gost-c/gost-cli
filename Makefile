COMMIT = $$(git describe --always)

generate:
	@go generate ./...

build: generate
	@echo "====> Build gost-cli in ./bin "
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\"" -o bin/gost