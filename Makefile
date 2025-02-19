# Variables
API_URL ?= http://localhost:3000

# Help
.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ API Testing

.PHONY: test-api
test-api: ## Run all API tests
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh all

.PHONY: test-register
test-register: ## Test register endpoint
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh register

.PHONY: test-login
test-login: ## Test login endpoint
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh login

.PHONY: test-users
test-users: ## Test get users endpoint
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh users

##@ Custom URL Testing

.PHONY: test-api-custom
test-api-custom: ## Run all API tests with custom URL (make test-api-custom API_URL=http://example.com)
	@chmod +x scripts/test-api.sh
	@API_URL=$(API_URL) ./scripts/test-api.sh all

.PHONY: test-register-custom
test-register-custom: ## Test register endpoint with custom URL
	@chmod +x scripts/test-api.sh
	@API_URL=$(API_URL) ./scripts/test-api.sh register

.PHONY: test-login-custom
test-login-custom: ## Test login endpoint with custom URL
	@chmod +x scripts/test-api.sh
	@API_URL=$(API_URL) ./scripts/test-api.sh login

.PHONY: test-users-custom
test-users-custom: ## Test get users endpoint with custom URL
	@chmod +x scripts/test-api.sh
	@API_URL=$(API_URL) ./scripts/test-api.sh users
