#!/bin/bash

eval $(cat .env | sed 's/#.*//g')\

migrate -path $PWD/sql \
 -database "postgres://$DB_USER:$DB_PASSWORD@$HOST:$PORT/$DB_NAME?sslmode=disable" \
 -verbose up