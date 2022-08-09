export MIGRATIONS_DIR=./migrations
export DB_DSN="host=localhost port=5555 user=user password=password dbname=movies sslmode=disable"

.PHONY: run-dev-infra
run-dev-infra:
	docker-compose up -d postgresql

.PHONY: add-migration
add-migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: migrate
migrate:
	goose -v -dir="${MIGRATIONS_DIR}" postgres ${DB_DSN} up