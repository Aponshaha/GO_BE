package routes

import (
	"ecom/internal/handlers"
	"ecom/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes using Gin
func SetupRoutes(router *gin.Engine) {
	// Setup Swagger documentation routes
	SetupSwagger(router)

	// Apply global middleware
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())

	// Health check endpoint
	router.GET("/health", healthCheck)
	router.GET("/api/health", healthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Category routes
		categoryHandler := handlers.NewCategoryHandler()
		{
			v1.POST("/categories", categoryHandler.CreateCategory)
			v1.GET("/categories", categoryHandler.GetAllCategories)
			v1.GET("/categories/:id", categoryHandler.GetCategory)
			v1.PUT("/categories/:id", categoryHandler.UpdateCategory)
			v1.DELETE("/categories/:id", categoryHandler.DeleteCategory)
			v1.GET("/categories/:id/products", categoryHandler.GetCategoryProducts)
		}

		// Product routes
		productHandler := handlers.NewProductHandler()
		{
			v1.POST("/products", productHandler.CreateProduct)
			v1.GET("/products", productHandler.GetAllProducts)
			v1.GET("/products/:id", productHandler.GetProduct)
			v1.PUT("/products/:id", productHandler.UpdateProduct)
			v1.DELETE("/products/:id", productHandler.DeleteProduct)
			v1.GET("/products/category/:category_id", productHandler.GetProductsByCategoryID)
		}

		// Customer routes (placeholder for Phase 4)
		// v1.POST("/customers", customerHandler.CreateCustomer)
		// v1.GET("/customers", customerHandler.GetAllCustomers)
		// v1.GET("/customers/:id", customerHandler.GetCustomer)
		// v1.PUT("/customers/:id", customerHandler.UpdateCustomer)
		// v1.DELETE("/customers/:id", customerHandler.DeleteCustomer)

		// Order routes (placeholder for Phase 5)
		// v1.POST("/orders", orderHandler.CreateOrder)
		// v1.GET("/orders", orderHandler.GetAllOrders)
		// v1.GET("/orders/:id", orderHandler.GetOrder)
		// v1.PUT("/orders/:id", orderHandler.UpdateOrder)
		// v1.DELETE("/orders/:id", orderHandler.DeleteOrder)

		// Payment routes (placeholder for Phase 6)
		// v1.POST("/payments", paymentHandler.CreatePayment)
		// v1.GET("/payments", paymentHandler.GetAllPayments)
		// v1.GET("/payments/:id", paymentHandler.GetPayment)
		// v1.PUT("/payments/:id", paymentHandler.UpdatePayment)
	}
}

// healthCheck is a simple health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Server is running",
	})
}

