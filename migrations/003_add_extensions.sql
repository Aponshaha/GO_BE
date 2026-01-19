-- Migration: 003_add_extensions.sql
-- Description: Adds PostgreSQL extensions for advanced features
-- Created: 2024-01-02

BEGIN;

-- Enable pg_trgm extension for fuzzy text search
-- This allows efficient LIKE queries and full-text search
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Enable btree_gin for GIN indexes on common data types
-- Improves performance of certain index types
CREATE EXTENSION IF NOT EXISTS btree_gin;

COMMIT;


