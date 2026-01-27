# REST API Development Plan with Swagger

## 📊 Project Overview

E-commerce API with full REST endpoints and OpenAPI (Swagger) documentation.

---

## 🏗️ Project Structure

```
internal/
├── handlers/
│   ├── category_handler.go
│   ├── product_handler.go
│   ├── customer_handler.go
│   ├── order_handler.go
│   ├── payment_handler.go
│   └── common.go (shared response structures)
├── routes/
│   ├── routes.go (main router setup)
│   └── swagger.go (swagger initialization)
├── dto/
│   ├── request.go (request DTOs)
│   ├── response.go (response DTOs)
│   └── pagination.go
└── services/
    ├── category_service.go
    ├── product_service.go
    ├── customer_service.go
    ├── order_service.go
    └── payment_service.go
```

---

## 📋 REST API Endpoints

### 1. Categories

| Method | Endpoint                          | Description                     | Status  |
| ------ | --------------------------------- | ------------------------------- | ------- |
| GET    | `/api/v1/categories`              | List all categories (paginated) | 🔴 TODO |
| GET    | `/api/v1/categories/:id`          | Get category by ID              | 🔴 TODO |
| POST   | `/api/v1/categories`              | Create new category             | 🔴 TODO |
| PUT    | `/api/v1/categories/:id`          | Update category                 | 🔴 TODO |
| DELETE | `/api/v1/categories/:id`          | Delete category (soft delete)   | 🔴 TODO |
| GET    | `/api/v1/categories/:id/products` | Get products by category        | 🔴 TODO |

### 2. Products

| Method | Endpoint                      | Description                               | Status  |
| ------ | ----------------------------- | ----------------------------------------- | ------- |
| GET    | `/api/v1/products`            | List all products (paginated, filterable) | 🔴 TODO |
| GET    | `/api/v1/products/:id`        | Get product by ID                         | 🔴 TODO |
| POST   | `/api/v1/products`            | Create new product                        | 🔴 TODO |
| PUT    | `/api/v1/products/:id`        | Update product                            | 🔴 TODO |
| DELETE | `/api/v1/products/:id`        | Delete product (soft delete)              | 🔴 TODO |
| GET    | `/api/v1/products/search`     | Search products (by name, sku)            | 🔴 TODO |
| GET    | `/api/v1/products/:id/images` | Get product images                        | 🔴 TODO |
| POST   | `/api/v1/products/:id/images` | Upload product images                     | 🔴 TODO |
| PUT    | `/api/v1/products/:id/stock`  | Update stock                              | 🔴 TODO |

### 3. Customers

| Method | Endpoint                          | Description                    | Status  |
| ------ | --------------------------------- | ------------------------------ | ------- |
| GET    | `/api/v1/customers`               | List all customers (paginated) | 🔴 TODO |
| GET    | `/api/v1/customers/:id`           | Get customer by ID             | 🔴 TODO |
| POST   | `/api/v1/customers`               | Create new customer            | 🔴 TODO |
| PUT    | `/api/v1/customers/:id`           | Update customer                | 🔴 TODO |
| DELETE | `/api/v1/customers/:id`           | Delete customer (soft delete)  | 🔴 TODO |
| GET    | `/api/v1/customers/:id/addresses` | Get customer addresses         | 🔴 TODO |
| POST   | `/api/v1/customers/:id/addresses` | Add address                    | 🔴 TODO |
| GET    | `/api/v1/customers/:id/orders`    | Get customer orders            | 🔴 TODO |

### 4. Orders

| Method | Endpoint                     | Description                 | Status  |
| ------ | ---------------------------- | --------------------------- | ------- |
| GET    | `/api/v1/orders`             | List all orders (paginated) | 🔴 TODO |
| GET    | `/api/v1/orders/:id`         | Get order by ID             | 🔴 TODO |
| POST   | `/api/v1/orders`             | Create new order            | 🔴 TODO |
| PUT    | `/api/v1/orders/:id`         | Update order status         | 🔴 TODO |
| DELETE | `/api/v1/orders/:id`         | Cancel order                | 🔴 TODO |
| GET    | `/api/v1/orders/:id/items`   | Get order items             | 🔴 TODO |
| GET    | `/api/v1/orders/:id/payment` | Get order payment info      | 🔴 TODO |
| POST   | `/api/v1/orders/:id/payment` | Process payment             | 🔴 TODO |

### 5. Payments

