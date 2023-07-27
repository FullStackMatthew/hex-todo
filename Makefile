# Makefile for hex-todo project

# Variables
APP_NAME := hex-todo
GO := go
DOCKER := docker
DOCKER_COMPOSE := docker-compose

# Targets
.PHONY: all build run clean test

all: build

build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o $(APP_NAME) ./backend

run:
	@echo "Starting $(APP_NAME)..."
	$(GO) run ./backend

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

test:
	@echo "Running tests..."
	$(GO) test -v ./backend/...

docker-build:
	@echo "Building Docker image..."
	$(DOCKER) build -t $(APP_NAME) ./backend

docker-run:
	@echo "Starting Docker container..."
	$(DOCKER_COMPOSE) up

docker-clean:
	@echo "Cleaning up Docker..."
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

