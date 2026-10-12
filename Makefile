COMPOSE=docker compose -f docker-compose.dev.yml

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up api

down:
	$(COMPOSE) down