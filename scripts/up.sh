docker compose up -d
trap "docker compose stop" EXIT
air

