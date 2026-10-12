COMPOSE=docker compose -f docker-compose.dev.yml

.PHONY: build up down format

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up api

down:
	$(COMPOSE) down

format:
	$(COMPOSE) run --rm api go fmt ./...
