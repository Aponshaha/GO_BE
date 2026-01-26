# Database Migrations Guide

This directory contains SQL migration files for the e-commerce/inventory system.

## Migration Files

### 001_create_base_tables.sql

Creates all base tables with:

- Proper data types (timestamps, decimals, enums)
- Foreign key constraints
- Check constraints
- Triggers for auto-updating timestamps
- Triggers for stock management
- Triggers for rating calculations

### 002_create_indexes.sql

Creates comprehensive indexes for:

- Foreign keys (faster JOINs)
- Frequently queried columns
- Search operations (full-text, trigram)
- Soft delete queries
- Composite indexes for complex queries
- Partial indexes for specific patterns

### 003_add_extensions.sql

Adds PostgreSQL extensions:

- `pg_trgm` - For fuzzy text search
- `btree_gin` - For GIN indexes on common types

## Running Migrations

### Option 1: Using Go Migration Runner (Recommended) ⭐

```bash
# Check migration status
go run cmd/migrate/main.go -command=status

# Run all pending migrations
go run cmd/migrate/main.go -command=up

# Rollback last migration
go run cmd/migrate/main.go -command=down
```

The migration runner:

- ✅ Tracks which migrations have been applied
- ✅ Only runs pending migrations
- ✅ Uses transactions (safe rollback on error)
- ✅ Shows clear status

### Option 2: Manual (psql)

```bash
# Connect to your Docker database
psql -h localhost -p 5432 -U postgres -d ecom

# Run migrations in order
\i migrations/001_create_base_tables.sql
\i migrations/003_add_extensions.sql
\i migrations/002_create_indexes.sql
```

### Option 3: Using psql from command line

```bash
# Run all migrations
psql -h localhost -p 5433 -U postgres -d ecom -f migrations/001_create_base_tables.sql
psql -h localhost -p 5433 -U postgres -d ecom -f migrations/003_add_extensions.sql
psql -h localhost -p 5433 -U postgres -d ecom -f migrations/002_create_indexes.sql
```

## Migration Best Practices

1. **Always use transactions**: Each migration should be wrapped in `BEGIN;` and `COMMIT;`
2. **Run in order**: Migrations must be run sequentially
3. **Test first**: Test migrations on a copy of production data
4. **Backup**: Always backup before running migrations
5. **Idempotent**: Use `IF NOT EXISTS` where possible
6. **Version control**: Track which migrations have been applied

## Creating New Migrations

1. Number sequentially: `004_description.sql`, `005_description.sql`
2. Include description comment at top
3. Use transactions
4. Test thoroughly
5. Document in this README

## Index Strategy Explained

### Why Indexes?

Indexes speed up queries but slow down writes. We use them strategically:

1. **Foreign Key Indexes**: Every foreign key should be indexed
2. **Search Indexes**: Columns used in WHERE clauses
3. **Sort Indexes**: Columns used in ORDER BY
4. **Composite Indexes**: Multiple columns queried together
5. **Partial Indexes**: Index only specific rows (e.g., active products)

### Index Types Used

- **B-tree**: Default, good for most queries
- **GIN (Generalized Inverted Index)**: For full-text search, arrays, JSONB
- **Partial**: Only index rows matching a condition

### Example Queries That Benefit

```sql
-- Fast: Uses idx_products_category_id
SELECT * FROM products WHERE category_id = 5;

-- Fast: Uses idx_products_status_price
SELECT * FROM products
WHERE status = 'active'
ORDER BY price ASC;

-- Fast: Uses idx_products_name_trgm (trigram)
SELECT * FROM products
WHERE name ILIKE '%laptop%';

-- Fast: Uses idx_products_active_only (partial)
SELECT * FROM products
WHERE deleted_at IS NULL AND status = 'active';
```

## Rollback Strategy

For production, always create rollback migrations:

```sql
-- 004_add_feature.sql (up)
-- 004_add_feature_rollback.sql (down)
```

## Monitoring Index Usage

```sql
-- Check index usage
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan as index_scans
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;

-- Find unused indexes
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan
FROM pg_stat_user_indexes
WHERE idx_scan = 0
AND indexname NOT LIKE 'pg_toast%';
```
