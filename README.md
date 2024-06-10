# dbo

## Running the Program Directly

```bash
go mod download
```

then set .env file with

```text
PORT=8080
PGHOST=localhost
PGPORT=5432
PGUSER=postgres
PGPASSWORD=postgres
PGDATABASE=dbo1
TOKEN_HOUR_LIFESPAN=1
API_SECRET=secret
```

then run

```bash
go build
./dbo
```

## Start With Docker

```bash
docker compose up -d --build
```
