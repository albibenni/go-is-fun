.PHONY: build test run clean dev compose migrate-up migrate-down lint generate

build:
	go build -o auth && ./auth

test:
	go test ./...

run:
	go run .

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
