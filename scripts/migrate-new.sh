#!/bin/bash

set -o errexit

migrationText=$(cat << EOF
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
EOF
)

read -p "Enter migration name: " name

echo "$migrationText" > "./migrations/trade-bot/__$name.sql"
