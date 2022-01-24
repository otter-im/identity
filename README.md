# Otter Social > Identity Service

A gRPC identity service, for user lookups and login validations.

## Prerequisites

- Redis instance, or cluster
- PostgreSQL instance, or cluster

## Development Setup

```shell
go generate
go run cmd/migrate_db/migrate.go init
go run cmd/service.go
```

## Runtime Envars

- `APP_ENV` (Default: `dev`) - Environment dev or prod
- `SERVICE_HOST` (Default: `0.0.0.0`) - Listening address
- `SERVICE_PORT` (Default: `50050`) - Listening port

#### PostgreSQL
- `POSTGRES_ADDRESS` (Default: `localhost:5432`) - database address, or pgbouncer address
- `POSTGRES_USER` - (Default: none) User to connect to the database
- `POSTGRES_PASSWORD` - (Default: none) Password for the database user
- `POSTGRES_DATABASE` (Default: `otter_identity`) - Database name

#### Redis
- `REDIS_NODES` (Default `localhost:6379`) - Comma delimited list of Redis nodes
- `REDIS_PASSWORD` (Default: none) - Connection password
- `REDIS_DATABASE` (Default: `0`) - Redis DB number

## License

We're MIT.
