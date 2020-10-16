.PHONY: go/build
go/build: ## Compile packages and dependencies
	go build -o $(GOPATH)/bin/clencli

.PHONY: go/install
go/install: ## Compile and install packages and dependencies
	go install

.PHONY: go/run
go/run: ## Compile and run Go program
ifdef command
	go run main.go $(command)
else
	go run main.go
endif

.PHONY: go/fmt
go/fmt: ## gofmt (reformat) package sources
	go fmt main.go
	go fmt cmd/*.go

.PHONY: go/generate
go/generate: ## Generate Go files by processing source
	go generate ./...

.PHONY: go/mod/tidy
go/mod/tidy: ## Add missing and remove unused modules
	go mod tidy

.PHONY: go/get
go/get: ## Add dependencies to current module and install them
	go get ./...

.PHONY: go/clean
go/clean: ## Remove object files and cached files
	go clean -cache -modcache -i -r

.PHONY: go/list
go/list: ## List packages or modules
	go list -m -versions $(package)

.PHONY: go/version
go/version: ## Print Go version
	go version

.PHONY: go/env
go/env: ## Print Go environment information
	go env

.PHONY: go/test
go/test:
	@echo "=== Testing commands"
	@rm -rf test
	@cd cmd && go test -v