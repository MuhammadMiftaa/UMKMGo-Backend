include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

api:
	go run ./cmd/api/main.go

test:
	go test -v -cover -race ./internal/service/...

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