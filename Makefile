export IMAGE_REPOSITORY?=docker.io/khulnasoft
export DF_IMG_TAG?=2.5.2

TEMP_DIR := ./.tmp

# Command templates #################################
LINT_CMD := $(TEMP_DIR)/golangci-lint run --tests=false
GOIMPORTS_CMD := $(TEMP_DIR)/gosimports -local github.com/khulnasoft-lab

# Tool versions #################################
GOLANGCILINT_VERSION := v1.55.1
GOSIMPORTS_VERSION := v0.3.8
GOLICENSES_VERSION := v5.0.1

# Formatting variables #################################
BOLD := $(shell tput -T linux bold)
PURPLE := $(shell tput -T linux setaf 5)
GREEN := $(shell tput -T linux setaf 2)
CYAN := $(shell tput -T linux setaf 6)
RED := $(shell tput -T linux setaf 1)
RESET := $(shell tput -T linux sgr0)
TITLE := $(BOLD)$(PURPLE)
SUCCESS := $(BOLD)$(GREEN)

all: yarasec

bootstrap:
	$(PWD)/bootstrap.sh

clean:
	-rm ./YaraSec

vendor: go.mod
	go mod tidy -v
	go mod vendor

yarasec: vendor $(PWD)/**/*.go $(PWD)/agent-plugins-grpc/**/*.go
	CGO_LDFLAGS="-ljansson -lcrypto -lmagic" PKG_CONFIG_PATH=/usr/local/yara/lib/pkgconfig:$(PKG_CONFIG_PATH) go build -buildmode=pie -ldflags="-s -w -extldflags=-static -X 'main.version=$(DF_IMG_TAG)'" -buildvcs=false -v .

.PHONY: clean bootstrap

.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 docker build -t $(IMAGE_REPOSITORY)/khulnasoft_malware_scanner_ce:$(DF_IMG_TAG) .
