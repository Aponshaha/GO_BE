package dto

import "time"

// ===========================
// Pagination DTOs
// ===========================

type PaginationParams struct {
	Page  int `form:"page" binding:"min=1" default:"1"`
	Limit int `form:"limit" binding:"min=1,max=100" default:"10"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

// ===========================
// Response Wrappers
// ===========================

type SuccessResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
}

type ListResponseData struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
}

type ErrorResponseData struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// ===========================
// Category Request DTOs
// ===========================

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Slug        string `json:"slug" binding:"required,max=100"`
	Description string `json:"description" binding:"max=5000"`
	ParentID    *int64 `json:"parent_id,omitempty"`
	ImageURL    string `json:"image_url" binding:"max=255"`
	IsActive    bool   `json:"is_active" default:"true"`
	SortOrder   int    `json:"sort_order" default:"0"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"max=100"`
	Slug        string `json:"slug" binding:"max=100"`
	Description string `json:"description" binding:"max=5000"`
	ParentID    *int64 `json:"parent_id"`
	ImageURL    string `json:"image_url" binding:"max=255"`
	IsActive    *bool  `json:"is_active"`
	SortOrder   *int   `json:"sort_order"`
}

// ===========================
// Product Request DTOs
// ===========================

type CreateProductRequest struct {
	SKU                string   `json:"sku" binding:"required,max=50"`
	Name               string   `json:"name" binding:"required,max=200"`
	Slug               string   `json:"slug" binding:"required,max=200"`
	Description        string   `json:"description" binding:"max=5000"`
	ShortDescription   string   `json:"short_description" binding:"max=500"`
	CategoryID         int64    `json:"category_id" binding:"required"`
	Status             string   `json:"status" binding:"required,oneof=active inactive out_of_stock discontinued"`
	Price              float64  `json:"price" binding:"required,gt=0"`
	CompareAtPrice     *float64 `json:"compare_at_price,omitempty"`
	CostPrice          *float64 `json:"cost_price,omitempty"`
	StockQuantity      int      `json:"stock_quantity" binding:"required,gte=0"`
	LowStockThreshold  int      `json:"low_stock_threshold" default:"10"`
	WeightKg           *float64 `json:"weight_kg,omitempty"`
	DimensionsCm       string   `json:"dimensions_cm" binding:"max=50"`
	Barcode            string   `json:"barcode" binding:"max=100"`
	Manufacturer       string   `json:"manufacturer" binding:"max=100"`
	Brand              string   `json:"brand" binding:"max=100"`
	IsFeautred         bool     `json:"is_featured" default:"false"`
	MetaTitle          string   `json:"meta_title" binding:"max=200"`
	MetaDescription    string   `json:"meta_description" binding:"max=500"`
}

type UpdateProductRequest struct {
	Name               string   `json:"name" binding:"max=200"`
	Slug               string   `json:"slug" binding:"max=200"`
	Description        string   `json:"description" binding:"max=5000"`
	ShortDescription   string   `json:"short_description" binding:"max=500"`
	CategoryID         *int64   `json:"category_id"`
	Status             string   `json:"status" binding:"oneof=active inactive out_of_stock discontinued"`
	Price              *float64 `json:"price" binding:"gt=0"`
	CompareAtPrice     *float64 `json:"compare_at_price"`
	StockQuantity      *int     `json:"stock_quantity" binding:"gte=0"`
	LowStockThreshold  *int     `json:"low_stock_threshold"`
	WeightKg           *float64 `json:"weight_kg"`
	DimensionsCm       string   `json:"dimensions_cm" binding:"max=50"`
	Brand              string   `json:"brand" binding:"max=100"`
	IsFeautred         *bool    `json:"is_featured"`
	MetaTitle          string   `json:"meta_title" binding:"max=200"`
	MetaDescription    string   `json:"meta_description" binding:"max=500"`
}

type UpdateProductStockRequest struct {
	Quantity int    `json:"quantity" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=add subtract set"`
	Notes    string `json:"notes"`
}

// ===========================
// Customer Request DTOs
// ===========================

type CreateCustomerRequest struct {
	Email       string `json:"email" binding:"required,email,max=255"`
	FirstName   string `json:"first_name" binding:"required,max=100"`
	LastName    string `json:"last_name" binding:"required,max=100"`
	Phone       string `json:"phone" binding:"max=20"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	IsActive    bool   `json:"is_active" default:"true"`
}

type UpdateCustomerRequest struct {
	Email       string     `json:"email" binding:"email,max=255"`
	FirstName   string     `json:"first_name" binding:"max=100"`
	LastName    string     `json:"last_name" binding:"max=100"`
	Phone       string     `json:"phone" binding:"max=20"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	IsActive    *bool      `json:"is_active"`
}

type CreateAddressRequest struct {
	AddressType string `json:"address_type" binding:"required,oneof=shipping billing both"`
	FirstName   string `json:"first_name" binding:"required,max=100"`
	LastName    string `json:"last_name" binding:"required,max=100"`
	Company     string `json:"company" binding:"max=100"`
	AddressLine1 string `json:"address_line1" binding:"required,max=200"`
	AddressLine2 string `json:"address_line2" binding:"max=200"`
	City        string `json:"city" binding:"required,max=100"`
	State       string `json:"state" binding:"max=100"`
	PostalCode  string `json:"postal_code" binding:"required,max=20"`
	Country     string `json:"country" binding:"required,max=100"`
	Phone       string `json:"phone" binding:"max=20"`
	IsDefault   bool   `json:"is_default" default:"false"`
}

// ===========================
// Order Request DTOs
// ===========================

type CreateOrderItemRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required,gt=0"`
}

type CreateOrderRequest struct {
	CustomerID         int64                     `json:"customer_id" binding:"required"`
	Items              []CreateOrderItemRequest  `json:"items" binding:"required,min=1"`
	ShippingAddressID  *int64                    `json:"shipping_address_id"`
	BillingAddressID   *int64                    `json:"billing_address_id"`
	ShippingAmount     float64                   `json:"shipping_amount" binding:"gte=0" default:"0"`
	DiscountAmount     float64                   `json:"discount_amount" binding:"gte=0" default:"0"`
	Notes              string                    `json:"notes"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending confirmed processing shipped delivered cancelled refunded"`
	Notes  string `json:"notes"`
}

type CancelOrderRequest struct {
	Reason string `json:"reason" binding:"required,max=500"`
}

// ===========================
// Payment Request DTOs
// ===========================

type CreatePaymentRequest struct {
	OrderID       int64   `json:"order_id" binding:"required"`
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=credit_card debit_card paypal bank_transfer cash_on_delivery"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
}

type UpdatePaymentStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending processing completed failed refunded"`
}
