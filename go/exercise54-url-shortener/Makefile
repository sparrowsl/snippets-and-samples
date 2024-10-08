dev:
	npm run dev --prefix client/

run:
	go run ./cmd/api/

generate:
	sqlc generate

migrateup:
	goose -dir="./internal/data/schema/" sqlite3 ./breve.db up
	
migratedown:
	goose -dir="./internal/data/schema/" sqlite3 ./breve.db down
