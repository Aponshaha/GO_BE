-- Seed: 004_seed_orders.sql
-- Description: Seed orders data

BEGIN;

INSERT INTO orders (order_number, customer_id, status, subtotal, tax_amount, shipping_amount, discount_amount, total_amount, currency, created_at, updated_at)
VALUES
    ('ORD-2026-0001', 1, 'delivered'::order_status, 2049.98, 164.00, 15.00, 50.00, 2178.98, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('ORD-2026-0002', 2, 'processing'::order_status, 1329.98, 106.40, 10.00, 0.00, 1446.38, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('ORD-2026-0003', 3, 'pending'::order_status, 79.98, 6.40, 5.00, 0.00, 91.38, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('ORD-2026-0004', 4, 'shipped'::order_status, 1359.97, 108.80, 20.00, 100.00, 1388.77, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('ORD-2026-0005', 5, 'delivered'::order_status, 149.97, 12.00, 10.00, 0.00, 171.97, 'USD', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (order_number) DO NOTHING;

COMMIT;
