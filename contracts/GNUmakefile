.DEFAULT_GOAL := build

GOPATH ?= $(HOME)/go
COMMIT_SHA ?= $(shell git rev-parse HEAD)
VERSION = $(shell cat VERSION)
GOBIN ?= $(GOPATH)/bin
GO_LDFLAGS := $(shell tools/bin/ldflags)
GOFLAGS = -ldflags "$(GO_LDFLAGS)"

.PHONY: ccip-abi-generate
ccip-abi-generate: ## Generate Public ABI file for a contract
	./scripts/ccip_public_abi_generate $(c)

.PHONY: ccip-abi-fix
ccip-abi-fix: ## Fix contract ABIs that don't match their expected public ABIs
	./scripts/ccip_public_abi_fix

.PHONY: ccip-abi-check
ccip-abi-check: ## Check contract ABIs against their expected public ABIs
	./scripts/ccip_public_abi_check

.PHONY: abigen
abigen: ## Build & install abigen.
	../tools/bin/build_abigen

.PHONY: wrappers
wrappers: abigen ## Recompiles solidity contracts and their go wrappers.
	./scripts/native_solc_compile_all
	go generate ./core/gethwrappers

.PHONY: snapshot
snapshot:
	forge snapshot --match-test _gas

help:
	@echo ""
	@echo "         .__           .__       .__  .__        __"
	@echo "    ____ |  |__ _____  |__| ____ |  | |__| ____ |  | __"
	@echo "  _/ ___\|  |  \\\\\\__  \ |  |/    \|  | |  |/    \|  |/ /"
	@echo "  \  \___|   Y  \/ __ \|  |   |  \  |_|  |   |  \    <"
	@echo "   \___  >___|  (____  /__|___|  /____/__|___|  /__|_ \\"
	@echo "       \/     \/     \/        \/             \/     \/"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
