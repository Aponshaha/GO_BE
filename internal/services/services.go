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
