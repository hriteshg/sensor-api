#!/bin/bash

# Load environment variables
set -a
[ -f .env ] && . .env
set +a

# Wait for the Go service to be up and running
until curl -s http://localhost:3333/health > /dev/null; do
  echo "Waiting for the Go service to start..."
  sleep 5
done

# Run the SQL script against PostgreSQL
docker-compose cp seed_test_data.sql postgres:/seed_test_data.sql
export PGPASSWORD="$DB_PASSWORD"
docker-compose exec -T postgres psql -h localhost -U "$DB_USER" -d "$DB_NAME" -a -f /seed_test_data.sql
unset PGPASSWORD
