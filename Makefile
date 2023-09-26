.PHONY: help run install

# Define variables
GO := go
APP_NAME := CoverCraft

help:
	@echo "Usage:"
	@echo "  make run             : Run the $(APP_NAME) application"
	@echo "  make install         : Install project dependencies"

run:
	@echo "Starting $(APP_NAME)..."
	@$(GO) run app/main.go

install:
	@echo "Installing project dependencies..."
	@$(GO) mod tidy

compose-build:
	@echo "Starting to build images"
	docker compose build
compose-up:
	@echo "Starting to run built images"
	docker compose up -d

compose-down:
	@echo "Starting to turn services down"
	docker compose down