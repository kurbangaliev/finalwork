#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -U finalwork -d finalwork -f /docker-entrypoint-initdb.d/finalwork.sql;
