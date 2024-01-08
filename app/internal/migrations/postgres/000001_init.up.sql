BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --

CREATE TABLE public.user
(
    id            BIGINT GENERATED ALWAYS AS IDENTITY,
    uuid          VARCHAR      NOT NULL,
    username      VARCHAR      NOT NULL,
    first_name    VARCHAR,
    last_name     VARCHAR,
    email         VARCHAR      NOT NULL,
    password_hash VARCHAR(256) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;


INSERT INTO public.user (username, email, password_hash)
VALUES ('test', 'admin@admin.ru', '2wq2sq2sq3d');

SELECT *
FROM public.user;