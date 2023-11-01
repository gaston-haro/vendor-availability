BINARY=service

# Actions to execute locally
build:
	go build -o ${BINARY} cmd/web/*.go

clean:
	go clean -testcache ./...

test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

test_local:
	go test -cover ./...

test_local_coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
	
web:
	@clear
	go build -o ${BINARY} cmd/web/*.go
	./${BINARY} -E dev

# Protobuf generation
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTO_DIR := ./proto
GO_OUT_DIR := ./pb
PROTO_FILES := $(shell find $(PROTO_DIR) -type f -name '*.proto')

generate:
	@rm -rf $(GO_OUT_DIR) && \
	mkdir -p $(GO_OUT_DIR) && \
	cd $(GO_OUT_DIR) && go mod init github.com/deliveryhero  && cd .. && \
	for f in $(PROTO_FILES); do protoc --go_out=paths=import:$(GO_OUT_DIR) --go_opt=module=github.com/deliveryhero --go-grpc_out=paths=import:$(GO_OUT_DIR) --go-grpc_opt=module=github.com/deliveryhero --proto_path=$(PROTO_DIR) $$f; done && \
	cd $(GO_OUT_DIR) && go mod tidy -go=1.21

check_protoc_gen_go:
ifndef PROTOC_GEN_GO
	$(info  PATH: $(PATH))
	$(info  SHELL: $(SHELL))
	$(error "protoc-gen-go not found in PATH. Please install it by running 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'")
endif

check_protoc_gen_go_grpc:
ifndef PROTOC_GEN_GO_GRPC
	$(error "protoc-gen-go-grpc not found in PATH. Please install it by running 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'")
endif

.PHONY: generate check_protoc_gen_go check_protoc_gen_go_grpc


# Optional actions for projects that may need them
#test_integration:
#	go clean -testcache && go test -v -tags=integration ./test/integration