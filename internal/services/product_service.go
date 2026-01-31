package services

import (
	"database/sql"
	"fmt"
	"log"

	"ecom/internal/dto"
)

// ProductService handles product business logic
type ProductService struct {
	db *sql.DB
}

// NewProductService creates a new product service
func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	var id int64

	query := `
		INSERT INTO products (sku, name, slug, description, short_description, category_id, status, price, compare_at_price, stock_quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	err := s.db.QueryRow(
		query,
		req.SKU,
		req.Name,
		req.Slug,
		req.Description,
		req.ShortDescription,
		req.CategoryID,
		req.Status,
		req.Price,
		req.CompareAtPrice,
		req.StockQuantity,
	).Scan(&id)

	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, fmt.Errorf("failed to create product: %w", err)
	}	

	// Fetch and return the created product
	return s.GetProductByID(id)
}

func(s *ProductService) GetProductByID(id int64) (*dto.ProductResponse, error) {
	var product dto.ProductResponse

	query := `SELECT id, sku, name, slug, description, short_description, category_id, status, price, compare_at_price, cost_price, stock_quantity, low_stock_threshold, weight_kg, dimensions_cm, barcode, manufacturer, brand, rating_average, rating_count, view_count, is_featured, meta_title, meta_description, created_at, updated_at 
	 FROM products 
	 WHERE id = $1 AND deleted_at IS NULL`

	 err := s.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.ShortDescription,
		&product.CategoryID,
		&product.Status,
		&product.Price,
		&product.CompareAtPrice,
		&product.CostPrice,
		&product.StockQuantity,
		&product.LowStockThreshold,
		&product.WeightKg,
		&product.DimensionsCm,
		&product.Barcode,
		&product.Manufacturer,
		&product.Brand,
		&product.RatingAverage,
		&product.RatingCount,
		&product.ViewCount,
		&product.IsFeautred,
		&product.MetaTitle,
		&product.MetaDescription,
		&product.CreatedAt,
		&product.UpdatedAt,	
	 )

	 if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	 }

	 if err != nil {
		log.Printf("Error fetching product by ID: %v", err)
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	 }

	 return &product, nil
}

// GetAllProductsByCategory retrieves products by category with pagination
func (s *ProductService) GetAllProductsByCategory(categoryID int64, page, limit int) ([]dto.ProductResponse, int, error) {
	offset := (page - 1) * limit
	
	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM products WHERE category_id = $1 AND deleted_at IS NULL`
	err := s.db.QueryRow(countQuery, categoryID).Scan(&total)
	if err != nil {
		log.Printf("Error counting products: %v", err)
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Get paginated results
	query := `
		SELECT id, sku, name, slug, description, short_description, category_id, status, price, 
		       compare_at_price, cost_price, stock_quantity, low_stock_threshold, weight_kg, 
		       dimensions_cm, barcode, manufacturer, brand, rating_average, rating_count, 
		       view_count, is_featured, meta_title, meta_description, created_at, updated_at
		FROM products
		WHERE category_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	rows, err := s.db.Query(query, categoryID, limit, offset)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()
	
	var products []dto.ProductResponse
	for rows.Next() {
		var product dto.ProductResponse
		err := rows.Scan(
			&product.ID,
			&product.SKU,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.ShortDescription,
			&product.CategoryID,
			&product.Status,
			&product.Price,
			&product.CompareAtPrice,
			&product.CostPrice,
			&product.StockQuantity,
			&product.LowStockThreshold,
			&product.WeightKg,
			&product.DimensionsCm,
			&product.Barcode,
			&product.Manufacturer,
			&product.Brand,
			&product.RatingAverage,
			&product.RatingCount,
			&product.ViewCount,
			&product.IsFeautred,
			&product.MetaTitle,
			&product.MetaDescription,
			&product.CreatedAt,
			&product.UpdatedAt,	
		)
		if err != nil {
			log.Printf("Error scanning product row: %v", err)
			return nil, 0, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, 0, fmt.Errorf("row iteration error: %w", err)
	}

	return products, total, nil
}	

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(productID int64, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	query := `
		UPDATE products SET
			name = COALESCE(NULLIF($1, ''), name),
			slug = COALESCE(NULLIF($2, ''), slug),
			description = COALESCE(NULLIF($3, ''), description),
			short_description = COALESCE(NULLIF($4, ''), short_description),
			category_id = COALESCE($5, category_id),
			status = COALESCE(NULLIF($6, ''), status),
			price = COALESCE($7, price),
			compare_at_price = COALESCE($8, compare_at_price),
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $9 AND deleted_at IS NULL
		RETURNING id
	`
	var id int64
	err := s.db.QueryRow(
		query,
		req.Name,
		req.Slug,
		req.Description,
		req.ShortDescription,
		req.CategoryID,		
		req.Status,	
		req.Price,
		req.CompareAtPrice,
		productID,
	).Scan(&id)
	
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	}

	if err != nil {
		log.Printf("Error updating product: %v", err)
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	// Fetch and return the updated product
	return s.GetProductByID(id)
}

// GetAllProducts retrieves all products with pagination
func (s *ProductService) GetAllProducts(page, limit int) ([]dto.ProductResponse, int, error) {
	offset := (page - 1) * limit
	
	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM products WHERE deleted_at IS NULL`
	err := s.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		log.Printf("Error counting products: %v", err)
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Get paginated results
	query := `
		SELECT id, sku, name, slug, description, short_description, category_id, status, price, 
		       compare_at_price, cost_price, stock_quantity, low_stock_threshold, weight_kg, 
		       dimensions_cm, barcode, manufacturer, brand, rating_average, rating_count, 
		       view_count, is_featured, meta_title, meta_description, created_at, updated_at
		FROM products
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()
	
	var products []dto.ProductResponse
	for rows.Next() {
		var product dto.ProductResponse
		err := rows.Scan(
			&product.ID,
			&product.SKU,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.ShortDescription,
			&product.CategoryID,
			&product.Status,
			&product.Price,
			&product.CompareAtPrice,
			&product.CostPrice,
			&product.StockQuantity,
			&product.LowStockThreshold,
			&product.WeightKg,
			&product.DimensionsCm,
			&product.Barcode,
			&product.Manufacturer,
			&product.Brand,
			&product.RatingAverage,
			&product.RatingCount,
			&product.ViewCount,
			&product.IsFeautred,
			&product.MetaTitle,
			&product.MetaDescription,
			&product.CreatedAt,
			&product.UpdatedAt,	
		)
		if err != nil {
			log.Printf("Error scanning product row: %v", err)
			return nil, 0, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, 0, fmt.Errorf("row iteration error: %w", err)
	}

	return products, total, nil
}

// DeleteProduct soft deletes a product
func (s *ProductService) DeleteProduct(id int64) error {
	query := `
		UPDATE products
		SET deleted_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting product: %v", err)
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}