# Go parameters
MAIN_PATH=cmd/todoapi/main.go
BINARY_NAME=$(BINARY_PATH)/todoapi
BINARY_PATH=bin

run: ## Build and run the application
	[ -f ./.env ] || cp .env.tpl .env
	go build -o $(BINARY_NAME) -race $(MAIN_PATH)
	./$(BINARY_NAME)

test: ## Run unittests
	go test -race -v -timeout=10s ./...

clean: ## Remove previous build
	go clean $(MAIN_PATH)
	rm -f $(BINARY_PATH)/*

dep: ## Get the dependencies
	go mod download

lint: ## Lint Golang files
	@golint ./...

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ./...
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go build -o $(BINARY_NAME) -race $(MAIN_PATH)

docker-start-dev: ## Start Docker DEV containers
	docker-compose -f docker-compose.dev.yml up -d --remove-orphans

docker-stop-dev: ## Stop Docker DEV containers
	docker-compose -f docker-compose.dev.yml stop
	docker-compose -f docker-compose.dev.yml rm -f

docker-start: ## Start Docker production containers
	docker-compose -f docker-compose.yml up -d

docker-stop: ## Stop Docker production containers
	docker-compose -f docker-compose.yml stop
	docker-compose -f docker-compose.yml rm -f

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'