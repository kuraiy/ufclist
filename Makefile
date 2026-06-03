migrate-up:
	goose sqlite3 fighters.db -dir db/migrations up

migrate-down:
	goose sqlite3 fighters.db -dir db/migrations down

run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

start:
	./bin/app
