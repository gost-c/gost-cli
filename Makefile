COMMIT = $$(git describe --always)

generate:
	@go generate ./...

build: generate
	@echo "====> Build gost-cli in ./bin "
	go build -ldflags "-X main.GitCommit=\"$(COMMIT)\" -X command.BaseUrl=\"http://localhost:8000/\" -X command.WebUrl=\"http://localhost:3000/\"" -o bin/gost