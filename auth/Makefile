.PHONY: build test run clean dev compose migrate-up migrate-down lint generate

build:
	go build -o auth && ./auth

test:
	go test ./...

run:
	go run .

generate:
	sqlc generate

migrate-up:
	goose postgres postgresql://albibenni:password@localhost:5432/go-db up -dir ./sql/schema

migrate-down:
	goose postgres postgresql://albibenni:password@localhost:5432/go-db down -dir ./sql/schema

dev:
	air # or use go run with file watching

clean:
	rm -rf server

compose:
	docker compose down && docker compose up -d

install:
	go mod download

lint:
	golangci-lint run

docker-build:
	docker build -t myapp .

deploy:
	./scripts/deploy.sh
