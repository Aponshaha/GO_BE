package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware for handling panics and errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred")
			}
		}()
		c.Next()
	}
}

// RequestLogger middleware for logging HTTP requests
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		log.Printf(
			"%s %s %s | Status: %d | Duration: %v | IP: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			statusCode,
			duration,
			c.ClientIP(),
		)
	}
}

// CORSMiddleware for handling CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestIDMiddleware adds a unique request ID to context
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("X-Request-ID")
		if requestID == "" {
			requestID = time.Now().Format("20060102150405") + "-" + c.ClientIP()
		}
		c.Set("requestID", requestID)
		c.Next()
	}
}

// TimeoutMiddleware for request timeouts
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// This is handled at server level, but you can add custom timeout logic here if needed
		c.Next()
	}
}
