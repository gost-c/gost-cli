
generate:
	@go generate ./...

build: generate
	@echo "====> Build gost-cli in ./pkg "
	./build.sh