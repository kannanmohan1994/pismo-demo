# Pismo API

Backend service that exposes account and transaction endpoints on top of a Postgres database. The project is containerised so you can bring everything up with Docker, but you can also run it directly with Go if you prefer.

## Prerequisites

- [REQUIRED] Docker 24+ (or Docker Desktop) and Docker Compose v2 (`docker compose` CLI).
- [OPTIONAL] GNU Make (bundled on macOS/Linux; install via [chocolatey](https://community.chocolatey.org/packages/make) on Windows).
- [OPTIONAL] Go 1.23 if you want to run the binary outside containers.

## Configuration

1. Copy the sample environment file:

   - Copy env.sample --> .env

If you need to change container credentials, edit `docker.env` first and then reflect the same values in `.env`.

## Quick start

1. **Build and start containers**
   ```bash
   docker compose up -d --build
   ```

   This launches a Postgres 16 instance and the API container using the configuration declared in `docker-compose.yml`.
2. **Verify services**
   ```bash
   docker compose ps
   docker compose logs -f pismo   # optional
   ```

   Wait until the `postgres` container reports healthy.
3. **Run database migrations once**
   - `make migrate-up` OR the command mentioned under make migrate-up in Makefile if make command is not supported.

Once migrations succeed, the API is available at `http://localhost:9001`. Open `http://localhost:9001/swagger/index.html` for interactive docs (disabled in production mode).

### Common tasks

- Run tests: `make test`
- Generate Swagger spec: `make swagger`
- Stop and remove containers: `docker compose down`
- Drop the schema (dev only): `make migrate-down`

## Troubleshooting

- **Permission errors writing to `/data/postgres`**: adjust the volume in `docker-compose.yml` to a path your user owns, for example `./.postgres:/var/lib/postgresql/data`

## Project layout

- `app/` – Gin HTTP server bootstrap and routing.
- `internal/` – Domain handlers, validation, use cases, repositories, and mocks.
- `db/migrations/` – SQL migrations executed via `make migrate-up`.
- `docs/` – Generated Swagger specification.
