#!/bin/sh

set -e

echo "Run migrate database"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start app"
exec "$@"