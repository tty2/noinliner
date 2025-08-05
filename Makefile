##@ Generate

update: ## Update version from golangci-lint
	@./update-version.sh

build: update ## Build binary
	golangci-lint custom -v

replace: build ## Replace golangci-lint with current build
	mv custom-gcl $$(go env GOPATH)/bin/golangci-lint

##@ Other
#------------------------------------------------------------------------------
help:  ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
#------------- <https://suva.sh/posts/well-documented-makefiles> --------------

.DEFAULT_GOAL := help
.PHONY: help update build replace
