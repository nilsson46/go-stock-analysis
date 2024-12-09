#!/bin/sh

host="$1"
shift
port="$1"
shift
sql_file="$1"
shift
cmd="$@"

until pg_isready -h "$host" -p "$port"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing SQL file"
psql -h "$host" -p "$port" -U postgres -f "$sql_file"

>&2 echo "Executing command"
exec $cmd