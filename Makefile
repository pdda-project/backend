build:
	@go build -o bin/server .

migrate:
	@go run cmd/migrate/main.go


seed:
	@go run cmd/seed/main.go

run:
	@go run cmd/migrate/main.go && \
		go build -o bin/server && \
		./bin/server
