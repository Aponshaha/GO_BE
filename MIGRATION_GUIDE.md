# Database Migration & Indexing Guide

## 🎯 What We've Built

A complete e-commerce/inventory system with:

- ✅ Proper database schema with correct data types
- ✅ Comprehensive indexing strategy
- ✅ Migration system for version control
- ✅ Triggers for automatic updates
- ✅ Foreign key constraints
- ✅ Soft delete support

## 📁 Migration Files

1. **001_create_base_tables.sql** - All tables, constraints, triggers
2. **002_create_indexes.sql** - All indexes for performance
3. **003_add_extensions.sql** - PostgreSQL extensions

## 🚀 Quick Start

### Step 1: Run Migrations

```bash
# Check status
go run cmd/migrate/main.go -command=status

# Run all migrations
go run cmd/migrate/main.go -command=up
```

### Step 2: Verify

```bash
# Connect to database
psql -h localhost -p 5433 -U postgres -d ecom

# Check tables
\dt

# Check indexes
\di

# Check a specific table structure
\d products
```

## 📊 Indexing Strategy Explained

### Why Indexes?

Indexes make queries faster by creating a "table of contents" for your data.

**Without index:**

```sql
SELECT * FROM products WHERE category_id = 5;
-- Scans entire table (slow for large tables)
```

**With index:**

```sql
-- Uses idx_products_category_id
SELECT * FROM products WHERE category_id = 5;
-- Direct lookup (fast!)
```

### Index Types We Use

1. **B-tree Indexes** (default)

   - For: Equality, range queries, sorting
   - Example: `WHERE category_id = 5`, `ORDER BY price`

2. **GIN Indexes** (Generalized Inverted Index)

   - For: Full-text search, arrays, JSONB
   - Example: `WHERE name ILIKE '%laptop%'`

3. **Partial Indexes**
   - Only index specific rows
   - Smaller, faster
   - Example: Only index active products

### Index Categories

#### 1. Foreign Key Indexes

Every foreign key gets an index for faster JOINs:

```sql
CREATE INDEX idx_products_category_id ON products(category_id);
```

#### 2. Search Indexes

Columns used in WHERE clauses:

```sql
CREATE INDEX idx_products_status ON products(status);
```

#### 3. Composite Indexes

Multiple columns queried together:

```sql
CREATE INDEX idx_products_category_status
    ON products(category_id, status);
```

#### 4. Partial Indexes

Only index specific rows:

```sql
CREATE INDEX idx_products_active_only
    ON products(id, name, price)
    WHERE deleted_at IS NULL AND status = 'active';
```

## 🔍 Query Performance Examples

### Fast Queries (Using Indexes)

```sql
-- Uses idx_products_category_id
SELECT * FROM products WHERE category_id = 5;

-- Uses idx_products_status_price
SELECT * FROM products
WHERE status = 'active'
ORDER BY price ASC;

-- Uses idx_products_name_trgm (trigram search)
SELECT * FROM products
WHERE name ILIKE '%laptop%';

-- Uses idx_orders_customer_status
SELECT * FROM orders
WHERE customer_id = 10 AND status = 'pending';
```

### Slow Queries (No Index)

```sql
-- No index on description (unless full-text)
SELECT * FROM products WHERE description LIKE '%something%';

-- No index on calculated fields
SELECT * FROM products WHERE price * 1.1 > 100;
```

## 📈 Monitoring Index Usage

### Check Which Indexes Are Used

```sql
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan as times_used
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;
```

### Find Unused Indexes

```sql
SELECT
    schemaname,
    tablename,
    indexname
FROM pg_stat_user_indexes
WHERE idx_scan = 0
AND indexname NOT LIKE 'pg_toast%';
```

### Check Index Size

```sql
SELECT
    tablename,
    indexname,
    pg_size_pretty(pg_relation_size(indexname::regclass)) as size
FROM pg_indexes
WHERE schemaname = 'public'
ORDER BY pg_relation_size(indexname::regclass) DESC;
```

## 🎓 Practice Exercises

### Exercise 1: Add a New Index

Create a migration file `004_add_product_search_index.sql`:

```sql
BEGIN;

-- Add index for searching products by SKU
CREATE INDEX IF NOT EXISTS idx_products_sku
    ON products(sku)
    WHERE deleted_at IS NULL;

COMMIT;
```

Run it:

```bash
go run cmd/migrate/main.go -command=up
```

### Exercise 2: Analyze Query Performance

```sql
-- Enable query timing
\timing

-- Run a query
EXPLAIN ANALYZE
SELECT * FROM products
WHERE category_id = 5 AND status = 'active'
ORDER BY price ASC;
```

Look for:

- `Index Scan` = Using index ✅
- `Seq Scan` = Full table scan ❌

### Exercise 3: Create a Composite Index

If you frequently query:

```sql
SELECT * FROM orders
WHERE customer_id = 10
AND status = 'pending'
ORDER BY created_at DESC;
```

Create index:

```sql
CREATE INDEX idx_orders_customer_status_date
    ON orders(customer_id, status, created_at DESC);
```

## 🔄 Migration Best Practices

1. **Always use transactions**

   ```sql
   BEGIN;
   -- your changes
   COMMIT;
   ```

2. **Test on copy first**

   ```bash
   # Create test database
   createdb ecom_test
   psql ecom_test < migrations/001_create_base_tables.sql
   ```

3. **Backup before migrations**

   ```bash
   pg_dump -h localhost -p 5433 -U postgres -d ecom > backup.sql
   ```

4. **Number sequentially**

   - `001_`, `002_`, `003_` etc.

5. **Make idempotent**
   ```sql
   CREATE INDEX IF NOT EXISTS ...
   CREATE TABLE IF NOT EXISTS ...
   ```

## 🐛 Common Issues

### Issue: Migration already applied

**Solution**: The migration runner tracks applied migrations. Check status:

```bash
go run cmd/migrate/main.go -command=status
```

### Issue: Index already exists

**Solution**: Use `IF NOT EXISTS`:

```sql
CREATE INDEX IF NOT EXISTS idx_name ...
```

### Issue: Slow queries

**Solution**:

1. Check if index exists: `\di`
2. Use `EXPLAIN ANALYZE` to see query plan
3. Add missing indexes

## 📚 Next Steps

1. **Run the migrations** on your Docker database
2. **Test queries** and see indexes in action
3. **Create new migrations** for additional features
4. **Monitor index usage** to optimize
5. **Update Go models** to match new schema

## 🎯 Key Takeaways

- **Indexes** = Faster queries, slower writes
- **Foreign keys** should always be indexed
- **Composite indexes** help multi-column queries
- **Partial indexes** save space and improve performance
- **Monitor** index usage to find unused indexes
- **Migrations** help version control your database schema

Happy migrating! 🚀
