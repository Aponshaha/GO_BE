package repositories

import (
	"database/sql"
	"ecom/internal/database"
	"ecom/internal/models"
	"fmt"
)

// UserRepository handles user-related database operations
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, created_at, updated_at 
	          FROM users WHERE id = $1`
	
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	
	return user, nil
}

// ProductRepository handles product-related database operations
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		db: database.GetDB(),
	}
}

// GetAllProducts retrieves all products
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at 
	          FROM products ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	
	return products, nil
}