| Method | Endpoint               | Description           | Status  |
| ------ | ---------------------- | --------------------- | ------- |
| GET    | `/api/v1/payments`     | List all payments     | 🔴 TODO |
| GET    | `/api/v1/payments/:id` | Get payment by ID     | 🔴 TODO |
| POST   | `/api/v1/payments`     | Create payment        | 🔴 TODO |
| PUT    | `/api/v1/payments/:id` | Update payment status | 🔴 TODO |

### 6. Health & Status

| Method | Endpoint         | Description  | Status    |
| ------ | ---------------- | ------------ | --------- |
| GET    | `/health`        | Health check | ✅ EXISTS |
| GET    | `/api/v1/status` | API status   | 🔴 TODO   |

---

## 🔐 Response Format

### Success Response (200)

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Product Name",
    ...
  },
  "message": "Operation successful",
  "timestamp": "2026-01-27T15:30:00Z"
}
```

### List Response with Pagination (200)

```json
{
  "success": true,
  "data": [
    { "id": 1, "name": "Product 1" },
    { "id": 2, "name": "Product 2" }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "pages": 10
  },
  "message": "Retrieved successfully",
  "timestamp": "2026-01-27T15:30:00Z"
}
```

### Error Response (4xx/5xx)

```json
{
  "success": false,
  "error": "Not found",
  "message": "Product with ID 999 not found",
  "timestamp": "2026-01-27T15:30:00Z"
}
```

---

## 🔌 Swagger/OpenAPI Setup

### 1. Install Swagger Tools

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 2. Generate Swagger Docs

```bash
# Generate from code comments
swag init -g cmd/server/main.go

# View docs at: http://localhost:8080/swagger/index.html
```

### 3. Swagger Annotations Example

```go
// GetProducts godoc
// @Summary Get all products
// @Description Get paginated list of products with optional filters
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param category_id query int false "Filter by category"
// @Param status query string false "Filter by status" Enums(active,inactive,discontinued)
// @Success 200 {object} dto.ListResponse{data=[]dto.ProductResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
    // implementation
}
```

---

## 📦 DTO Structures

### Request DTOs (dto/request.go)

```go
// CreateProductRequest
type CreateProductRequest struct {
    SKU              string  `json:"sku" binding:"required,max=50"`
    Name             string  `json:"name" binding:"required,max=200"`
    Description      string  `json:"description" binding:"max=5000"`
    Price            float64 `json:"price" binding:"required,gt=0"`
    Stock            int     `json:"stock_quantity" binding:"required,gte=0"`
    CategoryID       int64   `json:"category_id" binding:"required"`
    Status           string  `json:"status" binding:"required,oneof=active inactive discontinued"`
    Brand            string  `json:"brand" binding:"max=100"`
    IsFeautred       bool    `json:"is_featured"`
}

// CreateOrderRequest
type CreateOrderRequest struct {
    CustomerID      int64                    `json:"customer_id" binding:"required"`
    Items           []CreateOrderItemRequest `json:"items" binding:"required,min=1"`
    ShippingAmount  float64                  `json:"shipping_amount" binding:"gte=0"`
    DiscountAmount  float64                  `json:"discount_amount" binding:"gte=0"`
}

type CreateOrderItemRequest struct {
    ProductID int64 `json:"product_id" binding:"required"`
    Quantity  int   `json:"quantity" binding:"required,gt=0"`
}

