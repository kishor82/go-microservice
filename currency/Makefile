.PHONY: protos install-protoc install-go-plugins
# Makefile to install protoc and Go plugins

# Specify the version of protoc you want to install
PROTOC_VERSION := 25.1

# Specify the Go protoc-gen-go and protoc-gen-go-grpc plugin versions
PROTOC_GEN_GO_VERSION := v1.31.0
PROTOC_GEN_GO_GRPC_VERSION := v1.3.0

# Specify the operating system and architecture
OS := $(shell uname -s)
ARCH := $(shell uname -m)

# Specify the installation directory
INSTALL_DIR := /usr/local/bin

GO=go

protos:
	protoc -I protos/ --go_out=protos --go-grpc_out=protos protos/currency.proto

install-protoc:
	@if [ ! -x "$(INSTALL_DIR)/protoc" ]; then \
		wget https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(OS)-$(ARCH).zip; \
		unzip protoc-$(PROTOC_VERSION)-$(OS)-$(ARCH).zip -d protoc-$(PROTOC_VERSION); \
		sudo mv protoc-$(PROTOC_VERSION)/bin/protoc $(INSTALL_DIR)/protoc; \
		rm -rf protoc-$(PROTOC_VERSION) protoc-$(PROTOC_VERSION)-$(OS)-$(ARCH).zip; \
	else \
		echo "protoc is already installed"; \
	fi

install-go-plugins:
	@if [ ! -x "$(shell go env GOPATH)/bin/protoc-gen-go" ]; then \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION); \
	else \
		echo "protoc-gen-go is already installed"; \
	fi
	@if [ ! -x "$(shell go env GOPATH)/bin/protoc-gen-go-grpc" ]; then \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION); \
	else \
		echo "protoc-gen-go-grpc is already installed"; \
	fi

install: install-protoc install-go-plugins
