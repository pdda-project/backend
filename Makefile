build:
	@go build -o bin/server .

migrate:
	@go run cmd/migrate/main.go