package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"ecom/internal/database"
	"ecom/internal/models"
	"ecom/internal/services"
)

// Handler holds all handlers and their dependencies
type Handler struct {
	userService    *services.UserService
	productService *services.ProductService
}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	return &Handler{
		userService:    services.NewUserService(),
		productService: services.NewProductService(database.GetDB()),
	}
}

// HealthCheck handles the health check endpoint
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := models.HealthCheck{
		Status:    "healthy",
		Message:   "Server is running",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetUser handles GET /api/users/:id
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL path
	// In a real implementation, you'd use a router like gorilla/mux or chi
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAllProducts handles GET /api/products
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10
	
	products, _, err := h.productService.GetAllProducts(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

