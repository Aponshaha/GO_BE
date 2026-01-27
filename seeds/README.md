# Database Seeds

Seed files for populating the database with mock data.

## Structure

Seed files are SQL files that follow a naming convention: `NNN_seed_description.sql`

- `001_seed_categories.sql` - Seed categories
- `002_seed_products.sql` - Seed products
- `003_seed_customers.sql` - Seed customers
- `004_seed_orders.sql` - Seed orders

## Usage

### Run all seeds

```bash
go run cmd/seed/main.go
```

### Clear and reseed (delete all existing data first)

```bash
go run cmd/seed/main.go -clear
```

## Adding New Seeds

1. Create a new SQL file following the naming convention: `NNN_seed_description.sql`
2. Keep the number sequential
3. Use `ON CONFLICT DO NOTHING` to handle duplicates gracefully
4. Wrap changes in `BEGIN;` and `COMMIT;` transactions
5. The seed runner will execute files in alphabetical order

### Example

```sql
-- Seed: 005_seed_coupons.sql
-- Description: Seed coupon data

BEGIN;

INSERT INTO coupons (code, name, description, discount_type, discount_value, is_active, starts_at, expires_at, created_at, updated_at)
VALUES
    ('SUMMER2024', 'Summer Sale', '20% off summer collection', 'percentage', 20.00, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '30 days', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (code) DO NOTHING;

COMMIT;
```

## Best Practices

1. **Use idempotent statements** - Use `ON CONFLICT DO NOTHING` or `IF NOT EXISTS`
2. **Keep it organized** - One type of data per file
3. **Use transactions** - Wrap all changes in `BEGIN;` and `COMMIT;`
4. **Document your seeds** - Add comments explaining what the seed does
5. **Test before running** - Test seeds on a copy of your database first
6. **Separate from code** - Keep SQL in files, not in Go code
