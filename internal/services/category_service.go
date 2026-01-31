package services

import (
	"database/sql"
	"fmt"
	"log"

	"ecom/internal/dto"
)

// CategoryService handles category business logic
type CategoryService struct {
	db *sql.DB
}

// NewCategoryService creates a new category service
func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{db: db}
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	var id int64

	query := `
		INSERT INTO categories (name, slug, description, parent_id, image_url, is_active, sort_order, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	err := s.db.QueryRow(
		query,
		req.Name,
		req.Slug,
		req.Description,
		req.ParentID,
		req.ImageURL,
		req.IsActive,
		req.SortOrder,
	).Scan(&id)

	if err != nil {
		log.Printf("Error creating category: %v", err)
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	// Fetch and return the created category
	return s.GetCategoryByID(id)
}

// GetCategoryByID retrieves a category by ID
func (s *CategoryService) GetCategoryByID(id int64) (*dto.CategoryResponse, error) {
	var category dto.CategoryResponse

	query := `
		SELECT id, name, slug, description, parent_id, image_url, is_active, sort_order, created_at, updated_at
		FROM categories
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := s.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.ParentID,
		&category.ImageURL,
		&category.IsActive,
		&category.SortOrder,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("category not found")
	}
	if err != nil {
		log.Printf("Error fetching category: %v", err)
		return nil, fmt.Errorf("failed to fetch category: %w", err)
	}

	return &category, nil
}

// GetAllCategories retrieves all categories with pagination
func (s *CategoryService) GetAllCategories(page, limit int) ([]dto.CategoryResponse, int, error) {
	offset := (page - 1) * limit

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM categories WHERE deleted_at IS NULL`
	err := s.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		log.Printf("Error counting categories: %v", err)
		return nil, 0, fmt.Errorf("failed to count categories: %w", err)
	}

	// Get paginated results
	query := `
		SELECT id, name, slug, description, parent_id, image_url, is_active, sort_order, created_at, updated_at
		FROM categories
		WHERE deleted_at IS NULL
		ORDER BY sort_order ASC, id ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error fetching categories: %v", err)
		return nil, 0, fmt.Errorf("failed to fetch categories: %w", err)
	}
	defer rows.Close()

	var categories []dto.CategoryResponse
	for rows.Next() {
		var category dto.CategoryResponse
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.Description,
			&category.ParentID,
			&category.ImageURL,
			&category.IsActive,
			&category.SortOrder,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning category: %v", err)
			return nil, 0, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating categories: %v", err)
		return nil, 0, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, total, nil
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(id int64, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	// Build dynamic query based on provided fields
	query := `
		UPDATE categories
		SET `

	args := []interface{}{}
	argNum := 1

	if req.Name != "" {
		query += fmt.Sprintf("name = $%d, ", argNum)
		args = append(args, req.Name)
		argNum++
	}

	if req.Slug != "" {
		query += fmt.Sprintf("slug = $%d, ", argNum)
		args = append(args, req.Slug)
		argNum++
	}

	if req.Description != "" {
		query += fmt.Sprintf("description = $%d, ", argNum)
		args = append(args, req.Description)
		argNum++
	}

	if req.ParentID != nil {
		query += fmt.Sprintf("parent_id = $%d, ", argNum)
		args = append(args, req.ParentID)
		argNum++
	}

	if req.ImageURL != "" {
		query += fmt.Sprintf("image_url = $%d, ", argNum)
		args = append(args, req.ImageURL)
		argNum++
	}

	if req.IsActive != nil {
		query += fmt.Sprintf("is_active = $%d, ", argNum)
		args = append(args, req.IsActive)
		argNum++
	}

	if req.SortOrder != nil {
		query += fmt.Sprintf("sort_order = $%d, ", argNum)
		args = append(args, req.SortOrder)
		argNum++
	}

	// Add updated_at
	query += fmt.Sprintf("updated_at = CURRENT_TIMESTAMP WHERE id = $%d AND deleted_at IS NULL", argNum)
	args = append(args, id)

	result, err := s.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating category: %v", err)
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("category not found")
	}

	// Fetch and return the updated category
	return s.GetCategoryByID(id)
}

// DeleteCategory soft deletes a category
func (s *CategoryService) DeleteCategory(id int64) error {
	query := `
		UPDATE categories
		SET deleted_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting category: %v", err)
		return fmt.Errorf("failed to delete category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

// GetProductsByCategory gets all products in a category
func (s *CategoryService) GetProductsByCategory(categoryID int64, page, limit int) ([]dto.ProductResponse, int, error) {
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
			log.Printf("Error scanning product: %v", err)
			return nil, 0, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating products: %v", err)
		return nil, 0, fmt.Errorf("error iterating products: %w", err)
	}

	return products, total, nil
}
