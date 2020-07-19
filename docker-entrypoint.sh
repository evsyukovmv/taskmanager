#!/bin/sh

# shellcheck disable=SC2006
for i in `seq 1 10`;
do
  nc -z "$WAIT_DB_HOST" 5432 && echo DB Host Success && break
  echo -n .
  sleep 1
done

./migrate.linux-amd64 -path db/migrations -database "$DATABASE_URL" up
./app
