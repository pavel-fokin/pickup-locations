# Pickup Locations

`pickup-locations` takes the source and a list of destinations
and returns a list of routes between source and each destination.

## Requirements

- Go 1.20

## Makefile

`Makefile` can be used to run a few simple commands.

```sh
$ make help
Makefile help:

run               Run service.
tests             Run unittests.
swagger           Generate swagger documentation.
                  Requires installed https://github.com/swaggo/swag
```

## API Docs

Service has a generated Swagger documentation. It is available `http://service-url/swagger/`

## Run

### Local

```sh
$ go mod download
$ make run
2023/03/17 09:03:26 Error loading .env file
2023/03/17 09:03:26 Starting pickup-locations... 8080
```

### Environment Variables

Service supports configuration with environment variables either directly, or with `*.env` file.

```
// General.
PICKUP_LOCATIONS_ENV_FILE=<path-to-env-file>

// Server.
PORT=<port>
PICKUP_LOCATIONS_HANDLER_TIMEOUT=<timeout-in-seconds>
PICKUP_LOCATIONS_SHUTDOWN_TIMEOUT=<timeout-in-seconds>

// OSRM client.
OSRM_CLIENT_TIMEOUT_SECONDS=<timeout-in-seconds>
```

## Deployment

Service uses `Google Cloud` to manage deployments and infrastructure.

- `Cloud Build` - CI/CD 
- `Artifacts Registry` - a repository with Docker images
- `Cloud Run` runs `pickup-locations` service
