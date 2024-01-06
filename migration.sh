#!/bin/bash
source .env

export MIGRATION_DSN=$PG_DSN

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v