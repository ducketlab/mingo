MOD_DIR := $(shell go env GOMODCACHE)
PKG := "github.com/ducketlab/mingo/$(PROJECT_NAME)"

codegen: # Init Service
	@mingen -I=.  -I${MOD_DIR} --go-ext_out=module=${PKG}:. cmd/protoc-gen-go-ext/extension/tag/*.proto
	@protoc -I=.  -I${MOD_DIR} --go-ext_out=module=${PKG}:. pb/*/*.proto
	@go generate ./...