# Otter Social > Identity Provider

An OAuth2 identity provider that operates over gRPC.

## Prerequisites

- Redis instance, or cluster
- PostgreSQL instance, or cluster

## Development Setup

```shell
go run cmd/migrate_db/migrate.go init
go run cmd/service.go
```

## Runtime Envars

- `SERVICE_ENV` (Default: `dev`) - Environment dev or prod
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

## Deployments

### Using Docker

Docker BuildKit is preferred where available.

```shell
docker build -t otter-im/identity:latest -f build/package/Dockerfile .
``` 

### Using systemd

Make sure to update `otter-identity.service` with the environment for your Postgres and Redis configuration

```shell
./scripts/build.sh
sudo cp ./dist/otter-identity /usr/bin/otter-identity
sudo cp ./init/otter-identity.service /etc/systemd/system/otter-identity.service
sudo systemctl daemon-reload
sudo systemctl enable --now otter-identity.service
```

### Using Kubernetes
Coming soon.

## License

AGPLv3 License.

See the LICENCE file for the full terms.
