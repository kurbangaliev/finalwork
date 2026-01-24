#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER finalwork;
	CREATE DATABASE finalwork;
	GRANT ALL PRIVILEGES ON DATABASE finalwork TO finalwork;
	\c finalwork;
  CREATE TABLE IF NOT EXISTS public.users
  (
      id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
      created_at timestamp with time zone,
      updated_at timestamp with time zone,
      deleted_at timestamp with time zone,
      login character varying(255) COLLATE pg_catalog."default",
      password character varying(255) COLLATE pg_catalog."default",
      CONSTRAINT users_pkey PRIMARY KEY (id)
  )

  TABLESPACE pg_default;

  ALTER TABLE IF EXISTS public.users
      OWNER to postgres;
  -- Index: idx_users_deleted_at

  -- DROP INDEX IF EXISTS public.idx_users_deleted_at;

  CREATE INDEX IF NOT EXISTS idx_users_deleted_at
      ON public.users USING btree
      (deleted_at ASC NULLS LAST)
      WITH (fillfactor=100, deduplicate_items=True)
      TABLESPACE pg_default;
EOSQL
