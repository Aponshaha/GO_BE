package middleware

import (
	"net/http"
	"time"

	"ecom/internal/dto"
	"github.com/gin-gonic/gin"
)

// SuccessResponse returns a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	response := dto.SuccessResponse{
		Success:   true,
		Data:      data,
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	c.JSON(statusCode, response)
}

// ListResponse returns a paginated list response
func ListResponse(c *gin.Context, data interface{}, page, limit, total int, message string) {
	pages := (total + limit - 1) / limit // Calculate total pages
	if pages < 1 {
		pages = 1
	}

	response := dto.ListResponseData{
		Success: true,
		Data:    data,
		Pagination: &dto.Pagination{
			Page:  page,
			Limit: limit,
			Total: total,
			Pages: pages,
		},
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, response)
}

// ErrorResponse returns a standardized error response
func ErrorResponse(c *gin.Context, statusCode int, error, message string) {
	response := dto.ErrorResponseData{
		Success:   false,
		Error:     error,
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	c.JSON(statusCode, response)
}

// ===========================
// Convenience Methods
// ===========================

// Created returns 201 Created response
func Created(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusCreated, data, message)
}

// OK returns 200 OK response
func OK(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusOK, data, message)
}

// BadRequest returns 400 Bad Request response
func BadRequest(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, "Bad Request", message)
}

// Unauthorized returns 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", message)
}

// Forbidden returns 403 Forbidden response
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, "Forbidden", message)
}

// NotFound returns 404 Not Found response
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, "Not Found", message)
}

// Conflict returns 409 Conflict response
func Conflict(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, "Conflict", message)
}

// InternalError returns 500 Internal Server Error response
func InternalError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", message)
}

// ServiceUnavailable returns 503 Service Unavailable response
func ServiceUnavailable(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusServiceUnavailable, "Service Unavailable", message)
}
