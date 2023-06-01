.PHONY: dev build test clean

dev:
	@echo "Checking if Docker is running..."
	@if ! docker info > /dev/null 2>&1; then \
		echo "Docker doesn't seem to be running, start Docker and run the command again."; \
		exit 1; \
	fi
	@echo "Building Docker images and starting services..."
	docker-compose build
	docker-compose up -d
	@echo "Running migrations..."
	docker-compose run --rm migrate
	@echo "Development environment is set up and running."

build:
	go build -o bin/api main.go

run: build
	./bin/api

test:
	go test ./...

clean:
	go clean
	docker-compose down

