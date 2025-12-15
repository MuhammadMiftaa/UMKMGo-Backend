include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

api:
	go run ./cmd/api/main.go

migrate:
	@if [ -z "$(to)" ]; then \
		goose up; \
	else \
		goose up-to $(to); \
	fi

migration:
	@goose create $(name) sql

rollback:
	@if [ -z "$(to)" ]; then \
		goose down; \
	else \
		goose down-to $(to); \
	fi

migration-status:
	@goose status

seeder:
	@goose -dir ./config/db/seeder create $(name) sql

seed:
	@goose -dir ./config/db/seeder -no-versioning up

seed-reset:
	@goose -dir ./config/db/seeder -no-versioning reset

test:
	go test -cover ./internal/service/...

test-coverage:
	go test -coverprofile=coverage.out ./internal/service/...
	go tool cover -html=coverage.out

test-verbose:
	go test -v -cover ./internal/service/...