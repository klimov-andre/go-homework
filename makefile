export MIGRATIONS_DIR=./migrations
export DB_DSN="host=localhost port=5555 user=user password=password dbname=movies sslmode=disable"
export DB_DSN_TEST="host=localhost port=5556 user=user password=password dbname=movies-test sslmode=disable"

.PHONY: run-dev-infra
run-dev-infra:
	docker-compose up -d postgresql

.PHONY: add-migration
add-migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: migrate
migrate:
	goose -v -dir="${MIGRATIONS_DIR}" postgres ${DB_DSN} up

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: migrate-test
migrate-test:
	goose -v -dir="${MIGRATIONS_DIR}" postgres ${DB_DSN_TEST} up

.PHONY: run-test-infra
run-test-infra:
	docker-compose up -d postgresql-test
	make run-kafka

.PHONY: clear-test-db
clear-test-db:
	docker-compose exec postgresql-test psql -h localhost -p 5432 -U user -d movies-test -c "delete from public.Movie;"
	docker-compose exec postgresql-test psql -h localhost -p 5432 -U user -d movies-test -c "ALTER SEQUENCE movie_id_seq RESTART WITH 1;"

.PHONY: run-kafka
run-kafka:
	docker-compose up -d zookeeper
	docker-compose up -d kafka-1
	docker-compose up -d kafka-ui