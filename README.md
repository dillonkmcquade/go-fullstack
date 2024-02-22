# Go-Fullstack

This is an opinionated starter template for making a fullstack project with Go, Go-Chi, go templates, HTMX, and Posgresql.

It is minimal and yet provides a working server, development database, and template setup to get going quickly.

It is intended to be extendable and "no magic" in the sense that all the code is included and not hidden from you (except for the librairies of course). A basic structure has been layed out which I think makes sense.

### Whats included

- [Go-Chi](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router for building Go HTTP services.
- [Air](https://github.com/cosmtrek/air) configuration for hot reloading of the development server.
- [HTMX.org](https://htmx.org) - Gives you access to AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML.
- Docker compose configuration for a Postgresql database.
- Makefile for automating common tasks.
- [Godotenv](https://github.com/joho/godotenv) for loading environment variables.

# Getting Started

First, install the required go modules.

```bash
go mod download
make prep #replaces all occurrences of "go-fullstack with your chosen project name"
```

## Requirements

- Docker
- Go 1.21+
- Bash builtins
- Make
- [Staticcheck](https://github.com/dominikh/go-tools) & [gosec](https://github.com/securego/gosec)
- [Air](https://github.com/cosmtrek/air)

## Environment variables

Take a look at the `.env.example`, these are the minimal environment variables required for the database container to start.
Make sure you have docker installed.

- `PG_DATABASE` - The name of the database
- `PG_USERNAME` - Any username created by the user
- `PG_PASSWORD` - The password to the database
- `FORWARD_DB_PORT` - The port that will be exposed for outside access to the container.

## Availabe commands

All of the available commands are in `Makefile`. Utility scripts are located in `scripts/`.

- `make help` Show available commands
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
| | | | parse_templates.go      # Register functions for use in templates
| | | | routes.go               # Register your handlers and middleware to the router here
| | | handlers/                 # Define your http handler functions here
| | | | healthcheck.go
| | | middleware/               # Define your custom middleware here
| | | | cors.go
| public/                       # Put your static CSS and JS files here (if you want to send HTML)
| | assets/                     # Put your images here
| | htmx.js
| scripts/                      # Utility scripts
| | prep*
| | up*
| templates/                    # HTML templates
| | base.html
| test/                         # Tests
| .air.toml
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
