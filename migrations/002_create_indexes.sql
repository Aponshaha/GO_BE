-- Migration: 002_create_indexes.sql
-- Description: Creates indexes for optimal query performance
-- Created: 2024-01-02



-- ============================================================================
-- PRIMARY KEY INDEXES (automatically created, but documenting for reference)
-- ============================================================================
-- All PRIMARY KEY constraints automatically create unique indexes
-- categories_pkey, products_pkey, customers_pkey, orders_pkey, etc.

-- ============================================================================
-- FOREIGN KEY INDEXES
-- ============================================================================
-- Index foreign keys for faster JOINs and constraint checks

-- Categories
CREATE INDEX IF NOT EXISTS idx_categories_parent_id ON categories(parent_id) WHERE parent_id IS NOT NULL;

-- Products
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_status ON products(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_products_is_featured ON products(is_featured) WHERE is_featured = true AND deleted_at IS NULL;

-- Product Images
CREATE INDEX IF NOT EXISTS idx_product_images_product_id ON product_images(product_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_product_images_is_primary ON product_images(product_id, is_primary) WHERE is_primary = true;

-- Inventory Movements
CREATE INDEX IF NOT EXISTS idx_inventory_movements_product_id ON inventory_movements(product_id);
CREATE INDEX IF NOT EXISTS idx_inventory_movements_reference ON inventory_movements(reference_type, reference_id) WHERE reference_type IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_inventory_movements_created_at ON inventory_movements(created_at DESC);

-- Customer Addresses
CREATE INDEX IF NOT EXISTS idx_customer_addresses_customer_id ON customer_addresses(customer_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_customer_addresses_is_default ON customer_addresses(customer_id, is_default) WHERE is_default = true;

-- Orders
CREATE INDEX IF NOT EXISTS idx_orders_customer_id ON orders(customer_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at DESC) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_orders_order_number ON orders(order_number);

-- Order Items
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_order_items_product_id ON order_items(product_id) WHERE deleted_at IS NULL;

-- Payments
CREATE INDEX IF NOT EXISTS idx_payments_order_id ON payments(order_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_payments_transaction_id ON payments(transaction_id) WHERE transaction_id IS NOT NULL;

-- Order Coupons
CREATE INDEX IF NOT EXISTS idx_order_coupons_order_id ON order_coupons(order_id);
CREATE INDEX IF NOT EXISTS idx_order_coupons_coupon_id ON order_coupons(coupon_id);

-- Shipping Details
CREATE INDEX IF NOT EXISTS idx_shipping_details_order_id ON shipping_details(order_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_shipping_details_tracking_number ON shipping_details(tracking_number) WHERE tracking_number IS NOT NULL;

-- Product Reviews
CREATE INDEX IF NOT EXISTS idx_product_reviews_product_id ON product_reviews(product_id) WHERE deleted_at IS NULL AND is_approved = true;
CREATE INDEX IF NOT EXISTS idx_product_reviews_customer_id ON product_reviews(customer_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_product_reviews_rating ON product_reviews(product_id, rating) WHERE deleted_at IS NULL AND is_approved = true;

-- Wishlists
CREATE INDEX IF NOT EXISTS idx_wishlists_customer_id ON wishlists(customer_id) WHERE deleted_at IS NULL;

-- Wishlist Items
CREATE INDEX IF NOT EXISTS idx_wishlist_items_wishlist_id ON wishlist_items(wishlist_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_wishlist_items_product_id ON wishlist_items(product_id) WHERE deleted_at IS NULL;

-- ============================================================================
-- UNIQUE INDEXES (for frequently queried unique columns)
-- ============================================================================

-- Products: SKU and slug are already UNIQUE (creates unique indexes automatically)
-- But we can add composite unique indexes if needed

-- ============================================================================
-- SEARCH & FILTERING INDEXES
-- ============================================================================

-- Products: Full-text search on name and description
CREATE INDEX IF NOT EXISTS idx_products_name_trgm ON products USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_products_description_trgm ON products USING gin(description gin_trgm_ops);

-- Note: Requires pg_trgm extension for trigram search
-- Run: CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Products: Price range queries
CREATE INDEX IF NOT EXISTS idx_products_price ON products(price) WHERE deleted_at IS NULL AND status = 'active';

-- Products: Stock level queries
CREATE INDEX IF NOT EXISTS idx_products_stock_low ON products(id, stock_quantity) 
    WHERE stock_quantity <= low_stock_threshold AND deleted_at IS NULL;

-- Products: Rating queries
CREATE INDEX IF NOT EXISTS idx_products_rating ON products(rating_average DESC) 
    WHERE deleted_at IS NULL AND status = 'active' AND rating_count > 0;

-- Products: Popular products (by view count)
CREATE INDEX IF NOT EXISTS idx_products_views ON products(view_count DESC) 
    WHERE deleted_at IS NULL AND status = 'active';

-- Customers: Email lookup (already UNIQUE, but documenting)
-- Customers: Name search
CREATE INDEX IF NOT EXISTS idx_customers_name ON customers USING gin((first_name || ' ' || last_name) gin_trgm_ops);

-- Orders: Date range queries (already have created_at index, but adding composite)
CREATE INDEX IF NOT EXISTS idx_orders_customer_date ON orders(customer_id, created_at DESC) 
    WHERE deleted_at IS NULL;

-- Orders: Status and date for reporting
CREATE INDEX IF NOT EXISTS idx_orders_status_date ON orders(status, created_at DESC) 
    WHERE deleted_at IS NULL;

-- ============================================================================
-- SOFT DELETE INDEXES
-- ============================================================================
-- Index deleted_at for efficient soft delete queries

CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON categories(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_products_deleted_at ON products(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_customers_deleted_at ON customers(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_orders_deleted_at ON orders(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_order_items_deleted_at ON order_items(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_coupons_deleted_at ON coupons(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_product_images_deleted_at ON product_images(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_product_reviews_deleted_at ON product_reviews(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_shipping_details_deleted_at ON shipping_details(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_wishlists_deleted_at ON wishlists(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_wishlist_items_deleted_at ON wishlist_items(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_customer_addresses_deleted_at ON customer_addresses(deleted_at) WHERE deleted_at IS NULL;

-- ============================================================================
-- COMPOSITE INDEXES (for multi-column queries)
-- ============================================================================

-- Products: Category + Status + Featured
CREATE INDEX IF NOT EXISTS idx_products_category_status_featured 
    ON products(category_id, status, is_featured) 
    WHERE deleted_at IS NULL;

-- Products: Status + Price (for filtering active products by price)
CREATE INDEX IF NOT EXISTS idx_products_status_price 
    ON products(status, price) 
    WHERE deleted_at IS NULL AND status = 'active';

-- Orders: Customer + Status (for customer order history)
CREATE INDEX IF NOT EXISTS idx_orders_customer_status 
    ON orders(customer_id, status, created_at DESC) 
    WHERE deleted_at IS NULL;

-- Order Items: Order + Product (for order details)
CREATE INDEX IF NOT EXISTS idx_order_items_order_product 
    ON order_items(order_id, product_id) 
    WHERE deleted_at IS NULL;

-- Inventory: Product + Movement Type + Date
CREATE INDEX IF NOT EXISTS idx_inventory_product_type_date 
    ON inventory_movements(product_id, movement_type, created_at DESC);

-- ============================================================================
-- PARTIAL INDEXES (for specific query patterns)
-- ============================================================================

-- Active products only
CREATE INDEX IF NOT EXISTS idx_products_active_only 
    ON products(id, name, price, stock_quantity) 
    WHERE deleted_at IS NULL AND status = 'active';

-- Pending orders
CREATE INDEX IF NOT EXISTS idx_orders_pending 
    ON orders(id, customer_id, total_amount, created_at) 
    WHERE status = 'pending' AND deleted_at IS NULL;

-- Low stock products
CREATE INDEX IF NOT EXISTS idx_products_low_stock 
    ON products(id, name, stock_quantity, low_stock_threshold) 
    WHERE stock_quantity <= low_stock_threshold 
    AND deleted_at IS NULL 
    AND status = 'active';

-- Approved reviews only
CREATE INDEX IF NOT EXISTS idx_reviews_approved 
    ON product_reviews(product_id, rating, created_at DESC) 
    WHERE is_approved = true AND deleted_at IS NULL;



