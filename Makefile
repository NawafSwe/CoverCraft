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

