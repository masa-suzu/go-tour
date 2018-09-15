.PHONY:all test fmt help

all:f t ## Execute unit tests after formatting
t: ## Execute unit tests
	go test ./...

f: ## Format go files
	go fmt ./...
	goimports -w ./
	go vet ./...
	golint ./...

help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
