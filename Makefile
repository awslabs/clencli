# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

include lib/make/*/Makefile

.PHONY: clencli/test
clencli/test: go/test

.PHONY: clencli/build
clencli/build: clencli/version go/mod/tidy go/version go/get go/fmt go/generate go/build ## Builds the app

.PHONY: clencli/install
clencli/install: go/get go/fmt go/generate go/install ## Builds the app and install all dependencies

.PHONY: clencli/run
clencli/run: go/fmt ## Run a command
ifdef command
	make go/run command='$(command)'
else
	make go/run
endif

.PHONY: clencli/compile
clencli/compile: ## Compile to multiple architectures
	@mkdir -p dist
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o dist/clencli-darwin-amd64 main.go
	GOOS=solaris GOARCH=amd64 go build -o dist/clencli-solaris-amd64 main.go

	GOOS=freebsd GOARCH=386 go build -o dist/clencli-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o dist/clencli-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm go build -o dist/clencli-freebsd-arm main.go

	GOOS=openbsd GOARCH=386 go build -o dist/clencli-openbsd-386 main.go
	GOOS=openbsd GOARCH=amd64 go build -o dist/clencli-openbsd-amd64 main.go
	GOOS=openbsd GOARCH=arm go build -o dist/clencli-openbsd-arm main.go

	GOOS=linux GOARCH=386 go build -o dist/clencli-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/clencli-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o dist/clencli-linux-arm main.go

	GOOS=windows GOARCH=386 go build -o dist/clencli-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/clencli-windows-amd64 main.go

.PHONY: clencli/tag
clencli/tag: ## Tag a version
ifdef version
	git tag -a v$(version) -m 'Release version v$(version)'
else
	@echo "version not specified"
endif

.PHONY: clencli/clean
clencli/clean: ## Removes unnecessary files and directories
	rm -rf downloads/
	rm -rf generated-*/
	rm -rf dist/
	rm -rf build/

.PHONY: clencli/update-readme
clencli/update-readme: ## Renders template readme.tmpl with additional documents
	@echo "Updating README.tmpl to the latest version"
	@cp box/resources/init/clencli/readme.tmpl clencli/readme.tmpl
	@echo "Generate COMMANDS.md"
	@echo "## Commands" > COMMANDS.md
	@echo '```' >> COMMANDS.md
	@clencli --help >> COMMANDS.md
	@echo '```' >> COMMANDS.md
	@echo "COMMANDS.md generated successfully"
	@clencli render template --name readme

.PHONY: clencli/test
clencli/test: go/test

.DEFAULT_GOAL := clencli/help

.PHONY: clencli/help
clencli/help: ## This HELP message
	@fgrep -h ": ##" $(MAKEFILE_LIST) | sed -e 's/\(\:.*\#\#\)/\:\ /' | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

split = $(word $2,$(subst $3, ,$1))
word-slash = $(word $2,$(subst /, ,$1))

CURRENT_BRANCH := $(shell git branch --show-current)
CURRENT_COMMIT := $(shell git rev-parse --short HEAD)
LATEST_TAG := $(shell git describe --tags --abbrev=0)

RELEASE_VERSION=v$(call word-slash,$(CURRENT_BRANCH),2)
CANDIDATE_VERSION=$(LATEST_TAG)-rc.

.PHONY: clencli/release
clencli/release: go/mod/tidy
	@echo CURRENT BRANCH IS: $(CURRENT_BRANCH)
	@echo CURRENT COMMIT IS: $(CURRENT_COMMIT)
	@echo LATEST TAG IS: $(LATEST_TAG)
ifneq (,$(findstring release,$(CURRENT_BRANCH)))
	@echo RELEASE FINAL VERSION
	git tag $(RELEASE_VERSION)
else ifneq (,$(findstring develop,$(CURRENT_BRANCH)))
	@echo RELEASE CANDIDATE VERSION
	@echo git tag $(CANDIDATE_VERSION)
else ifneq (,$(findstring feature,$(CURRENT_BRANCH)))
	@echo RELEASE DEV SNAPSTHO
else
	@echo Not found
endif
