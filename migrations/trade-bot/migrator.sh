#!/bin/bash

set -o errexit

readonly CMD=$1
readonly CONF=migrations/trade-bot/dbconfig.yml

if [ "$CMD" == "up" ]
  then
    echo "Applying all migrations"
    # --dryrun - means not apply migrations but just print them
    sql-migrate up --config="$CONF" --env="$ENV" --dryrun
    sql-migrate up --config="$CONF" --env="$ENV"
    echo "all migrations are applied"

elif [ "$CMD" == "down" ]
  then
    echo "Rolling back last migration"
    # --dryrun - means not apply migrations but just print them
    sql-migrate down --limit=1 --config="$CONF" --env="$ENV" --dryrun
    sql-migrate down --limit=1 --config="$CONF" --env="$ENV"
    echo "migration rolled back"

elif [ "$CMD" == "status" ]
  then
    echo "Migrations status"
    sql-migrate status --config="$CONF" --env="$ENV"

else
  echo "Incorrect command passed. Use "up" or "down""
fi
