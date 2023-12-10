docker compose up -d
trap "docker compose stop && rm -r tmp" EXIT
air

