package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIDParam extracts and validates ID from URL parameter
func GetIDParam(c *gin.Context, paramName string) (int64, error) {
	idStr := c.Param(paramName)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, NewValidationError("invalid " + paramName)
	}
	return id, nil
}

// GetQueryInt gets an integer from query parameters
func GetQueryInt(c *gin.Context, key string, defaultVal int) int {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return intVal
}

// GetQueryString gets a string from query parameters
func GetQueryString(c *gin.Context, key string, defaultVal string) string {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// NewValidationError creates a validation error
func NewValidationError(message string) error {
	return &ValidationError{Message: message}
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
