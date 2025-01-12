NAME=sbc-bakery
BINARY=packer-plugin-${NAME}
GO_VERSION := 1.23.4
HASHICORP_PACKER_PLUGIN_SDK_VERSION?=$(shell go list -m github.com/hashicorp/packer-plugin-sdk | cut -d " " -f2)

build:
	@go build -o ${BINARY}

# Install packer software development command (packer-sdc) utility
install-packer-sdc:
	@go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@${HASHICORP_PACKER_PLUGIN_SDK_VERSION}

plugin-check: install-packer-sdc build
	@packer-sdc plugin-check ${BINARY}

set-go-version:
	go mod edit -go=$(GO_VERSION)
	sed -i "s/FROM golang:.*/FROM golang:$(GO_VERSION)/" .devcontainer/Dockerfile
	echo "$(GO_VERSION)" > .go-version
