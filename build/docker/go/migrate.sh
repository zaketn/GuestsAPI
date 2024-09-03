#!/bin/sh
./migrate -database "postgres://${PG_USER}:${PG_PASSWORD}@db:5432/${PG_DB_NAME}?sslmode=disable" -path ./internal/db/migrations up

exec "$@"