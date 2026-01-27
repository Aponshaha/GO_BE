-- Seed: 001_seed_categories.sql
-- Description: Seed categories data

BEGIN;

INSERT INTO categories (name, slug, description, is_active, sort_order, created_at, updated_at)
VALUES
    ('Electronics', 'electronics', 'Electronic devices and gadgets', true, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Laptops', 'laptops', 'Portable computers and laptops', true, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Accessories', 'accessories', 'Computer and phone accessories', true, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Books', 'books', 'Physical and digital books', true, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Clothing', 'clothing', 'Apparel and fashion items', true, 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT DO NOTHING;

COMMIT;
