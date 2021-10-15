MOD_DIR := $(shell go env GOPATH)
PROJECT_NAME := "mingo"
PKG := "github.com/ducketlab/$(PROJECT_NAME)"
MINGO_MAIN := "cmd/mingo/main.go"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v redis | grep -v broker | grep -v etcd | grep -v examples)

dep: ## Get the dependencies
	@go mod download

install: ## install mingo cli
	@go install ${PKG}/cmd/mingo
	@go install ${PKG}/cmd/protoc-gen-go-ext
	@go install ${PKG}/cmd/protoc-gen-go-http

build: dep ## Build the binary file
	@go build -o build/$(PROJECT_NAME) $(MINGO_MAIN)

vet: ## Run go vet
	@go vet ${PKG_LIST}

codegen: ## Init Service
	@protoc -I=.  -I${MOD_DIR}/src --go-ext_out=module=${PKG}:. cmd/protoc-gen-go-ext/extension/tag/*.proto
	@protoc -I=.  -I${MOD_DIR}/src --go-ext_out=module=${PKG}:. pb/*/*.proto
	@go generate ./...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'