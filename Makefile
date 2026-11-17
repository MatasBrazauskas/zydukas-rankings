include .env

MIGRATIONS_PATH = ./migrations

.PHONY: all
all:
	@air

.PHONY: bruno
bruno:
	@bruno 1>/dev/null 2>&1 &

.PHONY: docker-up
docker-up:
	@docker compose up --build

.PHONY: docker-down
docker-down:
	@docker compose down -v

.PHONY: migrate-create
migration-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: token
token:
	@go run ./cmd/token
