-- PostgreSQL initialization script for simple-blog
-- Run this script as a PostgreSQL superuser before starting the application.
-- Usage: psql -U postgres -f scripts/init-db.sql

-- Create database (skip if it already exists)
SELECT 'CREATE DATABASE simple_blog'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'simple_blog')\gexec

-- Create application user (skip if it already exists).
-- IMPORTANT: Replace 'CHANGE_ME_STRONG_PASSWORD' with a strong password before running.
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'blog_user') THEN
        CREATE USER blog_user WITH PASSWORD 'CHANGE_ME_STRONG_PASSWORD';
    END IF;
END
$$;

-- Grant privileges on the database
GRANT ALL PRIVILEGES ON DATABASE simple_blog TO blog_user;

-- Connect to the application database to set up schema permissions
\connect simple_blog

-- Grant schema privileges so the application user can create tables
GRANT ALL ON SCHEMA public TO blog_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO blog_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO blog_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO blog_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO blog_user;
