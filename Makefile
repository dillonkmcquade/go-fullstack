project_name := go-fullstack
docker_command := podman-compose

build:
	go build -o bin/$(project_name) cmd/$(project_name)/main.go

up: 
	./scripts/up

down:
	$(docker_command) down 

test:
	go test -v ./...

lint:
	staticcheck ./...
	gosec ./...
prep:
	./scripts/prep
