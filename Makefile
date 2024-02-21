.DEFAULT_GOAL := help
project_name := go-fullstack
docker_command := podman-compose

.PHONY: build
## Compile a binary 
build:
	go build -o bin/$(project_name) cmd/$(project_name)/main.go

.PHONY: up
## Start the server and database
up: 
	./scripts/up

.PHONY: down
## Stop the server and database
down:
	$(docker_command) down 

.PHONY: test
## Run the tests
test:
	go test -v ./...

.PHONY: lint
## Lint using staticcheck and gosec
lint:
	staticcheck ./...
	gosec ./...

.PHONY: prep
## Replace all occurrences of 'go-fullstack' with your project's name
prep:
	./scripts/prep

# See <https://gist.github.com/klmr/575726c7e05d8780505a> for explanation.
.PHONY: help
help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)";echo;sed -ne"/^## /{h;s/.*//;:d" -e"H;n;s/^## //;td" -e"s/:.*//;G;s/\\n## /---/;s/\\n/ /g;p;}" ${MAKEFILE_LIST}|LC_ALL='C' sort -f|awk -F --- -v n=$$(tput cols) -v i=19 -v a="$$(tput setaf 6)" -v z="$$(tput sgr0)" '{printf"%s%*s%s ",a,-i,$$1,z;m=split($$2,w," ");l=n-i;for(j=1;j<=m;j++){l-=length(w[j])+1;if(l<= 0){l=n-i-length(w[j])-1;printf"\n%*s ",-i," ";}printf"%s ",w[j];}printf"\n";}'|more $(shell test $(shell uname) == Darwin && echo '-Xr')