// PaginationParams
type PaginationParams struct {
    Page  int `form:"page" binding:"min=1" default:"1"`
    Limit int `form:"limit" binding:"min=1,max=100" default:"10"`
}
```

### Response DTOs (dto/response.go)

```go
// ProductResponse
type ProductResponse struct {
    ID                int64     `json:"id"`
    SKU               string    `json:"sku"`
    Name              string    `json:"name"`
    Price             float64   `json:"price"`
    CompareAtPrice    *float64  `json:"compare_at_price,omitempty"`
    Stock             int       `json:"stock_quantity"`
    Status            string    `json:"status"`
    Brand             string    `json:"brand"`
    IsFeautred        bool      `json:"is_featured"`
    RatingAverage     float64   `json:"rating_average"`
    RatingCount       int       `json:"rating_count"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
}

// OrderResponse
type OrderResponse struct {
    ID              int64                `json:"id"`
    OrderNumber     string               `json:"order_number"`
    CustomerID      int64                `json:"customer_id"`
    Status          string               `json:"status"`
    Subtotal        float64              `json:"subtotal"`
    Tax             float64              `json:"tax_amount"`
    ShippingAmount  float64              `json:"shipping_amount"`
    DiscountAmount  float64              `json:"discount_amount"`
    TotalAmount     float64              `json:"total_amount"`
    Items           []OrderItemResponse  `json:"items"`
    CreatedAt       time.Time            `json:"created_at"`
}

type OrderItemResponse struct {
    ID         int64   `json:"id"`
    ProductID  int64   `json:"product_id"`
    SKU        string  `json:"sku"`
    Name       string  `json:"name"`
    Quantity   int     `json:"quantity"`
    UnitPrice  float64 `json:"unit_price"`
    TotalPrice float64 `json:"total_price"`
}

// ListResponse
type ListResponse struct {
    Data       interface{} `json:"data"`
    Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
    Page  int `json:"page"`
    Limit int `json:"limit"`
    Total int `json:"total"`
    Pages int `json:"pages"`
}

// ErrorResponse
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
}
```

---

## 🔄 Implementation Steps

### Phase 1: Setup (Week 1)

- [ ] Install Swagger dependencies
- [ ] Create DTO structures
- [ ] Setup middleware for error handling & logging
- [ ] Create common response utilities
- [ ] Setup pagination helpers

### Phase 2: Category APIs (Week 1-2)

- [ ] Create category service
- [ ] Create category handler
- [ ] Add Swagger annotations
- [ ] Write unit tests
- [ ] Integration testing

### Phase 3: Product APIs (Week 2-3)

- [ ] Create product service
- [ ] Create product handler
- [ ] Add search/filter functionality
- [ ] Image upload handling
- [ ] Stock management
- [ ] Swagger annotations & tests

### Phase 4: Customer APIs (Week 3-4)

- [ ] Create customer service
- [ ] Address management
- [ ] Customer handler
- [ ] Swagger annotations & tests

### Phase 5: Order APIs (Week 4-5)

- [ ] Create order service
- [ ] Order processing logic
- [ ] Order handler
- [ ] Cart management
- [ ] Swagger annotations & tests

### Phase 6: Payment APIs (Week 5)

- [ ] Create payment service
- [ ] Payment processing
- [ ] Payment handler
- [ ] Swagger annotations & tests

### Phase 7: Documentation & Testing (Week 6)

- [ ] Complete Swagger documentation
- [ ] Integration tests
- [ ] Performance testing
- [ ] API documentation review

---

## 🔗 Middleware & Utilities

### Required Middleware

```go
// Error handling middleware
func ErrorHandler() gin.HandlerFunc {}

// Request logging middleware
func RequestLogger() gin.HandlerFunc {}

// Pagination middleware
func PaginationMiddleware() gin.HandlerFunc {}

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {}
```

### Response Utilities

```go
// Success response
func SuccessResponse(data interface{}, message string) {}

// List response with pagination
func ListResponse(data interface{}, page, limit, total int) {}

// Error response
func ErrorResponse(status int, message string) {}
```

---

## 📚 Swagger Annotations Checklist

For each endpoint:

- [ ] @Summary - Brief description
- [ ] @Description - Detailed description
- [ ] @Tags - Group endpoints
- [ ] @Accept - Input format (json, form-data)
- [ ] @Produce - Output format (json)
- [ ] @Param - All parameters with types & validations
- [ ] @Success - Success response structure
- [ ] @Failure - Error response structures
- [ ] @Router - HTTP method & path

---

## 🧪 Testing Strategy

### Unit Tests

- Handler tests
- Service tests
- DTO validation tests

### Integration Tests

- Full endpoint tests
- Database transaction tests
- Error handling tests

### API Documentation Tests

- Swagger validation
- Request/response examples

---

## 📊 Example Swagger Entry Point

```go
// @title E-commerce API
// @version 1.0
// @description E-commerce REST API with full CRUD operations
// @host localhost:8080
// @basePath /api/v1
// @schemes http https
// @consumes application/json
// @produces application/json

func SetupSwagger() error {
    // swagger docs setup
}
```

---

## 🎯 Next Steps

1. **Approve this plan** - Review and adjust as needed
2. **Setup Phase** - Create DTOs and utilities
3. **Category APIs** - First endpoint implementation
4. **Expand** - Follow the phase breakdown
5. **Document** - Keep Swagger updated

---

## 📝 Notes

- All endpoints use `/api/v1/` prefix for versioning
- Soft deletes maintained with `deleted_at` field
- Pagination defaults: page=1, limit=10, max=100
- All timestamps in UTC (ISO 8601 format)
- Status codes: 200 (success), 400 (bad request), 404 (not found), 500 (server error)
