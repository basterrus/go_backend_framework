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

CREATE TABLE public.roles
(
    id          SERIAL PRIMARY KEY,
    role_name   VARCHAR NOT NULL,
    description VARCHAR
);

CREATE TABLE public.user
(
    id            SERIAL PRIMARY KEY,
    role_id       INT REFERENCES public.roles (id),
    uuid          VARCHAR      NOT NULL,
    username      VARCHAR      NOT NULL,
    first_name    VARCHAR,
    last_name     VARCHAR,
    email         VARCHAR      NOT NULL,
    password_hash VARCHAR(256) NOT NULL,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE public.category
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR NOT NULL,
    description VARCHAR
);

CREATE TABLE public.product
(
    id          SERIAL PRIMARY KEY,
    category_id INT REFERENCES public.category (id) NOT NULL,
    name        VARCHAR                             NOT NULL
);

INSERT INTO public.roles(role_name) VALUES('ADMIN');
INSERT INTO public.roles(role_name) VALUES('USER');
INSERT INTO public.roles(role_name) VALUES('MANAGER');

INSERT INTO public.user(uuid, role_id, username, email, password_hash)
VALUES ('27d99dac-38b0-4fdc-a531-b4a8fa59b427', 1,'admin', 'admin@admin.ru', '12wqdersbtdrshgntdbrv');

INSERT INTO public.user(uuid, role_id, username, email, password_hash)
VALUES ('27d99dac-38b0-4dgc-a531-b4a8fa59b427', 2,'user', 'user@user.ru', '12wqdersbtdrshgntdbrv');

INSERT INTO public.category (name, description) VALUES ('test', 'admin@admin.ru');

INSERT INTO public.product (category_id, name) VALUES (1, '2wq2sq2sq3d');

COMMIT;