package services

import (
	"ecom/internal/models"
	"ecom/internal/repositories"
)

// UserService handles user-related business logic
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(),
	}
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

// ProductService handles product-related business logic
type ProductService struct {
	productRepo *repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repositories.NewProductRepository(),
	}
}

// GetAllProducts retrieves all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAllProducts()
}
