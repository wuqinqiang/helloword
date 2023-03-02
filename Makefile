.PHONY: checklint
checklint:
ifeq (, $(shell which golangci-lint))
	@echo 'error: golangci-lint is not installed, please exec `brew install golangci-lint`'
	@exit 1
endif

.PHONY: lint
lint: checklint
	golangci-lint run --skip-dirs-use-default


win:
	GOOS=windows GOARCH=amd64 go build -o helloword-windows-amd64.exe main.go
macos-adm64:
	GOOS=darwin GOARCH=amd64 go build -o helloword-macos-amd64 main.go
linux:
	GOOS=linux GOARCH=amd64 go build -o helloword-linux-amd64 main.go
macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o helloword-macos-arm64 main.go

build: win macos-adm64 linux macos-arm64


all:
	$(shell sh build.sh)