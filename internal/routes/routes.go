package routes

import (
	"net/http"

	"ecom/internal/handlers"
	"ecom/internal/middleware"
)

// SetupRoutes configures all application routes
func SetupRoutes() http.Handler {
	handler := handlers.NewHandler()

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/api/health", handler.HealthCheck)

	// API routes
	api := http.NewServeMux()
	api.HandleFunc("/users", handler.GetUser)
	api.HandleFunc("/products", handler.GetAllProducts)
	
	mux.Handle("/api/", http.StripPrefix("/api", api))

	// Apply middleware
	var finalHandler http.Handler = mux
	finalHandler = middleware.CORSMiddleware(finalHandler)
	finalHandler = middleware.LoggingMiddleware(finalHandler)

	return finalHandler
}

