-- Migration: 001_initial_schema
-- Description: Initial schema for simple-blog (users, posts, comments)
-- Note: GORM AutoMigrate handles table creation automatically.
--       This file documents the expected schema for reference and manual recovery.

CREATE TABLE IF NOT EXISTS users (
    id         BIGSERIAL    PRIMARY KEY,
    username   VARCHAR(100) NOT NULL UNIQUE,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   TEXT         NOT NULL,
    bio        VARCHAR(500),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS posts (
    id         BIGSERIAL    PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    content    TEXT,
    excerpt    VARCHAR(500),
    author_id  BIGINT       NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category   VARCHAR(100),
    tags       VARCHAR(255),
    published  BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS comments (
    id         BIGSERIAL   PRIMARY KEY,
    content    TEXT        NOT NULL,
    author_id  BIGINT      NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id    BIGINT      NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
