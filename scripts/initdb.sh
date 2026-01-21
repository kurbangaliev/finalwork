#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER finalwork;
	CREATE DATABASE finalwork;
	GRANT ALL PRIVILEGES ON DATABASE finalwork TO finalwork;
	\c finalwork;
EOSQL
