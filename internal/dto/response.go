package dto

import "time"

// ===========================
// Category Response DTOs
// ===========================

type CategoryResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	ParentID    *int64    `json:"parent_id,omitempty"`
	ImageURL    *string   `json:"image_url,omitempty"`
	IsActive    bool      `json:"is_active"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ===========================
// Product Response DTOs
// ===========================

type ProductResponse struct {
	ID                int64      `json:"id"`
	SKU               string     `json:"sku"`
	Name              string     `json:"name"`
	Slug              string     `json:"slug"`
	Description       string     `json:"description,omitempty"`
	ShortDescription  string     `json:"short_description,omitempty"`
	CategoryID        int64      `json:"category_id"`
	Category          *CategoryResponse `json:"category,omitempty"`
	Status            string     `json:"status"`
	Price             float64    `json:"price"`
	CompareAtPrice    *float64   `json:"compare_at_price,omitempty"`
	CostPrice         *float64   `json:"cost_price,omitempty"`
	StockQuantity     int        `json:"stock_quantity"`
	LowStockThreshold int        `json:"low_stock_threshold"`
	WeightKg          *float64   `json:"weight_kg,omitempty"`
	DimensionsCm      *string    `json:"dimensions_cm,omitempty"`
	Barcode           *string    `json:"barcode,omitempty"`
	Manufacturer      *string    `json:"manufacturer,omitempty"`
	Brand             *string    `json:"brand,omitempty"`
	RatingAverage     float64    `json:"rating_average"`
	RatingCount       int        `json:"rating_count"`
	ViewCount         int        `json:"view_count"`
	IsFeautred        bool       `json:"is_featured"`
	MetaTitle         *string    `json:"meta_title,omitempty"`
	MetaDescription   *string    `json:"meta_description,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type ProductImageResponse struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	AltText   string    `json:"alt_text,omitempty"`
	SortOrder int       `json:"sort_order"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}

// ===========================
// Customer Response DTOs
// ===========================

type CustomerResponse struct {
	ID               int64     `json:"id"`
	Email            string    `json:"email"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Phone            string    `json:"phone,omitempty"`
	DateOfBirth      *time.Time `json:"date_of_birth,omitempty"`
	IsActive         bool      `json:"is_active"`
	EmailVerifiedAt  *time.Time `json:"email_verified_at,omitempty"`
	LastLoginAt      *time.Time `json:"last_login_at,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CustomerDetailResponse struct {
	*CustomerResponse
	Addresses []AddressResponse `json:"addresses,omitempty"`
	Orders    []OrderResponse   `json:"orders,omitempty"`
}

type AddressResponse struct {
	ID           int64     `json:"id"`
	CustomerID   int64     `json:"customer_id"`
	AddressType  string    `json:"address_type"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Company      string    `json:"company,omitempty"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2,omitempty"`
	City         string    `json:"city"`
	State        string    `json:"state,omitempty"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	Phone        string    `json:"phone,omitempty"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ===========================
// Order Response DTOs
// ===========================

type OrderResponse struct {
	ID               int64              `json:"id"`
	OrderNumber      string             `json:"order_number"`
	CustomerID       int64              `json:"customer_id"`
	Customer         *CustomerResponse  `json:"customer,omitempty"`
	Status           string             `json:"status"`
	Items            []OrderItemResponse `json:"items,omitempty"`
	Subtotal         float64            `json:"subtotal"`
	TaxAmount        float64            `json:"tax_amount"`
	ShippingAmount   float64            `json:"shipping_amount"`
	DiscountAmount   float64            `json:"discount_amount"`
	TotalAmount      float64            `json:"total_amount"`
	Currency         string             `json:"currency"`
	ShippingAddress  *AddressResponse   `json:"shipping_address,omitempty"`
	BillingAddress   *AddressResponse   `json:"billing_address,omitempty"`
	Notes            string             `json:"notes,omitempty"`
	CancelledAt      *time.Time         `json:"cancelled_at,omitempty"`
	CancelledReason  string             `json:"cancelled_reason,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

type OrderItemResponse struct {
	ID             int64   `json:"id"`
	OrderID        int64   `json:"order_id"`
	ProductID      int64   `json:"product_id"`
	SKU            string  `json:"sku"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalPrice     float64 `json:"total_price"`
	CreatedAt      time.Time `json:"created_at"`
}

// ===========================
// Payment Response DTOs
// ===========================

type PaymentResponse struct {
	ID              int64              `json:"id"`
	OrderID         int64              `json:"order_id"`
	Order           *OrderResponse     `json:"order,omitempty"`
	PaymentMethod   string             `json:"payment_method"`
	Status          string             `json:"status"`
	Amount          float64            `json:"amount"`
	Currency        string             `json:"currency"`
	TransactionID   string             `json:"transaction_id,omitempty"`
	GatewayResponse interface{}        `json:"gateway_response,omitempty"`
	PaidAt          *time.Time         `json:"paid_at,omitempty"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}
