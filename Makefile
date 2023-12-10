project_name := go-fullstack

build:
	go build -o bin/$(project_name) cmd/$(project_name)/main.go

up: 
	./scripts/up.sh

down:
	docker compose down 

test:
	go test -v ./...

lint:
	staticcheck ./...
	gosec ./...
