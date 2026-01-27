-- Seed: 003_seed_customers.sql
-- Description: Seed customers data

BEGIN;

INSERT INTO customers (email, first_name, last_name, phone, is_active, created_at, updated_at)
VALUES
    ('john.doe@example.com', 'John', 'Doe', '+1-555-0101', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('jane.smith@example.com', 'Jane', 'Smith', '+1-555-0102', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('bob.johnson@example.com', 'Bob', 'Johnson', '+1-555-0103', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('alice.williams@example.com', 'Alice', 'Williams', '+1-555-0104', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('charlie.brown@example.com', 'Charlie', 'Brown', '+1-555-0105', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (email) DO NOTHING;

COMMIT;
