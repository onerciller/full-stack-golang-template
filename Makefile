# Include tools
include tools/common.mk

# Color definitions
include tools/color.mk

# Include tools
include tools/tool.mk

.PHONY: dev ui-start server-start install build clean swag

# Development commands
install:
	@echo "$(BLUE)Installing dependencies$(CNone)"
	@go mod download
	@cd ui && npm install

dev: 
	@echo "$(BLUE)Starting development environment$(CNone)"
	@make -j 2 ui-start server-start 

swag: $(SWAG)
	@echo "$(BLUE)Generating swagger docs$(CNone)"
	@$(SWAG) init

ui-start:
	@echo "$(BLUE)Starting UI on http://localhost:5173$(CNone)"
	@cd ui && npm run dev

server-start:
	@echo "$(BLUE)Starting Server on http://localhost:3000$(CNone)"
	@go run main.go

build:
	@echo "$(BLUE)Building application$(CNone)"
	@go build -o bin/server main.go
	@cd ui && npm run build

clean:
	@echo "$(BLUE)Cleaning build artifacts$(CNone)"
	@rm -rf bin/
	@cd ui && rm -rf dist/


docker.compose:
	@echo "$(BLUE)Starting Docker Compose$(CNone)"
	@docker compose up -d

docker-compose-down:
	@echo "$(BLUE)Stopping Docker Compose$(CNone)"
	@docker-compose down