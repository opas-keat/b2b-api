#!/bin/bash

if [ $# -eq 0 ]
  then
    echo "Require migration name on first argument"
fi

MIGRATION_NAME=$1

migrate create -ext sql -dir sql $MIGRATION_NAME
