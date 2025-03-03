.PHONY: run
run:
	go run ./cmd/api

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: migrate-up
migrate-up:
	goose -dir=migrations sqlite ./data-sharing.db up

.PHONY: migrate-down
migrate-down:
	goose -dir=migrations sqlite ./data-sharing.db down

.PHONY: generate
generate:
	goose -dir=migrations create ${name} sql
