#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until /usr/bin/psql $host -c '\l'; do
  >&2 echo "waiting for db"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd