docker compose up -d
trap "docker compose stop" EXIT
go run "cmd/$1/main.go"

