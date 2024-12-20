lint: ## Lint code
	golangci-lint run --fix

test: ## Run tests
	go test ./...
