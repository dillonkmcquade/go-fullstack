# Project

## About

This is an opinionated starter template for making a fullstack project with Go, Go-Chi, go templates, HTMX, and Posgresql.

It is minimal and yet provides a working server, development database, and template setup to get going quickly.

It is intended to be extendable and "no magic" in the sense that all the code is included and not hidden from you (except for the librairies of course). A basic structure has been layed out which I think makes sense.

# Getting Started

## Requirements

- Docker
- Go 1.21+
- Bash builtins
- Make
- [Staticcheck](https://github.com/dominikh/go-tools) & [gosec](https://github.com/securego/gosec)

## Environment variables

Take a look at the `.env.example`, these are the minimal environment variables required for the database container to start.
Make sure you have docker installed.

- `PG_DATABASE` - The name of the database
- `PG_USERNAME` - Any username created by the user
- `PG_PASSWORD` - The password to the database
- `FORWARD_DB_PORT` - The port that will be exposed for outside access to the container.

## Availabe commands

All of the available commands are in `Makefile`. Utility scripts are located in `scripts/`.

- `make up` Builds and starts the postgresql container and runs 'go run' on the main package. It runs `docker compose stop` on exit.
- `make build` builds the binary and puts it in bin/
- `make test` runs all of the tests
- `make down` Removes the docker container
- `make lint` Run staticcheck and gosec

## Project structure

```
go-fullstack/
| bin/
| | go-fullstack*               # The binary
| cmd/
| | go-fullstack/
| | | main.go                   # Entrypoint to your application
| internal/
| | | app/
| | | | app.go                  # Contains some server configuration options
| | | | parse_templates.go
| | | | routes.go               # Register your handlers and middleware to the router here
| | | handlers/                 # Define your http handler functions here
| | | | healthcheck.go
| | | middleware/               # Define your custom middleware here
| | | | cors.go
| public/                       # Put your static CSS and JS files here (if you want to send HTML)
| | assets/                     # Put your images here
| | htmx.js
| scripts/                      # Utility scripts
| | up.sh*
| templates/                    # HTML templates
| | base.html
| test/                         # Tests
| .env
| .env.example
| .gitignore
| LICENSE
| Makefile
| README.md
| docker-compose.yml
| go.mod
| go.sum

```

## License

[MIT](./LICENSE) license
