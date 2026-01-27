package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams extracts pagination parameters from request
func PaginationParams(c *gin.Context) (page int, limit int) {
	page = 1
	limit = 10

	// Get page from query params
	if p := c.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Get limit from query params
	if l := c.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			if parsedLimit > 100 {
				parsedLimit = 100 // Max limit is 100
			}
			limit = parsedLimit
		}
	}

	return page, limit
}

// Offset calculates the offset for database query
func Offset(page, limit int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * limit
}

// CalculatePages calculates total number of pages
func CalculatePages(total, limit int) int {
	if limit <= 0 {
		return 1
	}
	pages := (total + limit - 1) / limit
	if pages < 1 {
		pages = 1
	}
	return pages
}
