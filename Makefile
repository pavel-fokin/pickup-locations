.PHONY: all
all: help

.PHONY: help
help:
	@echo 'Makefile help:                                                         '
	@echo '                                                                       '
	@echo 'run               Run service.                                         '
	@echo 'lint              Run linter.                                          '
	@echo 'tests             Run unittests.                                       '
	@echo 'swagger           Generate swagger documentation.                      '
	@echo '                  Requires installed https://github.com/swaggo/swag    '

.PHONY: lint
lint:
	@staticcheck ./...

.PHONY: tests
tests:
	@go test -coverprofile=coverage.out ./... -count=1

.PHONY: swagger
swagger:
	swag init -d cmd/pickup-locations-server,internal/server

.PHONY: run
run:
	@go run cmd/pickup-locations-server/main.go