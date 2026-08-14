#!/usr/bin/env bash
set -euo pipefail

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-storemanager}"
DB_NAME="${DB_NAME:-storemanager}"
DB_PASSWORD="${DB_PASSWORD:-changeme}"

export PGPASSWORD="$DB_PASSWORD"
PSQL="psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -v ON_ERROR_STOP=1"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_DIR="$(cd "$SCRIPT_DIR/../migrations" && pwd)"

# Pull the table names out of the "-- +goose Down" sections of every migration.
mapfile -t tables < <(
    awk '/^-- \+goose Down/{f=1;next} /^-- \+goose Up/{f=0} f' "$MIGRATIONS_DIR"/*.sql |
        grep -oiE 'DROP TABLE (IF EXISTS )?[a-z_]+' |
        awk '{print $NF}' |
        sort -u
)

if [ ${#tables[@]} -eq 0 ]; then
    echo "No tables found in goose Down sections."
    exit 1
fi

echo "Dropping tables from goose Down sections: ${tables[*]}"

# Reverse so children are dropped before parents (CASCADE covers the rest).
for ((i = ${#tables[@]} - 1; i >= 0; i--)); do
    $PSQL -c "DROP TABLE IF EXISTS ${tables[$i]} CASCADE;"
done

# Drop goose's bookkeeping so "goose up" re-runs the migrations from scratch.
$PSQL -c "DROP TABLE IF EXISTS goose_db_version CASCADE;"

echo "All tables dropped."
